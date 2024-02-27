package tunnel

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"net"
	"sagooiot/internal/consts"
	interlModel "sagooiot/internal/model"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/service"
	"sagooiot/network/core/tunnel/action"
	"sagooiot/network/model"
	"strconv"
	"time"
)

// TunnelClient 网络链接
type TunnelClient struct {
	TunnelBase
	tunnelInfo *model.Tunnel
	net        string
}

func newTunnelClient(tunnel *model.Tunnel, net string) *TunnelClient {
	return &TunnelClient{
		TunnelBase: TunnelBase{TunnelId: strconv.Itoa(int(tunnel.Id))},
		tunnelInfo: tunnel,
		net:        net,
	}
}

// Open 打开
func (client *TunnelClient) Open(ctx context.Context) error {
	if client.running {
		return errors.New("client is opened")
	}
	action.TunnelOpenAction(0, client.TunnelId)

	//发起连接
	conn, err := net.Dial(client.net, client.tunnelInfo.Addr)
	if err != nil {
		client.Retry(ctx)
		return err
	}
	client.retry = 0
	client.Link = conn

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
		_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionError, strconv.Itoa(int(client.tunnelInfo.Id))), []byte(editTunnelError.Error()))
		return editTunnelError
	}
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOnline, strconv.Itoa(int(client.tunnelInfo.Id))), nil)
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
		_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionError, strconv.Itoa(int(client.tunnelInfo.Id))), []byte(tunnelInfoErr.Error()))
	} else {
		onlineErr := action.TunnelOnlineAction(ctx, 0, strconv.Itoa(int(client.tunnelInfo.Id)), tunnelInfo.DeviceKey)
		if onlineErr != nil {
			g.Log().Errorf(ctx, "tunnel online error: %v", onlineErr)
			return
		}
		buf := make([]byte, 1024)
		for {
			n, err := client.Link.Read(buf)
			if err != nil {
				client.OnClose()
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
			go client.TunnelBase.ReadData(ctx, tunnelInfo.DeviceKey, data)
		}
	}
	client.running = false
	client.online = false
	action.TunnelOfflineAction(ctx, 0, client.TunnelId, tunnelInfo.DeviceKey)
	client.Retry(ctx)
}

// Close 关闭
func (client *TunnelClient) Close() error {
	client.running = false
	action.TunnelCloseAction(0, client.TunnelId)

	if client.Link != nil {
		link := client.Link
		client.Link = nil
		return link.Close()
	}
	return errors.New("tunnel is closed")
}
