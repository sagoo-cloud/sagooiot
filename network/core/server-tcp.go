package core

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	interlModel "github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/core/tunnelinstance"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"net"
)

type ServerTCP struct {
	server *model.Server

	children map[int]*ServerTcpTunnel

	listener *net.TCPListener

	running bool
}

func newServerTCP(server *model.Server) *ServerTCP {
	svr := &ServerTCP{
		server:   server,
		children: make(map[int]*ServerTcpTunnel),
	}
	return svr
}

func (server *ServerTCP) Open(ctx context.Context) error {
	if server.running {
		return errors.New("server is opened")
	}
	ServerOpenAction(server.server.Id)

	addr, err := net.ResolveTCPAddr("tcp", resolvePort(server.server.Addr))
	if err != nil {
		return err
	}
	server.listener, err = net.ListenTCP("tcp", addr)
	if err != nil {
		return err
	}

	server.running = true
	go func() {
		for {
			c, err := server.listener.AcceptTCP()
			if err != nil {
				//todo print err
				break
			}

			buf := make([]byte, 128)
			n := 0
			n, err = c.Read(buf)
			if err != nil {
				_ = c.Close()
				continue
			}
			data := buf[:n]
			deviceKey, checkIsOk := server.server.Register.Check(data)
			if !checkIsOk {
				_ = c.Close()
				continue
			}
			_, tunnelList, err := service.NetworkTunnel().GetTunnelList(ctx, &interlModel.GetNetworkTunnelListInput{
				ServiceId: server.server.Id,
				DeviceKey: deviceKey,
			})
			if err != nil {
				_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionError, server.server.Id), []byte(err.Error()))
				continue
			}
			tunnelId := 0
			if len(tunnelList) == 0 {
				heartbeatData, _ := json.Marshal(server.server.Heartbeat)
				var addTunnelError error
				if tunnelId, addTunnelError = service.NetworkTunnel().AddTunnel(ctx, interlModel.NetworkTunnelAddInput{
					ServerId:  server.server.Id,
					Name:      deviceKey,
					Types:     "server-tcp",
					Addr:      c.LocalAddr().String(),
					Remote:    c.RemoteAddr().String(),
					Heartbeat: string(heartbeatData),
					Protoccol: "",
					Status:    0,
					Remark:    "",
				}); addTunnelError != nil {
					_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionError, server.server.Id), []byte(addTunnelError.Error()))
					continue
				}
			} else {
				tunnelId = tunnelList[0].Id
				if editTunnelError := service.NetworkTunnel().EditTunnel(ctx, interlModel.NetworkTunnelEditInput{
					Id: tunnelList[0].Id,
					NetworkTunnelAddInput: interlModel.NetworkTunnelAddInput{
						ServerId:  tunnelList[0].ServerId,
						Name:      tunnelList[0].Name,
						Types:     tunnelList[0].Types,
						Addr:      c.LocalAddr().String(),
						Remote:    c.RemoteAddr().String(),
						Retry:     tunnelList[0].Retry,
						Heartbeat: tunnelList[0].Heartbeat,
						Serial:    tunnelList[0].Serial,
						Protoccol: tunnelList[0].Protoccol,
						Status:    tunnelList[0].Status,
						Remark:    tunnelList[0].Remark,
					},
				}); editTunnelError != nil {
					_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionError, server.server.Id), []byte(editTunnelError.Error()))
					continue
				}
			}
			tnl := newServerTcpTunnel(server.server.Id, tunnelId, deviceKey, c)
			go tnl.receive(ctx)
			server.children[tunnelId] = tnl
			ServerTunnelAction(ctx, server.server.Id, deviceKey)
		}
		server.running = false
	}()
	return nil
}

func (server *ServerTCP) Close() (err error) {
	ServerCloseAction(server.server.Id)
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

func (server *ServerTCP) GetTunnel(id int) tunnelinstance.TunnelInstance {
	return server.children[id]
}

func (server *ServerTCP) RemoveTunnel(id int) {
	delete(server.children, id)
}

func (server *ServerTCP) Running() bool {
	return server.running
}
