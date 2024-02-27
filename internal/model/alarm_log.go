package model

import (
	"sagooiot/internal/model/entity"
)

const (
	AlarmLogStatusUnhandle int = iota // 告警日志状态：未处理
	AlarmLogStatusHandle              // 告警日志状态：已处理
	AlarmLogStatusIgnore              // 告警日志状态：忽略
)

// 告警日志
type AlarmLogOutput struct {
	*entity.AlarmLog

	AlarmLevel *AlarmLevel         `json:"alarmLevel" orm:"with:level"`
	Product    *DevProductWithName `json:"product" orm:"with:key=product_key"`
	Device     *DevDevice          `json:"devcie" orm:"with:key=device_key"`
}

// 告警日志写入
type AlarmLogAddInput struct {
	Type       uint   `json:"type" d:"1" dc:"告警类型：1=规则告警，2=设备自主告警"`
	RuleId     uint64 `json:"ruleId" dc:"规则id，type=2时为0"`
	RuleName   string `json:"ruleName" dc:"规则名称"`
	Level      uint   `json:"level" dc:"告警级别"`
	Data       string `json:"data" dc:"触发告警的数据"`
	Expression string `json:"expression" dc:"触发告警的表达式"`
	ProductKey string `json:"productKey" dc:"产品标识"`
	DeviceKey  string `json:"deviceKey" dc:"设备标识"`
}

// 告警处理
type AlarmLogHandleInput struct {
	Id      uint64 `json:"id" dc:"告警日志ID" v:"required#告警日志ID不能为空"`
	Status  int    `json:"status" d:"1" dc:"处理状态" v:"required|in:1,2#请选择处理状态|未知的处理状态，请正确选择"`
	Content string `json:"content" dc:"处理意见"`
}

// 日志级别统计
type AlarmLogLevelTotal struct {
	Level uint    `json:"level" dc:"告警级别"`
	Name  string  `json:"name" dc:"告警名称"`
	Num   int     `json:"num" dc:"该级别日志数量"`
	Ratio float64 `json:"ratio" dc:"该级别日志数量占比(百分比)"`
}

// 日志列表
type AlarmLogListInput struct {
	AlarmInput
	PaginationInput
}
type AlarmInput struct {
	Status string `p:"status"` //告警状态
}
type AlarmLogListOutput struct {
	List []AlarmLogOutput `json:"list" dc:"告警日志"`
	PaginationOutput
}
