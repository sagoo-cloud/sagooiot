// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJob is the golang structure for table sys_job.
type SysJob struct {
	JobId          int64       `json:"jobId"          description:"任务ID"`
	JobName        string      `json:"jobName"        description:"任务名称"`
	JobParams      string      `json:"jobParams"      description:"参数"`
	JobGroup       string      `json:"jobGroup"       description:"任务组名"`
	InvokeTarget   string      `json:"invokeTarget"   description:"调用目标字符串"`
	CronExpression string      `json:"cronExpression" description:"cron执行表达式"`
	MisfirePolicy  int         `json:"misfirePolicy"  description:"计划执行策略（1多次执行 2执行一次）"`
	Concurrent     int         `json:"concurrent"     description:"是否并发执行（0允许 1禁止）"`
	Status         int         `json:"status"         description:"状态（0正常 1暂停）"`
	CreatedBy      uint64      `json:"createdBy"      description:"创建者"`
	UpdatedBy      uint64      `json:"updatedBy"      description:"更新者"`
	Remark         string      `json:"remark"         description:"备注信息"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
}
