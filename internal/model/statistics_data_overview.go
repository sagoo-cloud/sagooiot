package model

type BrokenLineChildRes struct {
	AccessDay string  `json:"accessDay"           description:"日期"`
	Values    float64 `json:"values"              description:"数值"`
}

type TemperingRatioRes struct {
	TemperatureRange string `json:"temperatureRange"      description:"温度区间"`
	Num              string `json:"num"                   description:"区间数据量"`
	Rate             string `json:"rate"                  description:"占比率"`
}

// 物联概览统计数据
type ThingOverviewOutput struct {
	Overview   DeviceTotalOutput    `json:"overview" dc:"物联概览统计数据"`
	Device     ThingDevice          `json:"device" dc:"设备月度统计"`
	AlarmLevel []AlarmLogLevelTotal `json:"alarmLevel" dc:"告警日志级别统计"`
}
type ThingDevice struct {
	MsgTotal   map[int]int `json:"msgTotal" dc:"设备消息量月度统计"`
	AlarmTotal map[int]int `json:"alarmTotal" dc:"设备告警量月度统计"`
}
