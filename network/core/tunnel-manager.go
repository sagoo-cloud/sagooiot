package core

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/core/tunnelinstance"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"sync"
)

var allTunnels sync.Map

type Server struct {
	model.Server
	Instance ServerInstance
}

type Tunnel struct {
	Instance tunnelinstance.TunnelInstance
}

func startTunnel(ctx context.Context, tunnel *model.Tunnel) error {
	tnl, err := NewTunnel(tunnel)
	if err != nil {
		// log.Error(err)
		return err
	}
	return tnl.Open(ctx)
}

func LoadTunnels(ctx context.Context) error {
	allTunnelModels := make([]*model.Tunnel, 0)

	tunnelRunList, err := service.NetworkTunnel().GetTunnelRunList(ctx)
	if err != nil {
		return err
	}
	for _, node := range tunnelRunList {
		t := mapperTunnel(ctx, *node)
		allTunnelModels = append(allTunnelModels, &t)
	}

	for index := range allTunnelModels {
		go func(tunnel *model.Tunnel) {
			err = startTunnel(ctx, tunnel)
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}(allTunnelModels[index])
	}
	return nil
}

func LoadTunnel(ctx context.Context, id int) error {

	tunnelInfo, err := service.NetworkTunnel().GetTunnelById(ctx, id)
	if err != nil {
		return err
	}
	tunnel := mapperTunnel(ctx, *tunnelInfo)
	if tunnel.Disabled {
		return nil // TODO error ??
	}
	err = startTunnel(ctx, &tunnel)
	if err != nil {
		return err
	}
	return nil
}

func GetTunnel(id int) *Tunnel {
	d, ok := allTunnels.Load(id)
	if ok {
		return d.(*Tunnel)
	}
	return nil
}

func RemoveTunnel(id int) error {
	d, ok := allTunnels.LoadAndDelete(id)
	if ok {
		lnk := d.(*Tunnel)
		return lnk.Instance.Close()
	}
	return nil // error
}
