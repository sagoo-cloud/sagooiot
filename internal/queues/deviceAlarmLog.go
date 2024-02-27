package queues

import (
	"context"
	"encoding/json"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/worker"
)

var ScheduledDeviceAlarmLog = new(worker.Scheduled)

func ScheduledDeviceAlarmLogRun() {
	ScheduledDeviceAlarmLog = worker.RegisterProcess(DeviceAlarmLog)
}

// DeviceAlarmLog 设备告警日志
var DeviceAlarmLog = &qDeviceAlarmLog{}

type qDeviceAlarmLog struct{}

// GetTopic 主题
func (q *qDeviceAlarmLog) GetTopic() string {
	return consts.QueueDeviceAlarmLogTopic
}

// Handle 处理消息
func (q *qDeviceAlarmLog) Handle(ctx context.Context, p worker.Payload) (err error) {
	if p.Payload == nil || q.GetTopic() != p.Group {
		return nil
	}
	var data model.AlarmLogAddInput
	if err = json.Unmarshal(p.Payload, &data); err != nil {
		return err
	}
	//真正写日志
	_, err = service.AlarmLog().Add(ctx, &data)

	return
}
