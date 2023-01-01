package core

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	interlModel "github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"net"
	"time"
)

// TunnelClient 网络链接
type TunnelClient struct {
	tunnelBase
	tunnelInfo *model.Tunnel
	net        string
}

func newTunnelClient(tunnel *model.Tunnel, net string) *TunnelClient {
	return &TunnelClient{
		tunnelBase: tunnelBase{tunnelId: int(tunnel.Id)},
		tunnelInfo: tunnel,
		net:        net,
	}
}

// Open 打开
func (client *TunnelClient) Open(ctx context.Context) error {
	if client.running {
		return errors.New("client is opened")
	}
	TunnelOpenAction(client.tunnelId)

	//发起连接
	conn, err := net.Dial(client.net, client.tunnelInfo.Addr)
	if err != nil {
		client.Retry(ctx)
		return err
	}
	client.retry = 0
	client.link = conn

	//开始接收数据
	go client.receive(ctx)
	if editTunnelError := service.NetworkTunnel().EditTunnel(ctx, interlModel.NetworkTunnelEditInput{
		Id: int(client.tunnelInfo.Id),
		NetworkTunnelAddInput: interlModel.NetworkTunnelAddInput{
			ServerId: client.tunnelInfo.ServerId,
			Name:     client.tunnelInfo.Name,
			Addr:     conn.LocalAddr().String(),
			Remote:   conn.RemoteAddr().String(),
		},
	}); editTunnelError != nil {
		_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionError, int(client.tunnelInfo.Id)), []byte(editTunnelError.Error()))
		return editTunnelError
	}
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOnline, int(client.tunnelInfo.Id)), nil)
	return nil
}

func (client *TunnelClient) Retry(ctx context.Context) {
	//重连
	retry := &client.tunnelInfo.Retry
	if retry.Enable && (retry.Maximum == 0 || client.retry < retry.Maximum) {
		client.retry++
		client.retryTimer = time.AfterFunc(time.Second*time.Duration(retry.Timeout), func() {
			client.retryTimer = nil
			err := client.Open(ctx)
			if err != nil {
				g.Log().Error(ctx, err)
			}
		})
	}
}

func (client *TunnelClient) receive(ctx context.Context) {
	client.running = true
	client.online = true

	tunnelInfo, tunnelInfoErr := service.NetworkTunnel().GetTunnelById(ctx, int(client.tunnelInfo.Id))
	if tunnelInfoErr != nil {
		_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionError, int(client.tunnelInfo.Id)), []byte(tunnelInfoErr.Error()))
	} else {
		TunnelOnlineAction(ctx, int(client.tunnelInfo.Id), tunnelInfo.DeviceKey)
		buf := make([]byte, 1024)
		for {
			n, err := client.link.Read(buf)
			if err != nil {
				client.onClose()
				break
			}
			if n == 0 {
				continue
			}

			data := buf[:n]
			//过滤心跳包
			if client.tunnelInfo.Heartbeat.Enable && client.tunnelInfo.Heartbeat.Check(data) {
				continue
			}

			//透传转发
			if client.pipe != nil {
				_, err = client.pipe.Write(data)
				if err != nil {
					client.pipe = nil
				} else {
					continue
				}
			}
			go client.tunnelBase.ReadData(ctx, tunnelInfo.DeviceKey, data)
		}
	}
	client.running = false
	client.online = false
	TunnelOfflineAction(ctx, 0, client.tunnelId)
	client.Retry(ctx)
}

// Close 关闭
func (client *TunnelClient) Close() error {
	client.running = false
	TunnelCloseAction(client.tunnelId)

	if client.link != nil {
		link := client.link
		client.link = nil
		return link.Close()
	}
	return errors.New("tunnel is closed")
}
