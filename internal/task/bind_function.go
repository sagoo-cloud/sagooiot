/*
* @desc:定时任务配置
* @company:Sagoo Co.,Ltd
* @Author: microrain
* @Date:   2021/7/16 15:45
 */

package task

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/jobTask"
	"strings"

	"github.com/gogf/gf/v2/frame/g"
)

func StartInit() {

	ctx := context.TODO()

	sysConfigInfo, err := service.ConfigData().GetConfigByKey(ctx, consts.IsAutoRunJob)
	if err != nil {
		return
	}
	if sysConfigInfo != nil {
		if strings.EqualFold(sysConfigInfo.ConfigValue, "1") {
			task1 := &jobTask.TimeTask{
				FuncName: "test1",
				Run:      Test1,
			}
			task2 := &jobTask.TimeTask{
				FuncName: "test2",
				Run:      Test2,
			}
			clearOperationLogByDaysTask := &jobTask.TimeTask{
				FuncName: "clearOperationLogByDays",
				Run:      ClearOperationLogByDays,
			}

			clearNoticeLogByDaysTask := &jobTask.TimeTask{
				FuncName: "clearNoticeLogByDays",
				Run:      ClearNoticeLogByDays,
			}

			clearAlarmLogByDaysTask := &jobTask.TimeTask{
				FuncName: "clearAlarmLogByDays",
				Run:      ClearAlarmLogByDays,
			}

			getAccessURLTask := &jobTask.TimeTask{
				FuncName: "getAccessURL",
				Run:      GetAccessURL,
			}

			dataSourceTask := &jobTask.TimeTask{
				FuncName: "dataSource",
				Run:      DataSource,
			}

			dataTemplateTask := &jobTask.TimeTask{
				FuncName: "dataTemplate",
				Run:      DataTemplate,
			}

			deviceLogClearTask := &jobTask.TimeTask{
				FuncName: "deviceLogClear",
				Run:      DeviceLogClear,
			}

			jobTask.TimeTaskList.AddTask(task1).AddTask(task2).
				AddTask(clearOperationLogByDaysTask).
				AddTask(clearNoticeLogByDaysTask).
				AddTask(clearAlarmLogByDaysTask).
				AddTask(getAccessURLTask).
				AddTask(dataSourceTask).
				AddTask(dataTemplateTask).
				AddTask(deviceLogClearTask)

			//自动执行已开启的任务
			var jobs []*model.SysJobOut
			jobs, err = service.SysJob().GetJobs(ctx)
			if err != nil {
				g.Log().Error(ctx, err)
			}

			//批量启动任务
			err = service.SysJob().JobStartMult(ctx, jobs)
			if err != nil {
				g.Log().Error(ctx, err.Error())
			}
		}
	}

}
