package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type GetJobListInput struct {
	JobName  string `json:"jobName" description:"任务名称"`
	JobGroup string `json:"jobGroup" description:"任务组名"`
	Status   string `json:"status" description:"状态（0正常 1暂停）"`
	*PaginationInput
}

type SysJobRes struct {
	JobId          int64       `orm:"job_id,primary"    json:"jobId"`          // 任务ID
	JobName        string      `orm:"job_name,primary"  json:"jobName"`        // 任务名称
	JobParams      string      `orm:"job_params"        json:"jobParams"`      // 参数
	JobGroup       string      `orm:"job_group,primary" json:"jobGroup"`       // 任务组名
	InvokeTarget   string      `orm:"invoke_target"     json:"invokeTarget"`   // 调用目标字符串
	CronExpression string      `orm:"cron_expression"   json:"cronExpression"` // cron执行表达式
	MisfirePolicy  int         `orm:"misfire_policy"    json:"misfirePolicy"`  // 计划执行策略（1多次执行 2执行一次）
	Concurrent     int         `orm:"concurrent"        json:"concurrent"`     // 是否并发执行（0允许 1禁止）
	Status         int         `orm:"status"            json:"status"`         // 状态（0正常 1暂停）
	CreateBy       uint64      `orm:"create_by"         json:"createBy"`       // 创建者
	UpdateBy       uint64      `orm:"update_by"         json:"updateBy"`       // 更新者
	Remark         string      `orm:"remark"            json:"remark"`         // 备注信息
	CreatedAt      *gtime.Time `orm:"created_at"        json:"createdAt"`      // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"        json:"updatedAt"`      // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"        json:"deletedAt"`      // 删除时间
}
type SysJobFunListOut struct {
	FunName string `json:"fun_name"`
	Explain string `json:"explain"`
}

type SysJobOut struct {
	JobId          int64       `orm:"job_id,primary"    json:"jobId"`          // 任务ID
	JobName        string      `orm:"job_name,primary"  json:"jobName"`        // 任务名称
	JobParams      string      `orm:"job_params"        json:"jobParams"`      // 参数
	JobGroup       string      `orm:"job_group,primary" json:"jobGroup"`       // 任务组名
	InvokeTarget   string      `orm:"invoke_target"     json:"invokeTarget"`   // 调用目标字符串
	CronExpression string      `orm:"cron_expression"   json:"cronExpression"` // cron执行表达式
	MisfirePolicy  int         `orm:"misfire_policy"    json:"misfirePolicy"`  // 计划执行策略（1多次执行 2执行一次）
	Concurrent     int         `orm:"concurrent"        json:"concurrent"`     // 是否并发执行（0允许 1禁止）
	Status         int         `orm:"status"            json:"status"`         // 状态（0正常 1暂停）
	CreateBy       uint64      `orm:"create_by"         json:"createBy"`       // 创建者
	UpdateBy       uint64      `orm:"update_by"         json:"updateBy"`       // 更新者
	Remark         string      `orm:"remark"            json:"remark"`         // 备注信息
	CreatedAt      *gtime.Time `orm:"created_at"        json:"createdAt"`      // 创建时间
	UpdatedAt      *gtime.Time `orm:"updated_at"        json:"updatedAt"`      // 更新时间
	DeletedAt      *gtime.Time `orm:"deleted_at"        json:"deletedAt"`      // 删除时间
}

// SysJobAddInput 添加JOB
type SysJobAddInput struct {
	JobName        string `json:"jobName"  description:"任务名称"`
	JobParams      string `json:"jobParams"  description:"任务参数"`
	JobGroup       string `json:"jobGroup" description:"分组"`
	InvokeTarget   string `json:"invokeTarget" description:"执行方法"`
	CronExpression string `json:"cronExpression" description:"任务执行表达式" `
	MisfirePolicy  int    `json:"misfirePolicy"`
	Concurrent     int    `json:"concurrent" `
	Status         int    `json:"status" description:"状态" `
	Remark         string `json:"remark" `
	CreateBy       uint64
}

// SysJobEditInput 修改JOB
type SysJobEditInput struct {
	JobId          int64  `json:"job_id" v:"min:1#任务id不能为空"`
	JobName        string `json:"jobName"  description:"任务名称" `
	JobParams      string `json:"jobParams"  description:"任务参数"`
	JobGroup       string `json:"jobGroup" description:"分组"`
	InvokeTarget   string `json:"invokeTarget" description:"执行方法" `
	CronExpression string `json:"cronExpression" description:"任务执行表达式" `
	MisfirePolicy  int    `json:"misfirePolicy"`
	Concurrent     int    `json:"concurrent" `
	Status         int    `json:"status" description:"状态"`
	Remark         string `json:"remark" `
	CreateBy       uint64
	UpdateBy       uint64
}
