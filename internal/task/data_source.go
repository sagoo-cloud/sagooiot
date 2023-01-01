package task

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据源获取数据
func DataSource(ctx context.Context) {
	//获取参数
	t := jobTask.TimeTaskList.GetByName("dataSource")
	if t == nil {
		return
	}

	id := service.SysJob().Value(ctx)

	err := service.DataSource().UpdateData(ctx, id)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
