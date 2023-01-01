package core

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"time"
)

func TunnelOpenAction(tunnelId int) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOpen, tunnelId), nil)
}

func TunnelCloseAction(tunnelId int) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionClose, tunnelId), nil)
}

func TunnelOnlineAction(ctx context.Context, tunnelId int, deviceKey string) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOnline, tunnelId), nil)
	deviceRes, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	deviceDetail := GetDevice(uint64(deviceRes.Id))
	if deviceDetail != nil {
		err = deviceDetail.Start(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		} else {
			if deviceRes != nil && deviceRes.Status != consts.DeviceStatueOnline {
				onlineMessage, _ := json.Marshal(model.DeviceOnlineMessage{
					DeviceKey: deviceRes.DevDevice.Key,
					Timestamp: time.Now().Unix(),
				})
				if dataBusOfflineErr := mqtt.Publish(consts.GetDataBusWrapperTopic(deviceRes.Product.Key, deviceRes.DevDevice.Key, consts.DataBusOffline), onlineMessage); dataBusOfflineErr != nil {
					g.Log().Errorf(ctx, "publish data error: %w, topic:%s, message:%s, message ignored", dataBusOfflineErr,
						consts.GetDataBusWrapperTopic(deviceRes.Product.Key, deviceRes.DevDevice.Key, consts.DataBusOffline),
						string(onlineMessage))
					return
				}
			}
		}
	}
}

func TunnelOfflineAction(ctx context.Context, serverId, tunnelId int) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOffline, tunnelId), nil)
	tunnelInfo, err := service.NetworkTunnel().GetTunnelById(ctx, tunnelId)

	//TODO log error
	if err == nil && tunnelInfo != nil {
		deviceDetail, _ := service.DevDevice().Get(ctx, tunnelInfo.DeviceKey)
		if deviceDetail == nil {
			g.Log().Errorf(ctx, "deviceKey:%s not found,ignore", tunnelInfo.DeviceKey)
			return
		}
		if deviceStopError := GetDevice(uint64(deviceDetail.Id)).Stop(); deviceStopError != nil {
			g.Log().Errorf(ctx, "Stop device  error:%w ,ignore", deviceStopError)
		}
		if deviceDetail != nil && deviceDetail.Status == consts.DeviceStatueOnline {
			onlineMessage, _ := json.Marshal(model.DeviceOnlineMessage{
				DeviceKey: tunnelInfo.DeviceKey,
				Timestamp: time.Now().Unix(),
			})
			if dataBusOfflineErr := mqtt.Publish(consts.GetDataBusWrapperTopic(deviceDetail.Product.Key, tunnelInfo.DeviceKey, consts.DataBusOffline), onlineMessage); dataBusOfflineErr != nil {
				g.Log().Errorf(ctx, "publish data error: %w, topic:%s, message:%s, message ignored", dataBusOfflineErr,
					consts.GetDataBusWrapperTopic(deviceDetail.Product.Key, tunnelInfo.DeviceKey, consts.DataBusOffline),
					string(onlineMessage))
				return
			}
		}

	} else if err != nil {
		g.Log().Errorf(ctx, "getTunnel ifno error:%w ,ignore", err)
	}

	if serverId != 0 {
		server := GetServer(serverId)
		if server != nil {
			server.Instance.RemoveTunnel(tunnelId)
		}
	}
}
