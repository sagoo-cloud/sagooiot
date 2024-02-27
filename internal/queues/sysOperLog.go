package queues

import (
	"context"
	"encoding/json"
	"sagooiot/internal/consts"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/worker"
)

var ScheduledSysOperLog = new(worker.Scheduled)

func ScheduledSysOperLogRun() {
	ScheduledSysOperLog = worker.RegisterProcess(SysOperLog)
}

// SysOperLog 系统日志
var SysOperLog = &qSysOperLog{}

type qSysOperLog struct{}

// GetTopic 主题
func (q *qSysOperLog) GetTopic() string {
	return consts.QueueRequestLogTopic
}

// Handle 处理消息
func (q *qSysOperLog) Handle(ctx context.Context, p worker.Payload) (err error) {
	if p.Payload == nil || q.GetTopic() != p.Group {
		return nil
	}
	var data entity.SysOperLog
	if err = json.Unmarshal(p.Payload, &data); err != nil {
		return err
	}
	//真正写日志
	return service.SysOperLog().RealWrite(ctx, data)
}
