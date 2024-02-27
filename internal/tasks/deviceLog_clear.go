package tasks

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"sagooiot/internal/service"
)

// DeviceLogClear 设备日志清理
func (t TaskJob) DeviceLogClear() {
	ctx := context.Background()
	if err := service.TdLogTable().Clear(ctx); err != nil {
		g.Log().Error(ctx, err)
	}
	glog.Debug(ctx, "执行任务：清理设备日志")
}
