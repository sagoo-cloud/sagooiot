package network

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/network"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var Server = cNetworkServer{}

type cNetworkServer struct{}

// 获取列表
func (u *cNetworkServer) GetNetworkServerList(ctx context.Context, req *network.GetNetworkServerListReq) (res *network.GetNetworkServerListRes, err error) {
	var input *model.GetNetworkServerListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	total, out, err := service.NetworkServer().GetServerList(ctx, input)
	if err != nil {
		return
	}
	res = new(network.GetNetworkServerListRes)
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
func (u *cNetworkServer) GetServerById(ctx context.Context, req *network.GetNetworkServerByIdReq) (res *network.GetNetworkServerByIdRes, err error) {
	out, err := service.NetworkServer().GetServerById(ctx, req.Id)
	var data *model.NetworkServerRes
	if out != nil {
		if err = gconv.Scan(out, &data); err != nil {
			return
		}
	}
	res = &network.GetNetworkServerByIdRes{
		Data: data,
	}
	return
}

// 添加数据
func (u *cNetworkServer) AddServer(ctx context.Context, req *network.AddNetworkServerReq) (res *network.AddNetworkServerRes, err error) {
	var data = model.NetworkServerAddInput{}
	err = gconv.Scan(req, &data)
	err = service.NetworkServer().AddServer(ctx, data)
	return
}

// 修改数据
func (u *cNetworkServer) EditServer(ctx context.Context, req *network.EditNetworkServerReq) (res *network.EditNetworkServerRes, err error) {
	var data = model.NetworkServerEditInput{}
	err = gconv.Scan(req, &data)
	err = service.NetworkServer().EditServer(ctx, data)
	return
}

// 删除数据
func (u *cNetworkServer) DeleteServer(ctx context.Context, req *network.DeleteNetworkServerReq) (res *network.DeleteNetworkServerRes, err error) {
	if len(req.Ids) == 0 {
		err = gerror.New("ID参数错误")
	}
	err = service.NetworkServer().DeleteServer(ctx, req.Ids)
	return
}

// 修改数据
func (u *cNetworkServer) SetServerStatus(ctx context.Context, req *network.SetNetworkServerStatusReq) (res *network.SetNetworkServerStatusRes, err error) {
	err = service.NetworkServer().SetServerStatus(ctx, req.Id, req.Status)
	return
}
