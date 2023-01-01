package core

import (
	"context"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/network/core/tunnelinstance"
	"github.com/sagoo-cloud/sagooiot/network/model"
)

// Server 通道
type ServerInstance interface {
	Open(ctx context.Context) error
	Close() error
	GetTunnel(id int) tunnelinstance.TunnelInstance
	RemoveTunnel(id int)
	Running() bool
}

// NewServer 创建通道
func NewServer(server *model.Server) (ServerInstance, error) {
	var svr ServerInstance
	var err error
	switch server.Type {
	case "tcp":
		svr = newServerTCP(server)
		break
	case "udp":
		svr = newServerUDP(server)
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", server.Type)
	}

	return svr, err
}
