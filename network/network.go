package network

import (
	"context"
	"sagooiot/network/core"
	"sagooiot/network/core/device"
	"sagooiot/network/core/server"
	"sagooiot/network/core/tunnel"
)

var reloadNetWorkFunc = []func(ctx context.Context) error{
	// 开启主题订阅
	core.StartSubscriber,
	device.LoadDevices,
	tunnel.LoadTunnels,
	server.LoadServers,
}

func ReloadNetwork(c context.Context) (err error) {
	for _, f := range reloadNetWorkFunc {
		if err = f(c); err != nil {
			return err
		}
	}
	return
}
