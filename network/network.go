package network

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/network/core"
)

func ReloadNetwork(c context.Context) (err error) {
	for _, f := range []func(ctx context.Context) error{core.StartSubscriber, core.LoadDevices, core.LoadTunnels, core.LoadServers} {
		if err = f(c); err != nil {
			return err
		}
	}
	return
}
