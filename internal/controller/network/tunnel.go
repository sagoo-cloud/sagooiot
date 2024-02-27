package network

import (
	"context"
	"sagooiot/api/v1/network"
	"sagooiot/internal/model"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

var Tunnel = cNetworkTunnel{}

type cNetworkTunnel struct{}

// 获取列表
func (u *cNetworkTunnel) GetNetworkTunnelList(ctx context.Context, req *network.GetNetworkTunnelListReq) (res *network.GetNetworkTunnelListRes, err error) {
	var input *model.GetNetworkTunnelListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, out, err := service.NetworkTunnel().GetTunnelList(ctx, input)
	if err != nil {
		return
	}
	res = new(network.GetNetworkTunnelListRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.Data); err != nil {
			return
		}
	}
	return
}

// 获取指定ID数据
func (u *cNetworkTunnel) GetTunnelById(ctx context.Context, req *network.GetNetworkTunnelByIdReq) (res *network.GetNetworkTunnelByIdRes, err error) {
	out, err := service.NetworkTunnel().GetTunnelById(ctx, req.Id)
	if err != nil {
		return
	}
	var data *model.NetworkTunnelRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &network.GetNetworkTunnelByIdRes{
		Data: data,
	}
	return
}

// 添加数据
func (u *cNetworkTunnel) AddTunnel(ctx context.Context, req *network.AddNetworkTunnelReq) (res *network.AddNetworkTunnelRes, err error) {

	var data = model.NetworkTunnelAddInput{}
	err = gconv.Scan(req, &data)
	_, err = service.NetworkTunnel().AddTunnel(ctx, data)
	return
}

// 修改数据
func (u *cNetworkTunnel) EditTunnel(ctx context.Context, req *network.EditNetworkTunnelReq) (res *network.EditNetworkTunnelRes, err error) {
	var data = model.NetworkTunnelEditInput{}
	err = gconv.Scan(req, &data)
	err = service.NetworkTunnel().EditTunnel(ctx, data)
	return
}

// 删除数据
func (u *cNetworkTunnel) DeleteTunnel(ctx context.Context, req *network.DeleteNetworkTunnelReq) (res *network.DeleteNetworkTunnelRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.NetworkTunnel().DeleteTunnel(ctx, req.Ids)
	return
}

// 修改数据
func (u *cNetworkTunnel) SetTunnelStatus(ctx context.Context, req *network.SetNetworkTunnelStatusReq) (res *network.SetNetworkTunnelStatusRes, err error) {
	err = service.NetworkTunnel().SetTunnelStatus(ctx, req.Id, req.Status)
	return
}
