package monitorops

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sMonitoropsRemoteconf struct{}

func sMonitoropsRemoteconfNew() *sMonitoropsRemoteconf {
	return &sMonitoropsRemoteconf{}
}
func init() {
	service.RegisterMonitoropsRemoteconf(sMonitoropsRemoteconfNew())
}

// GetRemoteconfList 获取列表数据
func (s *sMonitoropsRemoteconf) GetRemoteconfList(ctx context.Context, in *model.GetRemoteconfListInput) (total, page int, list []*model.RemoteconfOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Remoteconf.Ctx(ctx)
		c := dao.Remoteconf.Columns()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		page = in.PageNum
		if in.PageSize == 0 {
			in.PageSize = consts.PageSize
		}
		if in.ProductKey != "" {
			m = m.Where(c.ProductKey, in.ProductKey)
		}
		total, err = m.Count()
		err = m.Page(page, in.PageSize).Order("utc_create desc").Scan(&list)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
		for i := range list {
			list[i].ConfigNumber = fmt.Sprintf("%02d", len(list)-i)
		}
	})
	return
}

// GetRemoteconfById 获取指定ID数据
func (s *sMonitoropsRemoteconf) GetRemoteconfById(ctx context.Context, id int) (out *model.RemoteconfOutput, err error) {
	err = dao.Remoteconf.Ctx(ctx).Where(dao.Remoteconf.Columns().Id, id).Scan(&out)
	return
}

// AddRemoteconf 添加数据
func (s *sMonitoropsRemoteconf) AddRemoteconf(ctx context.Context, in model.RemoteconfAddInput) (err error) {
	var p []*entity.DevProduct
	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, in.ProductKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}
	var param *do.Remoteconf
	err = gconv.Scan(in, &param)
	if err != nil {
		glog.Error(ctx, err)
		return
	}

	param.Id = guid.S()
	param.UtcCreate = gtime.Now().UTC()
	param.GmtCreate = gtime.Now().Format("Y/m/d H:i:s")

	_, err = dao.Remoteconf.Ctx(ctx).Insert(param)
	return
}

// EditRemoteconf 修改数据
func (s *sMonitoropsRemoteconf) EditRemoteconf(ctx context.Context, in model.RemoteconfEditInput) (err error) {
	dao.Remoteconf.Ctx(ctx).FieldsEx(dao.Remoteconf.Columns().Id).Where(dao.Remoteconf.Columns().Id, in.Id).Update(in)
	return
}

// DeleteRemoteconf 删除数据
func (s *sMonitoropsRemoteconf) DeleteRemoteconf(ctx context.Context, Ids []int) (err error) {
	_, err = dao.Remoteconf.Ctx(ctx).Delete(dao.Remoteconf.Columns().Id+" in (?)", Ids)
	return
}
