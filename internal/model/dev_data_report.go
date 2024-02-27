package model

// 上报属性数据
type ReportPropertyData map[string]ReportPropertyNode

// 属性值
type ReportPropertyNode struct {
	Value      any   // 属性值
	CreateTime int64 // 上报时间
}

// 上报事件数据
type ReportEventData struct {
	Key   string           // 事件标识
	Param ReportEventParam // 事件输出参数
}

// 事件输出参数
type ReportEventParam struct {
	Value      map[string]any // 事件输出参数
	CreateTime int64          // 上报时间
}

// 设备上下线状态
type ReportStatusData struct {
	Status     string // 状态：online、offline
	CreateTime int64  // 上下线时间
}
