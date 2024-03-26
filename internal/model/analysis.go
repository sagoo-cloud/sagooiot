package model

import "sagooiot/pkg/iotModel"

// DeviceOnlineOfflineCount 设备在线离线状态统计
type DeviceOnlineOfflineCount struct {
	Total   int `json:"total"`   // 设备总数
	Online  int `json:"online"`  // 在线设备数
	Offline int `json:"offline"` // 离线设备数
	Disable int `json:"disable"` // 禁用设备数
}

// CountData 统计数据
type CountData struct {
	Title string
	Value int64
}

// ProductCountRes 产品数量统计
type ProductCountRes struct {
	Total   int `json:"total"`   // 产品总数
	Disable int `json:"disable"` // 禁用产品数
	Enable  int `json:"enable"`  // 启用产品数
	Added   int `json:"added"`   // 新增产品数
}

// DeviceDataReq 设备数据请求
type DeviceDataReq struct {
	DeviceKey string `json:"deviceKey" v:"required#设备key不能为空" dc:"设备key"`
	PaginationInput
}

// DeviceDataRes 设备数据响应
type DeviceDataRes struct {
	DeviceKey  string                      `json:"deviceKey"`
	DeviceData iotModel.ReportPropertyData `json:"deviceData"`
}

type DeviceIndicatorTrendReq struct {
	ProductKey string `json:"productKey" v:"required#产品key不能为空"`
	DeviceKey  string `json:"deviceKey" v:"required#设备key不能为空"`
	Properties string `json:"properties" v:"required#设备属性不能为空"`
	StartDate  string `json:"startDate" v:"required#开始时间不能为空"`
	EndDate    string `json:"endDate" v:"required#结束时间不能为空"`
}

type DeviceIndicatorTrendRes struct {
	DataValue float64 `json:"dataValue" dc:"属性值"`
	Date      string  `json:"date" dc:"属性值上报时间"`
}

// DeviceIndicatorPolymerizeReq 设备指标聚合
type DeviceIndicatorPolymerizeReq struct {
	DateType   string `json:"dateType" v:"required#日期类型不能为空" dc:"日期类型：1 yyyy-MM-dd HH:mm 5分钟，2 一小时 yyyy-MM-dd HH ，3 一天 yyyy-MM-dd；对应时间范围为 一周，一个月和一年"`
	ProductKey string `json:"productKey" v:"required#产品key不能为空"`
	DeviceKey  string `json:"deviceKey" v:"required#设备key不能为空"`
	Properties string `json:"properties" v:"required#设备属性不能为空"`
	StartDate  string `json:"startDate" v:"required#开始时间不能为空"`
	EndDate    string `json:"endDate" v:"required#结束时间不能为空"`
}

type PolymerizeRes struct {
	DataValue float64 `json:"dataValue" dc:"属性值"`
	Date      string  `json:"date" dc:"属性值上报时间"`
}

type DeviceIndicatorPolymerizeRes struct {
	DataAverageValue float64 `json:"dataAverageValue" dc:"平均值"`
	DataMaxValue     float64 `json:"dataMaxValue" dc:"最大值"`
	DataMinValue     float64 `json:"dataMinValue" dc:"最小值"`
	Date             string  `json:"date" dc:"属性值上报时间"`
}
