package baseLogic

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/queues"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/statistics"
)

// InertTdLog 插入设备日志
func InertTdLog(ctx context.Context, logType, deviceKey string, obj interface{}) {

	statistics.CountDeviceData() // 统计设备数据

	str, strIsOk := obj.(string)
	content := str
	if !strIsOk {
		objStr, _ := json.Marshal(obj)
		content = string(objStr)
	}

	var deviceLog = model.TdLogAddInput{
		Ts:      gtime.Now(),
		Device:  deviceKey,
		Type:    logType,
		Content: content,
	}
	// 向设备缓存数据库插入数据
	if err := dcache.DB().InsertData(context.Background(), deviceKey, deviceLog); err != nil {
		g.Log().Debug(ctx, "Failed to insert data: %v", err)
	}
	data, _ := json.Marshal(deviceLog)
	err := queues.DeviceDataSaveWorker.Push(ctx, consts.QueueDeviceDataSaveTopic, data, 10)
	if err != nil {
		g.Log().Debug(ctx, "Run TaskDeviceDataSaveWorker: %v", err)
	}

}
