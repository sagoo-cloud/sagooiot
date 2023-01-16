package model

import (
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
)

const (
	AlarmTriggerTypeOnline   = iota + 1 // 触发类型：设备上线
	AlarmTriggerTypeOffline             // 触发类型：设备离线
	AlarmTriggerTypeProperty            // 触发类型：属性上报
	AlarmTriggerTypeEvent               // 触发类型：事件上报
)

var AlarmTriggerType = map[int]string{
	AlarmTriggerTypeOnline:   "设备上线",
	AlarmTriggerTypeOffline:  "设备离线",
	AlarmTriggerTypeProperty: "属性上报",
	AlarmTriggerTypeEvent:    "事件上报",
}

const (
	OperatorEq   = "eq"   // 操作符：等于
	OperatorNe   = "ne"   // 操作符：不等于
	OperatorGt   = "gt"   // 操作符：大于
	OperatorGte  = "gte"  // 操作符：大于等于
	OperatorLt   = "lt"   // 操作符：小于
	OperatorLte  = "lte"  // 操作符：小于等于
	OperatorBet  = "bet"  // 操作符：在...之间
	OperatorNbet = "nbet" // 操作符：不在...之间
)

const (
	AlarmRuleStatusOff int = iota // 告警规则状态：未启用
	AlarmRuleStatusOn             // 告警规则状态：已启用
)

type AlarmFilters struct {
	Key      string   `json:"key"      dc:"条件key"`
	Operator string   `json:"operator" dc:"操作符:eq,ne,gt,lt,gte,lte,bet,nbet"`
	Value    []string `json:"value"    dc:"条件值"`
	AndOr    int      `json:"andOr"    dc:"多个条件参数的关系：0=无，1=并且，2=或"`
}
type AlarmCondition struct {
	Filters []AlarmFilters `json:"filters" dc:"条件参数"`
	AndOr   int            `json:"andOr"   dc:"多个条件组的关系：0=无，1=并且，2=或"`
}
type AlarmTriggerCondition struct {
	TriggerCondition []AlarmCondition `json:"triggerCondition" dc:"触发条件" v:"required#请添加触发条件"`
}

type AlarmAction struct {
	SendGateway    string   `json:"sendGateway" dc:"通知发送通道：sms、work_weixin、dingding"`
	NoticeConfig   string   `json:"noticeConfig" dc:"通知配置"`
	NoticeTemplate string   `json:"noticeTemplate" dc:"通知模板"`
	Addressee      []string `json:"addressee" dc:"收信人"`
}
type AlarmPerformAction struct {
	Action []AlarmAction `json:"action" dc:"执行动作" v:"required#请添加执行动作"`
}

type AlarmRuleAddInput struct {
	Name        string `json:"name" dc:"告警规则名称" v:"required#请输入告警规则名称"`
	Level       uint   `json:"level" dc:"告警级别" v:"required#请选择告警级别"`
	ProductKey  string `json:"productKey" dc:"产品标识" v:"required#请选择产品"`
	DeviceKey   string `json:"deviceKey" dc:"设备标识"`
	TriggerType int    `json:"triggerType" dc:"触发类型：1=上线，2=离线，3=属性上报" v:"required#请选择触发类型"`
	AlarmTriggerCondition
	AlarmPerformAction
}

type AlarmRuleEditInput struct {
	Id uint64 `json:"id" dc:"告警规则ID" v:"required#告警规则ID不能为空"`
	AlarmRuleAddInput
}

type AlarmRuleOutput struct {
	*entity.AlarmRule

	TriggerTypeName string `json:"triggerTypeName" dc:"触发类型"`

	Condition AlarmTriggerCondition `json:"condition" dc:"触发条件"`

	PerformAction AlarmPerformAction `json:"performAction" dc:"执行动作"`

	AlarmLevel AlarmLevel `json:"alarmLevel" orm:"with:level" dc:"告警级别"`
}
type AlarmLevel struct {
	*entity.AlarmLevel
}

type OperatorOutput struct {
	Title string `json:"title" dc:"操作符名称"`
	Type  string `json:"type" dc:"操作符值"`
}

type TriggerTypeOutput struct {
	Title string `json:"title" dc:"触发类型"`
	Type  int    `json:"type" dc:"类型值"`
}

type TriggerParamOutput struct {
	Title    string `json:"title" dc:"条件参数"`
	ParamKey string `json:"paramKey" dc:"参数key"`
}

type AlarmRuleListInput struct {
	PaginationInput
}
type AlarmRuleListOutput struct {
	List []AlarmRuleOutput `json:"list" dc:"告警规则列表"`
	PaginationOutput
}
