package tunnel

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/service"
	"sagooiot/network/core/mapper"
	"sagooiot/network/core/server/base"
	base2 "sagooiot/network/core/tunnel/base"
	"sagooiot/network/model"
	"sync"
)

var allTunnels sync.Map

type Server struct {
	model.Server
	Instance base.ServerInstance
}

type Tunnel struct {
	Instance base2.TunnelInstance
}

func startTunnel(ctx context.Context, tunnel *model.Tunnel) error {
	tnl, err := NewTunnel(tunnel)
	if err != nil {
		g.Log().Error(ctx, err.Error())
		return nil
	}
	if tnl == nil {
		return fmt.Errorf("new tunnel: %+#v tunnel failed", tunnel)
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
		t := mapper.Tunnel(ctx, *node)
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
	tunnel := mapper.Tunnel(ctx, *tunnelInfo)
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
