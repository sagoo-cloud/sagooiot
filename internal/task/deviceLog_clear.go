package task

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"

	"github.com/gogf/gf/v2/frame/g"
)

// 设备日志清理
func DeviceLogClear(ctx context.Context) {
	//获取参数
	t := jobTask.TimeTaskList.GetByName("deviceLogClear")
	if t == nil {
		return
	}

	if err := service.TdLogTable().Clear(ctx); err != nil {
		g.Log().Error(ctx, err)
	}
}
