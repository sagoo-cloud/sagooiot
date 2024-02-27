// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysJob is the golang structure of table sys_job for DAO operations like Where/Data.
type SysJob struct {
	g.Meta         `orm:"table:sys_job, do:true"`
	JobId          interface{} // 任务ID
	JobName        interface{} // 任务名称
	JobParams      interface{} // 参数
	JobGroup       interface{} // 任务组名
	InvokeTarget   interface{} // 调用目标字符串
	CronExpression interface{} // cron执行表达式
	MisfirePolicy  interface{} // 计划执行策略（1多次执行 2执行一次）
	Concurrent     interface{} // 是否并发执行（0允许 1禁止）
	Status         interface{} // 状态（0正常 1暂停）
	CreatedBy      interface{} // 创建者
	UpdatedBy      interface{} // 更新者
	Remark         interface{} // 备注信息
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
