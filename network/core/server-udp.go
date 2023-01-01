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

// ServerUDP UDP服务器
type ServerUDP struct {
	server *model.Server

	children map[int]*ServerUdpTunnel
	tunnels  map[string]*ServerUdpTunnel

	listener *net.UDPConn
	running  bool
}

func newServerUDP(server *model.Server) *ServerUDP {
	svr := &ServerUDP{
		server:   server,
		children: make(map[int]*ServerUdpTunnel),
		tunnels:  make(map[string]*ServerUdpTunnel),
	}
	return svr
}

// Open 打开
func (server *ServerUDP) Open(ctx context.Context) error {
	if server.running {
		return errors.New("server is opened")
	}
	ServerOpenAction(server.server.Id)

	addr, err := net.ResolveUDPAddr("udp", resolvePort(server.server.Addr))
	if err != nil {
		return err
	}
	c, err := net.ListenUDP("udp", addr)
	if err != nil {
		//TODO 需要正确处理接收错误
		return err
	}
	server.listener = c //共用连接

	server.running = true
	go func() {
		for {
			buf := make([]byte, 1024)
			n, addr, err := c.ReadFromUDP(buf)
			if err != nil {
				_ = c.Close()
				//continue
				break
			}
			data := buf[:n]

			//如果已经保存了链接 TODO 要有超时处理
			tnl, ok := server.tunnels[addr.String()]
			if ok {
				tnl.onData(ctx, data)
				continue
			}

			deviceKey, checkIsOk := server.server.Register.Check(data)
			if !checkIsOk {
				_ = c.Close()
				continue
			}
			tunnelId := 0
			_, tunnelList, err := service.NetworkTunnel().GetTunnelList(ctx, &interlModel.GetNetworkTunnelListInput{
				ServiceId: server.server.Id,
				DeviceKey: deviceKey,
			})
			if err != nil {
				_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionError, server.server.Id), []byte(err.Error()))
				continue
			}
			if len(tunnelList) == 0 {
				heartbeatData, _ := json.Marshal(server.server.Heartbeat)
				var addTunnelError error
				if tunnelId, addTunnelError = service.NetworkTunnel().AddTunnel(ctx, interlModel.NetworkTunnelAddInput{
					ServerId:  server.server.Id,
					Name:      deviceKey,
					Types:     "server-udp",
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
			_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOnline, server.server.Id), nil)

			tnl = newServerUdpTunnel(deviceKey, tunnelId, c, addr)
			server.children[tunnelId] = tnl
			ServerTunnelAction(ctx, server.server.Id, deviceKey)
			TunnelOnlineAction(ctx, tunnelId, deviceKey)
		}

		server.running = false
	}()

	return nil
}

// Close 关闭
func (server *ServerUDP) Close() (err error) {
	ServerCloseAction(server.server.Id)
	//close tunnels
	if server.children != nil {
		for _, l := range server.children {
			_ = l.Close()
		}
	}
	return server.listener.Close()
}

// GetTunnel 获取链接
func (server *ServerUDP) GetTunnel(id int) tunnelinstance.TunnelInstance {
	return server.children[id]
}

func (server *ServerUDP) RemoveTunnel(id int) {
	delete(server.children, id)
}

func (server *ServerUDP) Running() bool {
	return server.running
}
