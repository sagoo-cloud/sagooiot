package network

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/core"
)

type sNetworkServer struct{}

func sNetworkServerNew() *sNetworkServer {
	return &sNetworkServer{}
}
func init() {
	service.RegisterNetworkServer(sNetworkServerNew())
}

const (
	ServerStatusOnline  = "online"
	ServerStatusOffline = "offline"
)

// GetServerList 获取列表数据
func (s *sNetworkServer) GetServerList(ctx context.Context, in *model.GetNetworkServerListInput) (total int, out []*model.NetworkServerOut, err error) {
	//g.Log().Debug(ctx, in)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NetworkServer.Ctx(ctx)

		if in.KeyWord != "" {
			m = m.WhereLike(dao.NetworkServer.Columns().Name, "%"+in.KeyWord+"%")
		}

		total, err = m.Count()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		if in.PageNum == 0 {
			in.PageNum = 1
		}
		if in.PageSize == 0 {
			in.PageSize = consts.PageSize
		}
		err = m.Page(in.PageNum, in.PageSize).Order("created_at desc").Scan(&out)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})

	return
}

// GetServerRunList 获取可运行的服务列表数据
func (s *sNetworkServer) GetServerRunList(ctx context.Context) (list []*model.NetworkServerRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NetworkServer.Ctx(ctx)
		err = m.Where(dao.NetworkServer.Columns().Status, 1).Scan(&list)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})
	return
}

// GetServerById 获取指定ID数据
func (s *sNetworkServer) GetServerById(ctx context.Context, id int) (out *model.NetworkServerOut, err error) {
	err = dao.NetworkServer.Ctx(ctx).Where("id", id).Scan(&out)
	return
}

// AddServer 添加数据 todo 需要处理
func (s *sNetworkServer) AddServer(ctx context.Context, in model.NetworkServerAddInput) (err error) {
	insertResult, insertResultErr := dao.NetworkServer.Ctx(ctx).Insert(in)
	if insertResultErr != nil {
		return insertResultErr
	}
	lastId, lastIdErr := insertResult.LastInsertId()
	if lastIdErr != nil {
		return lastIdErr
	}
	if err == nil && in.Status == consts.ServerStatusOnline {
		return core.LoadServer(ctx, int(lastId))
	}
	return
}

// EditServer 修改数据 todo 需要处理
func (s *sNetworkServer) EditServer(ctx context.Context, in model.NetworkServerEditInput) (err error) {
	_, err = dao.NetworkServer.Ctx(ctx).FieldsEx(dao.NetworkServer.Columns().Id, dao.NetworkServer.Columns().CreateBy).Where(dao.NetworkServer.Columns().Id, in.Id).Update(in)
	if err = core.RemoveServer(in.Id); err != nil {
		return err
	}
	if in.Status == 1 {
		if err = core.LoadServer(ctx, in.Id); err != nil {
			return err
		}
	}
	return
}

// 删除数据
// todo 需要处理
func (s *sNetworkServer) DeleteServer(ctx context.Context, ids []int) (err error) {
	_, err = dao.NetworkServer.Ctx(ctx).Delete(dao.NetworkServer.Columns().Id+" in (?)", ids)
	if err == nil {
		for _, node := range ids {
			if err = core.RemoveServer(node); err != nil {
				return err
			}
		}
	}
	return
}

// SetServerStatus 修改状态数据 todo 需要处理
func (s *sNetworkServer) SetServerStatus(ctx context.Context, id, status int) (err error) {

	var data = g.Map{
		dao.NetworkServer.Columns().Status: status,
	}
	_, err = dao.NetworkServer.Ctx(ctx).Where(dao.NetworkServer.Columns().Id, id).Update(data)
	if err != nil {
		return err
	}
	if status == 1 {
		if err = core.LoadServer(ctx, id); err != nil {
			return err
		}
	} else {
		if err = core.RemoveServer(id); err != nil {
			return err
		}
	}
	return
}
