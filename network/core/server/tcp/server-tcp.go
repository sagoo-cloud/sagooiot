package tcp

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"net"
	"sagooiot/network/core/server/common"
	"sagooiot/network/core/tunnel/base"
	"sagooiot/network/model"
)

type ServerTCP struct {
	server *model.Server

	children map[string]*ServerTcpTunnel

	listener *net.TCPListener

	running bool
}

func NewServerTCP(server *model.Server) *ServerTCP {
	svr := &ServerTCP{
		server:   server,
		children: make(map[string]*ServerTcpTunnel),
	}
	return svr
}

func (server *ServerTCP) Open(ctx context.Context) error {
	if server.running {
		return errors.New("server is opened")
	}
	common.ServerOpenAction(server.server.Id)

	addr, err := net.ResolveTCPAddr("tcp4", common.ResolvePort(server.server.Addr))
	if err != nil {
		return err
	}
	server.listener, err = net.ListenTCP("tcp4", addr)
	if err != nil {
		return err
	}

	server.running = true
	go func() {
		for {
			c, acceptErr := server.listener.AcceptTCP()
			if acceptErr != nil {
				g.Log().Errorf(ctx, "accept tcp  error: %s", acceptErr.Error())
				break
			}

			buf := make([]byte, 1024)
			n, readErr := c.Read(buf)
			if readErr != nil {
				g.Log().Errorf(ctx, "read tcp  error: %v,local_addr:%s remote_addr:%s", readErr, c.LocalAddr().String(), c.RemoteAddr().String())
				_ = c.Close()
				continue
			}
			data := buf[:n]
			deviceKey, checkIsOk := server.server.Register.Check(data)
			if !checkIsOk {
				g.Log().Errorf(ctx, "register check not right,check_data:%s ,local_addr:%s remote_addr:%s", string(data), c.LocalAddr().String(), c.RemoteAddr().String())
				_ = c.Close()
				continue
			}
			tnl, tnlErr := newServerTcpTunnel(ctx, server.server.Id, deviceKey, c)
			if tnlErr != nil {
				g.Log().Errorf(ctx, "new tcp tunnel error: %v,local_addr:%s remote_addr:%s", tnlErr, c.LocalAddr().String(), c.RemoteAddr().String())
				continue
			}
			server.children[tnl.TunnelId] = tnl
			go func() {
				tnl.receive(ctx)
				server.RemoveTunnel(tnl.TunnelId)
			}()
			common.ServerTunnelAction(ctx, server.server.Id, deviceKey)
		}
		server.running = false
	}()
	return nil
}

func (server *ServerTCP) Close() (err error) {
	common.ServerCloseAction(server.server.Id)
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

func (server *ServerTCP) GetTunnel(id string) base.TunnelInstance {
	return server.children[id]
}

func (server *ServerTCP) RemoveTunnel(id string) {
	delete(server.children, id)
}

func (server *ServerTCP) Running() bool {
	return server.running
}
