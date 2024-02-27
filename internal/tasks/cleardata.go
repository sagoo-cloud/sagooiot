package tasks

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/service"
)

// ClearOperationLogByDays 清理超过指定天数的操作日志
func (t TaskJob) ClearOperationLogByDays(days string) {
	ctx := context.Background()
	glog.Debugf(ctx, "执行任务：清理超过%v天的操作日志", days)
	err := service.SysOperLog().ClearOperationLogByDays(ctx, gconv.Int(days))
	if err != nil {
		glog.Error(ctx, err)
	}
}

// ClearNoticeLogByDays 清理超过指定天数的通知服务日志
func (t TaskJob) ClearNoticeLogByDays(days string) {
	ctx := context.Background()
	glog.Debugf(ctx, "执行任务：清理超过%d天的通知服务发送日志", gconv.Int(days))
	err := service.NoticeLog().ClearLogByDays(ctx, gconv.Int(days))
	if err != nil {
		glog.Error(ctx, err)
	}
}

// ClearAlarmLogByDays 清理超过指定天数的告警日志
func (t TaskJob) ClearAlarmLogByDays(days string) {
	ctx := context.Background()
	glog.Debugf(ctx, "执行任务：清理超过%d天的告警日志", gconv.Int(days))
	err := service.AlarmLog().ClearLogByDays(ctx, gconv.Int(days))
	if err != nil {
		glog.Error(ctx, err)
	}
}

// ClearTDengineLogByDays 清理超过指定天数的TD日志
func (t TaskJob) ClearTDengineLogByDays(days string) {
	ctx := context.Background()
	glog.Debugf(ctx, "执行任务：清理超过%d天的TD日志", gconv.Int(days))
	err := service.TdEngine().ClearLogByDays(ctx, gconv.Int(days))
	if err != nil {
		glog.Error(ctx, err)
	}
}
