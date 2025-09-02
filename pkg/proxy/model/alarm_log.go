package model

import "github.com/gogf/gf/v2/os/gtime"

type AlarmLog struct {
	Id         int64       `json:"id"         description:""`
	DeptId     int         `json:"deptId"     description:"部门ID"`
	Type       uint        `json:"type"       description:"告警类型：1=规则告警，2=设备自主告警"`
	RuleId     uint64      `json:"ruleId"     description:"规则id"`
	RuleName   string      `json:"ruleName"   description:"规则名称"`
	Level      uint        `json:"level"      description:"告警级别"`
	Data       string      `json:"data"       description:"触发告警的数据"`
	Expression string      `json:"expression" description:"触发告警的表达式"`
	ProductKey string      `json:"productKey" description:"产品标识"`
	DeviceKey  string      `json:"deviceKey"  description:"设备标识"`
	Status     int         `json:"status"     description:"告警状态：0=未处理，1=已处理"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"告警时间"`
	UpdatedBy  uint        `json:"updatedBy"  description:"告警处理人员"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"处理时间"`
	Content    string      `json:"content"    description:"处理意见"`
}
