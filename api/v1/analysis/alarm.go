package analysis

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

// DeviceAlertCountByYearMonthReq 按年度每月设备告警数统计
type DeviceAlertCountByYearMonthReq struct {
	g.Meta `path:"/deviceAlertCountByYearMonth" method:"get" summary:"按年度每月设备告警数统计" tags:"IOT数据分析"`
	Year   string `json:"year" v:"required#日期不能为空" dc:"日期：年-yyyy"`
}
type DeviceAlertCountByYearMonthRes struct {
	Data []model.CountData `json:"data" dc:"月设备告警计数"`
}

// DeviceAlertCountByMonthDayReq 按月度每日设备告警数统计
type DeviceAlertCountByMonthDayReq struct {
	g.Meta `path:"/deviceAlertCountByMonthDay" method:"get" summary:"按月度每日设备告警数统计" tags:"IOT数据分析"`
	Month  string `json:"month" v:"required#日期不能为空" dc:"日期：年-月 yyyy-MM"`
}
type DeviceAlertCountByMonthDayRes struct {
	Data []model.CountData `json:"data" dc:"日设备告警计数"`
}

// DeviceAlertCountByDayHourReq 按日每小时设备告警数统计
type DeviceAlertCountByDayHourReq struct {
	g.Meta `path:"/deviceAlertCountsByDayHour" method:"get" summary:"按日每小时设备告警数统计" tags:"IOT数据分析"`
	Day    string `json:"day" v:"required#日期不能为空" dc:"日期：年-月-日 yyyy-MM-dd"`
}
type DeviceAlertCountByDayHourRes struct {
	Data []model.CountData `json:"data" dc:"小时设备告警计数"`
}

// DeviceAlarmTotalCountReq 指定时间的告警总数统计，按年，月，日分别指定
type DeviceAlarmTotalCountReq struct {
	g.Meta   `path:"/deviceAlarmTotalCount" method:"get" summary:"告警总数统计（当年、当月、当日）" tags:"IOT数据分析"`
	DateType string `json:"dateType" v:"required#日期类型不能为空" dc:"日期类型：year 年，month 月，day 日"`
	Date     string `json:"date" v:"required#日期不能为空" dc:"日期：年 yyyy，月 yyyy-MM，日 yyyy-MM-dd"`
}
type DeviceAlarmTotalCountRes struct {
	Number int64 `json:"number" dc:"告警总数"`
}

// DeviceAlarmLevelCountReq 告警级别统计
type DeviceAlarmLevelCountReq struct {
	g.Meta   `path:"/deviceAlarmLevelCount" method:"get" summary:"告警按级别统计" tags:"IOT数据分析"`
	DateType string `json:"dateType" v:"required#日期类型不能为空" dc:"日期类型：year 年，month 月，day 日"`
	Date     string `json:"date" v:"required#日期不能为空" dc:"日期：年 yyyy，月 MM，日 dd"`
}
type DeviceAlarmLevelCountRes struct {
	Data []model.CountData `json:"data" dc:"告警级别统计"`
}
