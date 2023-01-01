package core

import (
	"context"
	"errors"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"net"
)

// ServerTcpTunnel 网络连接
type ServerTcpTunnel struct {
	serverId  int
	deviceKey string
	tunnelBase
}

func newServerTcpTunnel(serverId, tunnelId int, deviceKey string, conn net.Conn) *ServerTcpTunnel {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOnline, tunnelId), nil)
	return &ServerTcpTunnel{serverId: serverId, deviceKey: deviceKey, tunnelBase: tunnelBase{
		tunnelId: tunnelId,
		link:     conn,
	}}
}

func (l *ServerTcpTunnel) Open(ctx context.Context) error {
	return errors.New("ServerTcpTunnel cannot open")
}

func (l *ServerTcpTunnel) receive(ctx context.Context) {
	l.running = true
	l.online = true
	TunnelOnlineAction(ctx, l.tunnelId, l.deviceKey)

	buf := make([]byte, 1024)
	for {
		n, err := l.link.Read(buf)
		if err != nil {
			l.onClose()
			break
		}
		if n == 0 {
			continue
		}
		if l.pipe != nil {
			_, err = l.pipe.Write(buf[:n])
			if err != nil {
				l.pipe = nil
			} else {
				continue
			}
		}
		go l.tunnelBase.ReadData(ctx, l.deviceKey, buf[:n])
	}
	l.running = false
	l.online = false
	TunnelOfflineAction(ctx, l.serverId, l.tunnelId)
}
