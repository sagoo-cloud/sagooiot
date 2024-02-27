package base

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

const (
	SagooServerTunnelPrefix = "sagoo-server-tunnel"
)

type ServerTunnel struct {
	// server的id
	ServerId int
	// server通道id
	TunnelId string
	// 设备key
	DeviceKey string
	// server通道类型,暂时只支持tcp, udp
	Type string
	// server通道状态，0:离线，1:在线
	Status int
	// server通道内地地址
	LocalAddr string
	// server通道远程地址
	RemoteAddr string
	// server通道备注
	Remark string `json:"remark"    description:"备注"`
}

// AddOrEditServerTunnel 添加或编辑server通道
func AddOrEditServerTunnel(ctx context.Context, t ServerTunnel) (tunnelId string, err error) {
	tunnelJson, _ := json.Marshal(t)
	tunnelId = fmt.Sprintf("%s-%s", SagooServerTunnelPrefix, t.DeviceKey)
	t.TunnelId = tunnelId
	_, err = g.Redis().Do(ctx, "SET", fmt.Sprintf("%s-%s", SagooServerTunnelPrefix, t.DeviceKey), string(tunnelJson))
	return tunnelId, err
}

// GetServerTunnel 获取server通道
func GetServerTunnel(ctx context.Context, deviceKey string) (t ServerTunnel, err error) {
	tunnelVar, err := g.Redis().Do(ctx, "GET", fmt.Sprintf("%s-%s", SagooServerTunnelPrefix, deviceKey))
	if err != nil {
		return t, err
	}
	err = json.Unmarshal([]byte(tunnelVar.String()), &t)
	return
}

// GetServerTunnelList 获取server通道列表
func GetServerTunnelList(ctx context.Context) (listT []ServerTunnel, err error) {
	tunnelListVar, err := g.Redis().Do(ctx, "KEYS", fmt.Sprintf("%s-*", SagooServerTunnelPrefix))
	if err != nil {
		return listT, err
	}
	for _, node := range tunnelListVar.Strings() {
		var t ServerTunnel
		if err = json.Unmarshal([]byte(node), &t); err != nil {
			return listT, err
		}
		listT = append(listT, t)
	}
	return listT, nil
}

func GetTunnelIdByDeviceKey(ctx context.Context, deviceKey string) (tunnelId string) {
	return fmt.Sprintf("%s-%s", SagooServerTunnelPrefix, deviceKey)
}

// DeleteServerTunnel 删除server通道
func DeleteServerTunnel(ctx context.Context, deviceKey string) (t ServerTunnel, err error) {
	t, err = GetServerTunnel(ctx, deviceKey)
	if err != nil {
		return t, err
	}
	_, err = g.Redis().Do(ctx, "DEL", fmt.Sprintf("%s-%s", SagooServerTunnelPrefix, deviceKey))
	return t, err
}

// TunnelOnline server通道上线
func (t ServerTunnel) TunnelOnline(ctx context.Context) error {
	t, err := GetServerTunnel(ctx, t.DeviceKey)
	if err != nil {
		return err
	}
	if t.Status == 1 {
		return nil
	} else {
		t.Status = 1
		_, err = AddOrEditServerTunnel(ctx, t)
		return err
	}
}

// TunnelOffline server通道下线
func (t ServerTunnel) TunnelOffline(ctx context.Context) error {
	_, err := GetServerTunnel(ctx, t.DeviceKey)
	if err != nil {
		return err
	}
	_, err = DeleteServerTunnel(ctx, t.DeviceKey)
	return err
}
