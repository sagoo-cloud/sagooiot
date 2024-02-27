package server

import (
	"fmt"
	"sagooiot/network/core/server/base"
	"sagooiot/network/core/server/tcp"
	"sagooiot/network/model"
)

// NewServer 创建通道
func NewServer(server *model.Server) (base.ServerInstance, error) {
	var svr base.ServerInstance
	var err error
	switch server.Type {
	case "tcp":
		svr = tcp.NewServerTCP(server)
		break
	case "udp":
		break
	case "http":
		break
	case "websocket":
		break
	default:
		return nil, fmt.Errorf("Unsupport type %s ", server.Type)
	}

	return svr, err
}
