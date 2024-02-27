// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AlarmLog is the golang structure of table alarm_log for DAO operations like Where/Data.
type AlarmLog struct {
	g.Meta     `orm:"table:alarm_log, do:true"`
	Id         interface{} //
	DeptId     interface{} // 部门ID
	Type       interface{} // 告警类型：1=规则告警，2=设备自主告警
	RuleId     interface{} // 规则id
	RuleName   interface{} // 规则名称
	Level      interface{} // 告警级别
	Data       interface{} // 触发告警的数据
	Expression interface{} // 触发告警的表达式
	ProductKey interface{} // 产品标识
	DeviceKey  interface{} // 设备标识
	Status     interface{} // 告警状态：0=未处理，1=已处理
	CreatedAt  *gtime.Time // 告警时间
	UpdatedBy  interface{} // 告警处理人员
	UpdatedAt  *gtime.Time // 处理时间
	Content    interface{} // 处理意见
}
