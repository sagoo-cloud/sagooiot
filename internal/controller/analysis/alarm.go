package analysis

import (
	"context"
	"sagooiot/api/v1/analysis"
	"sagooiot/internal/service"
)

var Alarm = cAlarm{}

type cAlarm struct{}

// GetDeviceAlarmLevelStats 获取告警级别统计
func (c *cAlarm) GetDeviceAlarmLevelStats(ctx context.Context, req *analysis.DeviceAlarmLevelCountReq) (res *analysis.DeviceAlarmLevelCountRes, err error) {
	data, err := service.AnalysisAlarm().GetAlarmLevelCount(ctx, req.DateType, req.Date)
	if err != nil {
		return
	}
	res = &analysis.DeviceAlarmLevelCountRes{
		Data: data,
	}
	return
}

// GetDeviceAlarmTotalCount 告警总数统计（当年、当月、当日）
func (c *cAlarm) GetDeviceAlarmTotalCount(ctx context.Context, req *analysis.DeviceAlarmTotalCountReq) (res *analysis.DeviceAlarmTotalCountRes, err error) {
	number, err := service.AnalysisAlarm().GetAlarmTotalCount(ctx, req.DateType, req.Date)
	if err != nil {
		return
	}
	res = &analysis.DeviceAlarmTotalCountRes{
		Number: number,
	}
	return
}

// GetDeviceAlertCountByYearMonth 按年度每月设备告警数统计
func (c *cAlarm) GetDeviceAlertCountByYearMonth(ctx context.Context, req *analysis.DeviceAlertCountByYearMonthReq) (res *analysis.DeviceAlertCountByYearMonthRes, err error) {
	data, err := service.AnalysisAlarm().GetDeviceAlertCountByYearMonth(ctx, req.Year)
	if err != nil {
		return
	}
	res = &analysis.DeviceAlertCountByYearMonthRes{
		Data: data,
	}
	return
}

// GetDeviceAlertCountByMonthDay 按月度每日设备告警数统计
func (c *cAlarm) GetDeviceAlertCountByMonthDay(ctx context.Context, req *analysis.DeviceAlertCountByMonthDayReq) (res *analysis.DeviceAlertCountByMonthDayRes, err error) {
	data, err := service.AnalysisAlarm().GetDeviceAlertCountByYearMonth(ctx, req.Month)
	if err != nil {
		return
	}
	res = &analysis.DeviceAlertCountByMonthDayRes{
		Data: data,
	}
	return
}

// GetDeviceAlertCountsByDayHour 按日每小时设备告警数统计
func (c *cAlarm) GetDeviceAlertCountsByDayHour(ctx context.Context, req *analysis.DeviceAlertCountByDayHourReq) (res *analysis.DeviceAlertCountByDayHourRes, err error) {
	data, err := service.AnalysisAlarm().GetDeviceAlertCountByYearMonth(ctx, req.Day)
	if err != nil {
		return
	}
	res = &analysis.DeviceAlertCountByDayHourRes{
		Data: data,
	}
	return
}
