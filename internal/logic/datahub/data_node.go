package datahub

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sDataNode struct{}

func init() {
	service.RegisterDataNode(dataNodeNew())
}

func dataNodeNew() *sDataNode {
	return &sDataNode{}
}

func (s *sDataNode) Add(ctx context.Context, in *model.DataNodeAddInput) (err error) {
	id, _ := dao.DataNode.Ctx(ctx).
		Fields(dao.DataNode.Columns().NodeId).
		Where(dao.DataNode.Columns().SourceId, in.SourceId).
		Where(dao.DataNode.Columns().Key, in.Key).
		Value()
	if id.Int64() > 0 {
		return gerror.New("数据节点标识重复")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataNode
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)

	if in.Rule != nil {
		rule, err := json.Marshal(in.Rule)
		if err != nil {
			return gerror.New("规则配置格式错误")
		}
		param.Rule = rule
	}

	err = dao.DataNode.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		rs, err := dao.DataNode.Ctx(ctx).Data(param).Insert()
		if err != nil {
			return err
		}

		nodeId, _ := rs.LastInsertId()

		dataSource, _ := service.DataSource().Detail(ctx, in.SourceId)
		if dataSource != nil && dataSource.DataTable != "" {
			// 表结构已存在，字段新增
			err = addColumn(ctx, uint64(nodeId))
			if err != nil {
				return err
			}
		}

		return nil
	})

	return
}

func (s *sDataNode) Edit(ctx context.Context, in *model.DataNodeEditInput) (err error) {
	id, _ := dao.DataNode.Ctx(ctx).
		Fields(dao.DataNode.Columns().NodeId).
		Where(dao.DataNode.Columns().NodeId, in.NodeId).
		Value()
	if id.Int64() == 0 {
		return gerror.New("数据节点不存在")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.DataNode
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.NodeId = nil

	_, err = dao.DataNode.Ctx(ctx).Data(param).Where(dao.DataNode.Columns().NodeId, in.NodeId).Update()

	return
}

func (s *sDataNode) Del(ctx context.Context, nodeId uint64) (err error) {
	var p *entity.DataNode
	err = dao.DataNode.Ctx(ctx).Where(dao.DataNode.Columns().NodeId, nodeId).Scan(&p)
	if p == nil {
		return gerror.New("数据节点不存在")
	}

	var ds *entity.DataSource
	err = dao.DataSource.Ctx(ctx).Where(dao.DataSource.Columns().SourceId, p.SourceId).Scan(&ds)
	if err != nil {
		return
	}
	if ds != nil && ds.Status != model.DataSourceStatusOff {
		return gerror.New("数据源已发布，请先撤回，再删除")
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	err = dao.DataNode.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		if ds != nil && ds.DataTable != "" {
			// 表结构已存在，字段须删除处理
			if err = dropColumn(ctx, nodeId); err != nil {
				return err
			}
		}

		_, err = dao.DataNode.Ctx(ctx).
			Data(do.DataNode{
				DeletedBy: uint(loginUserId),
				DeletedAt: gtime.Now(),
			}).
			Where(dao.DataNode.Columns().NodeId, nodeId).
			Unscoped().
			Update()
		return err
	})

	return
}

func (s *sDataNode) List(ctx context.Context, sourceId uint64) (list []*model.DataNodeOutput, err error) {
	var p []*entity.DataNode
	err = dao.DataNode.Ctx(ctx).OrderAsc(dao.DataNode.Columns().NodeId).Where(dao.DataNode.Columns().SourceId, sourceId).Scan(&p)
	if err != nil || p == nil {
		return
	}

	list = make([]*model.DataNodeOutput, len(p))
	for i, v := range p {
		// 规则配置
		var rule []*model.DataSourceRule
		if v.Rule != "" {
			j, _ := gjson.DecodeToJson(v.Rule)
			if err = j.Scan(&rule); err != nil {
				return nil, err
			}
		}

		out := new(model.DataNodeOutput)
		out.DataNode = v
		out.NodeRule = rule

		list[i] = out
	}

	return
}

// 详情
func (s *sDataNode) Detail(ctx context.Context, nodeId uint64) (out *model.DataNodeOutput, err error) {
	var p *entity.DataNode
	err = dao.DataNode.Ctx(ctx).Where(dao.DataNode.Columns().NodeId, nodeId).Scan(&p)
	if err != nil || p == nil {
		return
	}

	// 规则配置
	var rule []*model.DataSourceRule
	if p.Rule != "" {
		j, _ := gjson.DecodeToJson(p.Rule)
		if err = j.Scan(&rule); err != nil {
			return nil, err
		}
	}

	out = new(model.DataNodeOutput)
	out.DataNode = p
	out.NodeRule = rule

	return
}
