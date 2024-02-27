package device

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/internal/service"
	"sagooiot/network/core/logic/baseLogic"
	"sagooiot/network/model"
	"time"
)

func StartAction(ctx context.Context, deviceKey string) {
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
				if deviceOnlineErr := baseLogic.Online(ctx, model.DeviceOnlineMessage{
					DeviceKey:  deviceRes.DevDevice.Key,
					ProductKey: deviceRes.Product.Key,
					Timestamp:  time.Now().Unix(),
				}); deviceOnlineErr != nil {
					g.Log().Errorf(ctx, "device online error:%v", deviceOnlineErr)
				}
				return
			}
		}
	}
}

func OffAction(ctx context.Context, tunnelId int) {
	tunnelInfo, err := service.NetworkTunnel().GetTunnelById(ctx, tunnelId)

	//TODO log error
	if err == nil && tunnelInfo != nil {
		deviceDetail, _ := service.DevDevice().Get(ctx, tunnelInfo.DeviceKey)
		if deviceDetail == nil {
			g.Log().Errorf(ctx, "deviceKey:%s not found,ignore", tunnelInfo.DeviceKey)
			return
		}
		if deviceStopError := GetDevice(uint64(deviceDetail.Id)).Stop(); deviceStopError != nil {
			g.Log().Errorf(ctx, "Stop device  error:%v ,ignore", deviceStopError)
		}
		if deviceDetail != nil && deviceDetail.Status == consts.DeviceStatueOnline {
			if deviceOnlineErr := baseLogic.Offline(ctx, model.DeviceOfflineMessage{
				DeviceKey:  tunnelInfo.DeviceKey,
				ProductKey: deviceDetail.Product.Key,
				Timestamp:  time.Now().Unix(),
			}); deviceOnlineErr != nil {
				g.Log().Errorf(ctx, "device online error:%v", deviceOnlineErr)
			}
		}

	} else if err != nil {
		g.Log().Errorf(ctx, "getTunnel ifno error:%v ,ignore", err)
	}

}
