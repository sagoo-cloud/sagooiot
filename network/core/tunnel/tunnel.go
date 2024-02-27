package tunnel

import (
	"fmt"
	"sagooiot/network/core/tunnel/base"
	"sagooiot/network/model"
)

// NewTunnel 创建通道
func NewTunnel(tunnel *model.Tunnel) (base.TunnelInstance, error) {
	var tnl base.TunnelInstance
	switch tunnel.Type {
	case "serial":
		//TODO 等待补全
		//tnl = newTunnelSerial(tunnel)
		break
	case "tcp-client":
		tnl = newTunnelClient(tunnel, "tcp")
		break
	case "udp-client":
		tnl = newTunnelClient(tunnel, "udp")
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", tunnel.Type)
	}
	return tnl, nil
}
