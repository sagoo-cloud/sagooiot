package datahub

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDataTemplateNode struct{}

func init() {
	service.RegisterDataTemplateNode(dataTemplateNodeNew())
}

func dataTemplateNodeNew() *sDataTemplateNode {
	return &sDataTemplateNode{}
}

func (s *sDataTemplateNode) Add(ctx context.Context, in *model.DataTemplateNodeAddInput) (err error) {
	id, _ := dao.DataTemplateNode.Ctx(ctx).
		Fields(dao.DataTemplateNode.Columns().Id).
		Where(dao.DataTemplateNode.Columns().Tid, in.Tid).
		Where(dao.DataTemplateNode.Columns().Key, in.Key).
		Value()
	if id.Int64() > 0 {
		return gerror.New("数据模型节点标识重复")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataTemplateNode
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)
	if in.Method == "" {
		param.Method = nil
	}

	err = dao.DataTemplateNode.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		rs, err := dao.DataTemplateNode.Ctx(ctx).Data(param).Insert()
		if err != nil {
			return err
		}

		nodeId, _ := rs.LastInsertId()

		dataTemplate, _ := service.DataTemplate().Detail(ctx, in.Tid)
		if dataTemplate != nil && dataTemplate.DataTable != "" {
			// 表结构已存在，字段新增
			err = addTplColumn(ctx, uint64(nodeId))
			if err != nil {
				return err
			}
		}

		// 处理排序字段
		if in.IsSorting == 1 {
			c := dao.DataTemplate.Columns()
			_, err = dao.DataTemplate.Ctx(ctx).
				Data(g.Map{c.SortNodeKey: in.Key, c.SortDesc: in.IsDesc}).
				Where(c.Id, in.Tid).
				Update()
			if err != nil {
				return err
			}
		}

		return nil
	})

	return
}

func (s *sDataTemplateNode) Edit(ctx context.Context, in *model.DataTemplateNodeEditInput) (err error) {
	var p *entity.DataTemplateNode
	err = dao.DataTemplateNode.Ctx(ctx).Where(dao.DataTemplateNode.Columns().Id, in.Id).Scan(&p)
	if p == nil {
		return gerror.New("数据模型节点不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataTemplateNode
	err = gconv.Scan(in, &param)
	param.UpdateBy = uint(loginUserId)
	param.Id = nil
	if in.SourceId == p.SourceId {
		param.SourceId = nil
		if in.NodeId == p.NodeId {
			param.NodeId = nil
		}
	}

	err = dao.DataTemplateNode.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err = dao.DataTemplateNode.Ctx(ctx).Data(param).Where(dao.DataTemplateNode.Columns().Id, in.Id).Update()
		if err != nil {
			return err
		}

		// 处理排序字段
		sortNode := ""
		sortDesc := 0
		if in.IsSorting == 1 {
			sortNode = p.Key
			sortDesc = in.IsDesc
		}
		c := dao.DataTemplate.Columns()
		_, err = dao.DataTemplate.Ctx(ctx).
			Data(g.Map{c.SortNodeKey: sortNode, c.SortDesc: sortDesc}).
			Where(c.Id, p.Tid).
			WhereNot(c.SortNodeKey, sortNode).
			Update()
		return err
	})

	return
}

func (s *sDataTemplateNode) Del(ctx context.Context, id uint64) (err error) {
	var p *entity.DataTemplateNode
	err = dao.DataTemplateNode.Ctx(ctx).Where(dao.DataTemplateNode.Columns().Id, id).Scan(&p)
	if p == nil {
		return gerror.New("数据模型节点不存在")
	}

	var dt *entity.DataTemplate
	err = dao.DataTemplate.Ctx(ctx).Where(dao.DataTemplate.Columns().Id, p.Tid).Scan(&dt)
	if err != nil {
		return
	}
	if dt != nil && dt.Status != model.DataTemplateStatusOff {
		return gerror.New("数据模型已发布，请先撤回，再删除")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	err = dao.DataTemplateNode.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if dt != nil && dt.DataTable != "" {
			// 表结构已存在，字段须删除处理
			if err = dropTplColumn(ctx, id); err != nil {
				return err
			}
		}

		_, err = dao.DataTemplateNode.Ctx(ctx).
			Data(do.DataTemplateNode{
				DeletedBy: uint(loginUserId),
				DeletedAt: gtime.Now(),
			}).
			Where(dao.DataTemplateNode.Columns().Id, id).
			Unscoped().
			Update()
		if err != nil {
			return err
		}

		// 处理排序字段
		if dt.SortNodeKey == p.Key {
			c := dao.DataTemplate.Columns()
			_, err = dao.DataTemplate.Ctx(ctx).
				Data(g.Map{c.SortNodeKey: "", c.SortDesc: 0}).
				Where(c.Id, p.Tid).
				Update()
			if err != nil {
				return err
			}
		}

		return nil
	})

	return
}

func (s *sDataTemplateNode) List(ctx context.Context, tid uint64) (list []*model.DataTemplateNodeOutput, err error) {
	err = dao.DataTemplateNode.Ctx(ctx).WithAll().OrderAsc(dao.DataTemplateNode.Columns().Id).Where(dao.DataTemplateNode.Columns().Tid, tid).Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	dt, err := service.DataTemplate().Detail(ctx, tid)
	if err != nil || dt == nil {
		return
	}

	for _, v := range list {
		if dt.SortNodeKey == v.Key {
			v.IsSorting = 1
			v.IsDesc = dt.SortDesc
			break
		}
	}

	return
}
