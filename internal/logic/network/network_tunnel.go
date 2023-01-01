package network

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/core"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sNetworkTunnel struct{}

func sNetworkTunnelNew() *sNetworkTunnel {
	return &sNetworkTunnel{}
}
func init() {
	service.RegisterNetworkTunnel(sNetworkTunnelNew())
}

const (
	TunnelIsOffline = iota
	TunnelIsOnLine
)

// 获取列表数据
func (s *sNetworkTunnel) GetTunnelList(ctx context.Context, in *model.GetNetworkTunnelListInput) (total int, out []*model.NetworkTunnelOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NetworkTunnel.Ctx(ctx)

		if in.ServiceId > 0 {
			m = m.Where(dao.NetworkTunnel.Columns().ServerId, in.ServiceId)
		}
		if in.DeviceKey != "" {
			m = m.Where(dao.NetworkTunnel.Columns().DeviceKey, in.DeviceKey)
		}
		if in.KeyWord != "" {
			m = m.WhereLike(dao.NetworkTunnel.Columns().Name, "%"+in.KeyWord+"%")
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

// 获取列表数据
func (s *sNetworkTunnel) GetTunnelRunList(ctx context.Context) (out []*model.NetworkTunnelOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NetworkTunnel.Ctx(ctx)
		err = m.Where(dao.NetworkTunnel.Columns().Status, 1).Scan(&out)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
	})
	return
}

// 获取指定ID数据
func (s *sNetworkTunnel) GetTunnelById(ctx context.Context, id int) (out *model.NetworkTunnelOut, err error) {
	err = dao.NetworkTunnel.Ctx(ctx).Where("id", id).Scan(&out)
	return
}

// TODO 这里更改了请求参数，需要确认是否ok
// 添加数据
func (s *sNetworkTunnel) AddTunnel(ctx context.Context, in model.NetworkTunnelAddInput) (id int, err error) {
	rs, err := dao.NetworkTunnel.Ctx(ctx).Insert(in)
	if err != nil {
		return
	}
	newId, _ := rs.LastInsertId()
	id = int(newId)

	if err == nil && in.Status == TunnelIsOnLine {
		var networkTunnelEditInput model.NetworkTunnelEditInput
		if err = dao.NetworkTunnel.Ctx(ctx).Where("name", in.Name).Scan(&networkTunnelEditInput); err != nil {
			return
		} else {
			err = core.LoadTunnel(ctx, networkTunnelEditInput.Id)
			return
		}
	}
	return
}

// 修改数据
func (s *sNetworkTunnel) EditTunnel(ctx context.Context, in model.NetworkTunnelEditInput) (err error) {
	var param do.NetworkTunnel
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.Id = nil
	if in.Remote == "" {
		param.Remote = nil
	}

	_, err = dao.NetworkTunnel.Ctx(ctx).FieldsEx(dao.NetworkTunnel.Columns().Id).Where(dao.NetworkTunnel.Columns().Id, in.Id).Update(param)
	if err != nil {
		return
	}
	if err = core.RemoveTunnel(in.Id); err != nil {
		return err
	}
	if in.Status == TunnelIsOnLine {
		if err = core.LoadTunnel(ctx, in.Id); err != nil {
			return err
		}
	}
	return
}

// 删除数据
func (s *sNetworkTunnel) DeleteTunnel(ctx context.Context, ids []int) (err error) {
	_, err = dao.NetworkTunnel.Ctx(ctx).Delete(dao.NetworkTunnel.Columns().Id+" in (?)", ids)
	//TODO 这里需要注意中间删除失败的情况
	if err == nil {
		for _, node := range ids {
			if err = core.RemoveTunnel(node); err != nil {
				return err
			}
		}
	}
	return
}

// 修改状态数据
func (s *sNetworkTunnel) SetTunnelStatus(ctx context.Context, id, status int) (err error) {
	var data = g.Map{
		dao.NetworkTunnel.Columns().Status: status,
	}

	//TODO 这儿里还需要进行通道的启用处理，启用成功更新数据状态

	_, err = dao.NetworkTunnel.Ctx(ctx).Where(dao.NetworkTunnel.Columns().Id, id).Update(data)
	if err != nil {
		return
	}
	if status == 0 {
		if err = core.LoadTunnel(ctx, id); err != nil {
			return err
		}
	} else {
		if err = core.RemoveTunnel(id); err != nil {
			return err
		}
	}
	return
}
