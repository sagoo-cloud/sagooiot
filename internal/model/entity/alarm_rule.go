// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AlarmRule is the golang structure for table alarm_rule.
type AlarmRule struct {
	Id               uint64      `json:"id"               description:""`
	DeptId           int         `json:"deptId"           description:"部门ID"`
	Name             string      `json:"name"             description:"告警规则名称"`
	Level            uint        `json:"level"            description:"告警级别，默认：4（一般）"`
	ProductKey       string      `json:"productKey"       description:"产品标识"`
	DeviceKey        string      `json:"deviceKey"        description:"设备标识"`
	TriggerMode      int         `json:"triggerMode"      description:"触发方式：1=设备触发，2=定时触发"`
	TriggerType      int         `json:"triggerType"      description:"触发类型：1=上线，2=离线，3=属性上报, 4=事件上报"`
	EventKey         string      `json:"eventKey"         description:"事件标识"`
	TriggerCondition string      `json:"triggerCondition" description:"触发条件"`
	Action           string      `json:"action"           description:"执行动作"`
	Status           int         `json:"status"           description:"状态：0=未启用，1=已启用"`
	CreatedBy        uint        `json:"createdBy"        description:"创建者"`
	UpdatedBy        uint        `json:"updatedBy"        description:"更新者"`
	DeletedBy        uint        `json:"deletedBy"        description:"删除者"`
	CreatedAt        *gtime.Time `json:"createdAt"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        description:"更新时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        description:"删除时间"`
}
