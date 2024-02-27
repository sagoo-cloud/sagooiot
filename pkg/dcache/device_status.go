package dcache

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/queues"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/iotModel"
	"time"
)

// UpdateStatus 更新设备状态
func UpdateStatus(ctx context.Context, device *model.DeviceOutput) {
	timeout := 0
	if device.OnlineTimeout > 0 {
		timeout = device.OnlineTimeout
	}
	if timeout == 0 {
		//设备超时时间
		defaultTimeout, err := service.ConfigData().GetConfigByKey(ctx, consts.DeviceDefaultTimeoutTime)
		if err != nil || defaultTimeout == nil {
			defaultTimeout = &entity.SysConfig{
				ConfigValue: "30",
			}
		}
		timeout = gconv.Int(defaultTimeout.ConfigValue)
	}
	deviceStatus := GetDeviceStatus(ctx, device.Key)

	if deviceStatus == consts.DeviceStatueOnline {
		//正常数据上报，更新设备在线的缓存时间
		_, err := cache.Instance().UpdateExpire(ctx, consts.DeviceStatusPrefix+device.Key, time.Duration(timeout+1)*time.Second)
		if err != nil {
			g.Log().Debug(ctx, device.Key, "更新设备在线的缓存时间失败")
		}

	} else {
		var deviceStatusLog = new(iotModel.DeviceStatusLog)
		deviceStatusLog.Status = 2
		deviceStatusLog.Timestamp = time.Now()
		deviceStatusLog.DeviceKey = device.Key
		//设备首次上线，设置设备在线的缓存时间
		if err := cache.Instance().Set(ctx, consts.DeviceStatusPrefix+device.Key, deviceStatusLog, time.Duration(timeout)*time.Second); err != nil {
			g.Log().Debug(ctx, device.Key, "设置设备在线的缓存时间失败")
		}

		//添加延时下线消息判断
		go func(deviceKey string, timeOutSecond time.Duration) {
			if err := online(ctx, device); err != nil {
				g.Log().Debug(ctx, device.Key, "设备上线处理失败")
			}
			for {
				ds := GetDeviceStatus(ctx, deviceKey)
				if ds == consts.DeviceStatueOffline { //离线
					break
				}
				time.Sleep(timeOutSecond * time.Second)
			}

			//设备下线
			err := offline(ctx, device)
			if err != nil {
				return
			}
		}(device.Key, time.Duration(timeout)*time.Second)
	}
}

// pushDeviceStatus 推送设备状态
func pushDeviceStatus(deviceKey string, status int) {
	var deviceStatusLog = new(iotModel.DeviceStatusLog)
	deviceStatusLog.Status = status
	deviceStatusLog.Timestamp = time.Now()
	deviceStatusLog.DeviceKey = deviceKey
	data, _ := json.Marshal(deviceStatusLog)
	err := queues.DeviceStatusInfoUpdateWorker.Push(context.Background(), consts.QueueDeviceStatusInfoUpdate, data, 10)
	if err != nil {
		g.Log().Debug(context.Background(), "Push DeviceStatusInfoUpdateWorker: %v", err)
	}
}

// online 设备上线
func online(ctx context.Context, device *model.DeviceOutput) (err error) {
	//插入设备上线日志
	InertDeviceLog(ctx, consts.MsgTypeOnline, device.Key, iotModel.DeviceOnlineMessage{
		Timestamp: time.Now().UnixMilli(),
		Desc:      "",
	})
	pushDeviceStatus(device.Key, 2)

	//告警处理
	go func() {
		// 上线告警提醒
		data := iotModel.ReportStatusData{
			Status:     "online",
			CreateTime: gtime.Now().Unix(),
		}
		if err == nil {
			err = service.AlarmRule().Check(ctx, device.Product.Key, device.Key, consts.AlarmTriggerTypeOnline, data)
			if err != nil {
				g.Log().Errorf(ctx, "告警检测失败: %s", err.Error())
			}
		}
	}()

	return
}

// offline 设备下线
func offline(ctx context.Context, device *model.DeviceOutput) (err error) {
	InertDeviceLog(ctx, consts.MsgTypeOffline, device.Key, iotModel.DeviceOfflineMessage{
		Timestamp: time.Now().UnixMilli(),
		Desc:      "",
	})
	pushDeviceStatus(device.Key, 1)

	// 离线告警提醒
	data := iotModel.ReportStatusData{
		Status:     "offline",
		CreateTime: gtime.Now().Unix(),
	}
	if err == nil {
		err = service.AlarmRule().Check(ctx, device.ProductKey, device.Key, consts.AlarmTriggerTypeOffline, data)
	}
	return
}
