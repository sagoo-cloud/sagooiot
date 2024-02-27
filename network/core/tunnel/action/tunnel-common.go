package action

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/service"
	"sagooiot/network/core/device"
	"sagooiot/network/core/logic/baseLogic"
	"sagooiot/network/model"
	"time"
)

func TunnelOpenAction(serverId int, tunnelId string) {
	topic := consts.DataBusTunnel
	if serverId != 0 {
		topic = consts.DataBusServerTunnel
	}
	_ = mqtt.Publish(consts.GetWrapperTopic(topic, consts.ActionOpen, tunnelId), nil)
}

func TunnelCloseAction(serverId int, tunnelId string) {
	topic := consts.DataBusTunnel
	if serverId != 0 {
		topic = consts.DataBusServerTunnel
	}
	_ = mqtt.Publish(consts.GetWrapperTopic(topic, consts.ActionClose, tunnelId), nil)
}

func TunnelOnlineAction(ctx context.Context, serverId int, tunnelId string, deviceKey string) (err error) {
	topic := consts.DataBusTunnel
	if serverId != 0 {
		topic = consts.DataBusServerTunnel
	}
	_ = mqtt.Publish(consts.GetWrapperTopic(topic, consts.ActionOnline, tunnelId), nil)
	deviceRes, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		err = errors.New("device not found")
		return
	}
	deviceDetail := device.GetDevice(uint64(deviceRes.Id))
	if deviceDetail != nil {
		err = deviceDetail.Start(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
			return
		} else {
			if deviceRes != nil && deviceRes.Status != consts.DeviceStatueOnline {
				if dataBusOfflineErr := baseLogic.Online(ctx, model.DeviceOnlineMessage{
					DeviceKey:  deviceRes.DevDevice.Key,
					ProductKey: deviceRes.Product.Key,
					Timestamp:  time.Now().Unix(),
				}); dataBusOfflineErr != nil {
					g.Log().Errorf(ctx, "online err:%v", dataBusOfflineErr)
				}
			}
		}
	}
	return
}

func TunnelOfflineAction(ctx context.Context, serverId int, tunnelId, deviceKey string) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusTunnel, consts.ActionOffline, tunnelId), nil)
	if deviceKey != "" {
		deviceDetail, _ := service.DevDevice().Get(ctx, deviceKey)
		if deviceDetail == nil {
			g.Log().Errorf(ctx, "deviceKey:%s not found,ignore", deviceKey)
			return
		}
		if deviceStopError := device.GetDevice(uint64(deviceDetail.Id)).Stop(); deviceStopError != nil {
			g.Log().Errorf(ctx, "Stop device  error:%v ,ignore", deviceStopError)
		}
		if deviceDetail != nil && deviceDetail.Status == consts.DeviceStatueOnline {
			if offlineErr := baseLogic.Offline(ctx, model.DeviceOfflineMessage{
				DeviceKey:  deviceKey,
				ProductKey: deviceDetail.Product.Key,
				Timestamp:  time.Now().Unix(),
			}); offlineErr != nil {
				g.Log().Errorf(ctx, "offline err:%v", offlineErr)
			}
		}

	}
}
