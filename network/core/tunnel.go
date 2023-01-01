package core

import (
	"fmt"
	"github.com/sagoo-cloud/sagooiot/network/core/tunnelinstance"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"strings"
)

// NewTunnel 创建通道
func NewTunnel(tunnel *model.Tunnel) (tunnelinstance.TunnelInstance, error) {
	var tnl tunnelinstance.TunnelInstance
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

func resolvePort(addr string) string {
	if strings.IndexByte(addr, ':') == -1 {
		return ":" + addr
	}
	return addr
}
