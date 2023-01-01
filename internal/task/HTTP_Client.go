package task

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"
)

// GetAccessURL 执行访问URL
func GetAccessURL(ctx context.Context) {

	//获取参数
	t := jobTask.TimeTaskList.GetByName("getAccessURL")
	if t == nil {
		return
	}
	glog.Infof(ctx, "执行任务：访问URL", gconv.Int(t.Param[0]))

	res, err := g.Client().Get(ctx, gconv.String(t.Param[0]))
	if err != nil {
		glog.Error(ctx, err)
	}
	defer res.Close()
}
