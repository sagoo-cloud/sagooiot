package core

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

func ServerOpenAction(serverId int) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionOpen, serverId), nil)
}

func ServerCloseAction(serverId int) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionClose, serverId), nil)
}

func ServerTunnelAction(ctx context.Context, serverId int, deviceKey string) {
	_ = mqtt.Publish(consts.GetWrapperTopic(consts.DataBusServer, consts.ActionTunnel, serverId), nil)
	deviceDetail, _ := service.DevDevice().Get(ctx, deviceKey)
	if deviceDetail == nil {
		g.Log().Errorf(ctx, "deviceKey:%s not found,ignore", deviceKey)
		return
	}
	_, err := LoadDevice(ctx, deviceDetail.Id)
	if err != nil {
		g.Log().Error(ctx, err)
	}
	if deviceDetail != nil && deviceDetail.Status != consts.DeviceStatueOnline {
		onlineMessage, _ := json.Marshal(model.DeviceOnlineMessage{
			DeviceKey: deviceKey,
			Timestamp: time.Now().Unix(),
		})
		if dataBusOnlineErr := mqtt.Publish(consts.GetDataBusWrapperTopic(deviceDetail.Product.Key, deviceKey, consts.DataBusOnline), onlineMessage); dataBusOnlineErr != nil {
			g.Log().Errorf(ctx, "publish  data error: %w, topic:%s, message:%s, message ignored", dataBusOnlineErr,
				consts.GetDataBusWrapperTopic(deviceDetail.Product.Key, deviceKey, consts.DataBusOnline),
				string(onlineMessage))
			return
		}
	}
}
