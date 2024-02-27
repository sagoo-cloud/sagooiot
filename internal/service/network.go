// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sagooiot/internal/model"
)

type (
	INetworkServer interface {
		// GetServerList 获取列表数据
		GetServerList(ctx context.Context, in *model.GetNetworkServerListInput) (total int, out []*model.NetworkServerOut, err error)
		// GetServerRunList 获取可运行的服务列表数据
		GetServerRunList(ctx context.Context) (list []*model.NetworkServerRes, err error)
		// GetServerById 获取指定ID数据
		GetServerById(ctx context.Context, id int) (out *model.NetworkServerOut, err error)
		// AddServer 添加数据 todo 需要处理
		AddServer(ctx context.Context, in model.NetworkServerAddInput) (err error)
		// EditServer 修改数据 todo 需要处理
		EditServer(ctx context.Context, in model.NetworkServerEditInput) (err error)
		// 删除数据
		// todo 需要处理
		DeleteServer(ctx context.Context, ids []int) (err error)
		// SetServerStatus 修改状态数据 todo 需要处理
		SetServerStatus(ctx context.Context, id, status int) (err error)
	}
	INetworkTunnel interface {
		// GetTunnelList 获取列表数据
		GetTunnelList(ctx context.Context, in *model.GetNetworkTunnelListInput) (total int, out []*model.NetworkTunnelOut, err error)
		// GetTunnelRunList 获取列表数据
		GetTunnelRunList(ctx context.Context) (out []*model.NetworkTunnelOut, err error)
		// 获取指定ID数据
		GetTunnelById(ctx context.Context, id int) (out *model.NetworkTunnelOut, err error)
		// TODO 这里更改了请求参数，需要确认是否ok
		// AddTunnel 添加数据
		AddTunnel(ctx context.Context, in model.NetworkTunnelAddInput) (id int, err error)
		// EditTunnel 修改数据
		EditTunnel(ctx context.Context, in model.NetworkTunnelEditInput) (err error)
		// DeleteTunnel 删除数据
		DeleteTunnel(ctx context.Context, ids []int) (err error)
		// SetTunnelStatus 修改状态数据
		SetTunnelStatus(ctx context.Context, id, status int) (err error)
	}
)

var (
	localNetworkServer INetworkServer
	localNetworkTunnel INetworkTunnel
)

func NetworkServer() INetworkServer {
	if localNetworkServer == nil {
		panic("implement not found for interface INetworkServer, forgot register?")
	}
	return localNetworkServer
}

func RegisterNetworkServer(i INetworkServer) {
	localNetworkServer = i
}

func NetworkTunnel() INetworkTunnel {
	if localNetworkTunnel == nil {
		panic("implement not found for interface INetworkTunnel, forgot register?")
	}
	return localNetworkTunnel
}

func RegisterNetworkTunnel(i INetworkTunnel) {
	localNetworkTunnel = i
}
