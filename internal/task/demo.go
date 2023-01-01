/*
* @desc:测试定时任务
* @company:Sagoo Co.,Ltd
* @Author: microrain
* @Date:   2021/7/16 15:52
 */

package task

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"
)

func Test1(ctx context.Context) {
	g.Log().Debug(ctx, "无参测试", gtime.Datetime())
}

func Test2(ctx context.Context) {
	//获取参数
	t := jobTask.TimeTaskList.GetByName("test2")
	if t == nil {
		return
	}
	for _, v := range t.Param {
		g.Log().Debugf(ctx, "参数:%s;  ", v, gtime.Datetime())

	}
	fmt.Println()
}
