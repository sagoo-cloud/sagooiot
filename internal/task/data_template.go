package task

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据模型聚合数据
func DataTemplate(ctx context.Context) {
	//获取参数
	t := jobTask.TimeTaskList.GetByName("dataTemplate")
	if t == nil {
		return
	}

	id := service.SysJob().Value(ctx)

	err := service.DataTemplate().UpdateData(ctx, id)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
