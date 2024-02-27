package tcp

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"net"
	"sagooiot/internal/consts"
	"sagooiot/network/core/server/base"
	"sagooiot/network/core/tunnel"
	"sagooiot/network/core/tunnel/action"
)

// ServerTcpTunnel 网络连接
type ServerTcpTunnel struct {
	serverId  int
	deviceKey string
	*tunnel.TunnelBase
}

func newServerTcpTunnel(ctx context.Context, serverId int, deviceKey string, conn net.Conn) (*ServerTcpTunnel, error) {
	baseServerTunnelInfo := base.ServerTunnel{
		ServerId:   serverId,
		DeviceKey:  deviceKey,
		Type:       "tcp",
		Status:     consts.TunnelIsOnLine,
		LocalAddr:  conn.LocalAddr().String(),
		RemoteAddr: conn.RemoteAddr().String(),
		Remark:     "",
	}
	var err error
	baseServerTunnelInfo.TunnelId, err = base.AddOrEditServerTunnel(ctx, baseServerTunnelInfo)
	if err != nil {
		return nil, err
	}
	g.Log().Debug(ctx, "newServerTcpTunnel", serverId, deviceKey, conn.LocalAddr().String(), conn.RemoteAddr().String())
	tunnelBase := tunnel.TunnelBase{
		TunnelId: baseServerTunnelInfo.TunnelId,
		Link:     conn,
		ServerId: serverId,
	}
	tunnelBase.SetRunning(true)
	tunnelBase.SetOnline(true)
	return &ServerTcpTunnel{serverId: serverId, deviceKey: deviceKey, TunnelBase: &tunnelBase}, nil
}

func (l *ServerTcpTunnel) Open(ctx context.Context) error {
	return errors.New("ServerTcpTunnel cannot open")
}

func (l *ServerTcpTunnel) receive(ctx context.Context) {

	if err := action.TunnelOnlineAction(ctx, l.serverId, l.TunnelId, l.deviceKey); err != nil {
		g.Log().Errorf(ctx, "tunnel online error: %v", err)
		return
	}
	buf := make([]byte, 1024)
	for {
		n, err := l.Link.Read(buf)
		if err != nil {
			l.OnClose()
			break
		}
		if n == 0 {
			continue
		}
		if l.GetPipe() != nil {
			_, err = l.GetPipe().Write(buf[:n])
			if err != nil {
				l.SetPipe(nil)
			} else {
				continue
			}
		}
		go l.TunnelBase.ReadData(ctx, l.deviceKey, buf[:n])
	}
	l.SetRunning(false)
	l.SetOnline(false)

	action.TunnelOfflineAction(ctx, l.serverId, l.TunnelId, l.deviceKey)

}
