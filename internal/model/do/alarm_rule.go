// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AlarmRule is the golang structure of table alarm_rule for DAO operations like Where/Data.
type AlarmRule struct {
	g.Meta           `orm:"table:alarm_rule, do:true"`
	Id               interface{} //
	DeptId           interface{} // 部门ID
	Name             interface{} // 告警规则名称
	Level            interface{} // 告警级别，默认：4（一般）
	ProductKey       interface{} // 产品标识
	DeviceKey        interface{} // 设备标识
	TriggerMode      interface{} // 触发方式：1=设备触发，2=定时触发
	TriggerType      interface{} // 触发类型：1=上线，2=离线，3=属性上报, 4=事件上报
	EventKey         interface{} // 事件标识
	TriggerCondition interface{} // 触发条件
	Action           interface{} // 执行动作
	Status           interface{} // 状态：0=未启用，1=已启用
	CreatedBy        interface{} // 创建者
	UpdatedBy        interface{} // 更新者
	DeletedBy        interface{} // 删除者
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
	DeletedAt        *gtime.Time // 删除时间
}
