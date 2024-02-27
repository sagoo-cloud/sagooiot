package analysis

import (
	"context"
	"sagooiot/api/v1/analysis"
	"sagooiot/internal/service"
)

var Device = cDevice{}

type cDevice struct{}

// GetDeviceDataTotalCount 获取设备消息总数
func (c *cDevice) GetDeviceDataTotalCount(ctx context.Context, req *analysis.DeviceDataTotalCountReq) (res *analysis.DeviceDataTotalCountRes, err error) {
	number, err := service.AnalysisDevice().GetDeviceDataTotalCount(ctx, req.DateType)
	if err != nil {
		return
	}
	res = &analysis.DeviceDataTotalCountRes{
		Number: number,
	}
	return
}

// GetDeviceOnlineOfflineCount 设备在线离线统计, 设备总数，在线数，离线数
func (c *cDevice) GetDeviceOnlineOfflineCount(ctx context.Context, req *analysis.DeviceOnlineOfflineCountReq) (res *analysis.DeviceOnlineOfflineCountRes, err error) {
	data, err := service.AnalysisDevice().GetDeviceOnlineOfflineCount(ctx)
	if err != nil {
		return
	}
	res = &analysis.DeviceOnlineOfflineCountRes{
		Data: data,
	}
	return
}

// GetDeviceDataCount 设备数据统计，按年、月、日三种类型
func (c *cDevice) GetDeviceDataCount(ctx context.Context, req *analysis.DeviceDataCountReq) (res *analysis.DeviceDataCountRes, err error) {
	data, err := service.AnalysisDevice().GetDeviceDataCountList(ctx, req.DateType)
	if err != nil {
		return
	}
	res = &analysis.DeviceDataCountRes{
		Data: data,
	}
	return
}

//// GetDeviceDataCountByMonthDay  指定月份1-30日设备消息量统计数据
//func (c *cDevice) GetDeviceDataCountByMonthDay(ctx context.Context, req *analysis.DeviceDataCountByMonthDayReq) (res *analysis.DeviceDataCountByMonthDayRes, err error) {
//	data, err := service.AnalysisDevice().GetCountDeviceDataCountList(ctx, req.Month)
//	if err != nil {
//		return
//	}
//	res = &analysis.DeviceDataCountByMonthDayRes{
//		Data: data,
//	}
//	return
//}
//
//// GetDeviceDataCountByDayHour  按日每小时设备消息统计
//func (c *cDevice) GetDeviceDataCountByDayHour(ctx context.Context, req *analysis.DeviceDataCountByDayHourReq) (res *analysis.DeviceDataCountByDayHourRes, err error) {
//	data, err := service.AnalysisDevice().GetCountDeviceDataCountList(ctx, req.Day)
//	if err != nil {
//		return
//	}
//	res = &analysis.DeviceDataCountByDayHourRes{
//		Data: data,
//	}
//	return
//}
