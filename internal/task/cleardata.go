package task

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"
)

// ClearOperationLogByDays 清理超过指定天数的操作日志
func ClearOperationLogByDays(ctx context.Context) {
	//获取参数
	t := jobTask.TimeTaskList.GetByName("clearOperationLogByDays")
	if t == nil {
		return
	}
	glog.Infof(ctx, "执行任务：清理超过%d天的操作日志", gconv.Int(t.Param[0]))
	err := service.SysOperLog().ClearOperationLogByDays(ctx, gconv.Int(t.Param[0]))
	if err != nil {
		glog.Error(ctx, err)
	}
}

// ClearNoticeLogByDays 清理超过指定天数的通知服务日志
func ClearNoticeLogByDays(ctx context.Context) {
	//获取参数
	t := jobTask.TimeTaskList.GetByName("clearNoticeLogByDays")
	if t == nil {
		return
	}
	glog.Infof(ctx, "执行任务：清理超过%d天的通知服务发送日志", gconv.Int(t.Param[0]))
	err := service.NoticeLog().ClearLogByDays(ctx, gconv.Int(t.Param[0]))
	if err != nil {
		glog.Error(ctx, err)
	}
}

// ClearAlarmLogByDays 清理超过指定天数的告警日志
func ClearAlarmLogByDays(ctx context.Context) {
	//获取参数
	t := jobTask.TimeTaskList.GetByName("clearAlarmLogByDays")
	if t == nil {
		return
	}
	glog.Infof(ctx, "执行任务：清理超过%d天的告警日志", gconv.Int(t.Param[0]))
	err := service.AlarmLog().ClearLogByDays(ctx, gconv.Int(t.Param[0]))
	if err != nil {
		glog.Error(ctx, err)
	}
}
