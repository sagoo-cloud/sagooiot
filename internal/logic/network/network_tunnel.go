package network

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/network/core/tunnel"

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

// GetTunnelList 获取列表数据
func (s *sNetworkTunnel) GetTunnelList(ctx context.Context, in *model.GetNetworkTunnelListInput) (total int, out []*model.NetworkTunnelOut, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.NetworkTunnel.Ctx(ctx)

		if in.PaginationInput == nil {
			in.PaginationInput = &model.PaginationInput{}
		}

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

// GetTunnelRunList 获取列表数据
func (s *sNetworkTunnel) GetTunnelRunList(ctx context.Context) (out []*model.NetworkTunnelOut, err error) {
	m := dao.NetworkTunnel.Ctx(ctx)
	err = m.Where(dao.NetworkTunnel.Columns().Status, 1).Scan(&out)
	if err != nil {
		err = gerror.New("获取数据失败")
	}

	return
}

// 获取指定ID数据
func (s *sNetworkTunnel) GetTunnelById(ctx context.Context, id int) (out *model.NetworkTunnelOut, err error) {
	err = dao.NetworkTunnel.Ctx(ctx).Where(dao.NetworkTunnel.Columns().Id, id).Scan(&out)
	return
}

// TODO 这里更改了请求参数，需要确认是否ok
// AddTunnel 添加数据
func (s *sNetworkTunnel) AddTunnel(ctx context.Context, in model.NetworkTunnelAddInput) (id int, err error) {
	//查询通道名称是否存在
	num, err := dao.NetworkTunnel.Ctx(ctx).Where(g.Map{dao.NetworkTunnel.Columns().Name: in.Name}).Count()
	if err != nil {
		return
	}
	if num > 0 {
		err = gerror.New("通道名称已存在")
		return
	}
	err = dao.NetworkTunnel.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		rs, err := dao.NetworkTunnel.Ctx(ctx).Data(do.NetworkTunnel{
			DeptId:    service.Context().GetUserDeptId(ctx),
			ServerId:  in.ServerId,
			Name:      in.Name,
			Types:     in.Types,
			Addr:      in.Addr,
			Remote:    in.Remote,
			Retry:     in.Retry,
			Heartbeat: in.Heartbeat,
			Serial:    in.Serial,
			Protoccol: in.Protocol,
			Status:    in.Status,
			CreatedAt: gtime.Now(),
			Remark:    in.Remark,
		}).Insert()
		if err != nil {
			return
		}
		newId, _ := rs.LastInsertId()
		id = int(newId)

		if err == nil && in.Status == consts.TunnelIsOnLine {
			var networkTunnelEditInput model.NetworkTunnelEditInput
			if err = dao.NetworkTunnel.Ctx(ctx).Where(dao.NetworkTunnel.Columns().Name, in.Name).Scan(&networkTunnelEditInput); err != nil {
				return
			} else {
				err = tunnel.LoadTunnel(ctx, networkTunnelEditInput.Id)
				return
			}
		}

		return
	})
	if err != nil {
		return 0, err
	}
	return
}

// EditTunnel 修改数据
func (s *sNetworkTunnel) EditTunnel(ctx context.Context, in model.NetworkTunnelEditInput) (err error) {
	var netWorkTunnel *entity.NetworkTunnel
	err = dao.NetworkTunnel.Ctx(ctx).Where(dao.NetworkTunnel.Columns().Id, in.Id).Scan(&netWorkTunnel)
	if err != nil {
		return err
	}
	if netWorkTunnel == nil {
		return gerror.New("ID错误")
	}

	//查询通道名称是否存在
	num, err := dao.NetworkTunnel.Ctx(ctx).Where(g.Map{dao.NetworkTunnel.Columns().Name: in.Name}).WhereNot(dao.NetworkTunnel.Columns().Id, in.Id).Count()
	if err != nil {
		return
	}
	if num > 0 {
		err = gerror.New("通道名称已存在")
		return
	}

	var param do.NetworkTunnel
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.Id = nil
	if in.Remote == "" {
		param.Remote = nil
	}

	_, err = dao.NetworkTunnel.Ctx(ctx).Data(param).Where(dao.NetworkTunnel.Columns().Id, in.Id).Update()
	if err != nil {
		return
	}
	if err = tunnel.RemoveTunnel(in.Id); err != nil {
		return err
	}
	if in.Status == consts.TunnelIsOnLine {
		if err = tunnel.LoadTunnel(ctx, in.Id); err != nil {
			return err
		}
	}
	return
}

// DeleteTunnel 删除数据
func (s *sNetworkTunnel) DeleteTunnel(ctx context.Context, ids []int) (err error) {
	for _, id := range ids {
		var netWorkTunnel *entity.NetworkTunnel
		err = dao.NetworkTunnel.Ctx(ctx).Where(dao.NetworkTunnel.Columns().Id, id).Scan(&netWorkTunnel)
		if err != nil {
			return err
		}
		if netWorkTunnel == nil {
			return gerror.New("ID错误")
		}
	}

	_, err = dao.NetworkTunnel.Ctx(ctx).Delete(dao.NetworkTunnel.Columns().Id+" in (?)", ids)
	//TODO 这里需要注意中间删除失败的情况
	if err == nil {
		for _, node := range ids {
			if err = tunnel.RemoveTunnel(node); err != nil {
				return err
			}
		}
	}
	return
}

// SetTunnelStatus 修改状态数据
func (s *sNetworkTunnel) SetTunnelStatus(ctx context.Context, id, status int) (err error) {

	var netWorkTunnel *entity.NetworkTunnel
	err = dao.NetworkTunnel.Ctx(ctx).Where(dao.NetworkTunnel.Columns().Id, id).Scan(&netWorkTunnel)
	if err != nil {
		return err
	}
	if netWorkTunnel == nil {
		return gerror.New("ID错误")
	}

	var data = g.Map{
		dao.NetworkTunnel.Columns().Status: status,
	}

	//TODO 这儿里还需要进行通道的启用处理，启用成功更新数据状态

	_, err = dao.NetworkTunnel.Ctx(ctx).Data(data).Where(dao.NetworkTunnel.Columns().Id, id).Update()
	if err != nil {
		return
	}
	if status == 0 {
		if err = tunnel.RemoveTunnel(id); err != nil {
			return err
		}
	} else {
		if err = tunnel.LoadTunnel(ctx, id); err != nil {
			return err
		}
	}
	return
}
