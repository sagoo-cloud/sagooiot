package model

type DataOverviewRes struct {
	CityNO         string  `json:"cityNO"           description:"城市编号"`
	Area           float64 `json:"area"             description:"供热面积"`
	AllArea        float64 `json:"allArea"          description:"总面积"`
	Calorie        float64 `json:"calorie"          description:"总耗热"`
	SingleCalorie  float64 `json:"singleCalorie"    description:"热总单耗"`
	Electric       float64 `json:"electric"         description:"总耗电"`
	SingleElectric float64 `json:"singleElectric"   description:"电总单耗"`
	Water          float64 `json:"water"            description:"总耗水"`
	SingleWater    float64 `json:"singleWater"      description:"水总单耗"`
}

type BrokenLineRes struct {
	Calorie  []*BrokenLineChildRes `json:"calorie"   description:"总热耗"`
	Electric []*BrokenLineChildRes `json:"electric"  description:"总电耗"`
	Water    []*BrokenLineChildRes `json:"water"     description:"总水耗"`
}

type BrokenLineChildRes struct {
	AccessDay string  `json:"accessDay"           description:"日期"`
	Values    float64 `json:"values"              description:"数值"`
}

type BarChartRes struct {
	HuanLuNo        string  `json:"huanLuNo"              description:"换热站编号"`
	HuanLuName      string  `json:"huanLuName"            description:"换热站名称"`
	InPressure1     float64 `json:"inPressure1"           description:"一网供水压力"`
	InPressure2     float64 `json:"inPressure2"           description:"二网供水压力"`
	InTemperature1  float64 `json:"inTemperature1"        description:"一网供水温度"`
	InTemperature2  float64 `json:"inTemperature2"        description:"二网供水温度"`
	OutPressure1    float64 `json:"outPressure1"          description:"一网回水压力"`
	OutPressure2    float64 `json:"outPressure2"          description:"二网回水压力"`
	OutTemperature1 float64 `json:"outTemperature1"       description:"一网回水温度"`
	OutTemperature2 float64 `json:"outTemperature2"       description:"二网回水温度"`
}

type TemperingRatioRes struct {
	TemperatureRange string `json:"temperatureRange"      description:"温度区间"`
	Num              string `json:"num"                   description:"区间数据量"`
	Rate             string `json:"rate"                  description:"占比率"`
}

// 物联概览统计数据
type ThingOverviewOutput struct {
	Overview     DeviceTotalOutput    `json:"overview" dc:"物联概览统计数据"`
	Device       ThingDevice          `json:"device" dc:"设备月度统计"`
	DeviceForDay ThingDeviceForDay    `json:"deviceForDay" dc:"设备近一个月统计"`
	AlarmLevel   []AlarmLogLevelTotal `json:"alarmLevel" dc:"告警日志级别统计"`
}
type ThingDevice struct {
	MsgTotal   map[int]int `json:"msgTotal" dc:"设备消息量月度统计"`
	AlarmTotal map[int]int `json:"alarmTotal" dc:"设备告警量月度统计"`
}
type ThingDeviceForDay struct {
	MsgTotal   map[string]int `json:"msgTotal" dc:"设备消息量近一个月统计"`
	AlarmTotal map[string]int `json:"alarmTotal" dc:"设备告警量近一个月统计"`
}
