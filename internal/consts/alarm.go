package consts

const (
	AlarmTriggerModeDevice = iota + 1 // 触发方式：设备触发
	AlarmTriggerModeCron              // 触发方式：定时触发
)

const (
	AlarmTriggerTypeOnline   = iota + 1 // 触发类型：设备上线
	AlarmTriggerTypeOffline             // 触发类型：设备离线
	AlarmTriggerTypeProperty            // 触发类型：属性上报
	AlarmTriggerTypeEvent               // 触发类型：事件上报
)

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
