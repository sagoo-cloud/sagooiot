package common

import (
	"context"
	"sagooiot/internal/consts"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/service"
	"sagooiot/network/core/device"
	"sagooiot/network/core/logic/baseLogic"
	"sagooiot/network/model"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func ServerOpenAction(serverId int) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServerTunnel, consts.ActionOpen, strconv.Itoa(serverId)), nil)
}

func ServerCloseAction(serverId int) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServerTunnel, consts.ActionClose, strconv.Itoa(serverId)), nil)
}

func ServerTunnelAction(ctx context.Context, serverId int, deviceKey string) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServerTunnel, consts.ActionTunnel, strconv.Itoa(serverId)), nil)
	deviceDetail, _ := service.DevDevice().Get(ctx, deviceKey)
	if deviceDetail == nil {
		g.Log().Debugf(ctx, "deviceKey:%s not found,ignore", deviceKey)
		return
	}
	_, err := device.LoadDevice(ctx, deviceDetail.Key)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	if deviceDetail != nil && deviceDetail.Status != consts.DeviceStatueOnline {
		if dataBusOfflineErr := baseLogic.Online(ctx, model.DeviceOnlineMessage{
			DeviceKey:  deviceKey,
			ProductKey: deviceDetail.Product.Key,
			Timestamp:  time.Now().Unix(),
		}); dataBusOfflineErr != nil {
			g.Log().Errorf(ctx, "online err:%v", dataBusOfflineErr)
		}
	}
}
