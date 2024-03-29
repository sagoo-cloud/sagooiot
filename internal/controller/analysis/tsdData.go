package analysis

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/api/v1/analysis"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var TsdData = cTsdData{}

type cTsdData struct {
}

// GetDeviceIndicatorTrend 获取设备指标趋势
func (c *cTsdData) GetDeviceIndicatorTrend(ctx context.Context, req *analysis.DeviceIndicatorTrendReq) (res *analysis.DeviceIndicatorTrendRes, err error) {

	if req.DateRange == nil || len(req.DateRange) == 0 {
		searchDate := make([]string, 2)
		searchDate[0] = gtime.Now().Format("Y-m-d 00:00:00")
		searchDate[1] = gtime.Now().AddDate(0, 0, 1).Format("Y-m-d 23:59:59")
		req.DateRange = searchDate
	}
	trendReq := model.DeviceIndicatorTrendReq{
		ProductKey: req.ProductKey,
		DeviceKey:  req.DeviceKey,
		Properties: req.Properties,
		StartDate:  req.DateRange[0],
		EndDate:    req.DateRange[1],
	}
	data, err := service.AnalysisTsdData().GetDeviceIndicatorTrend(ctx, trendReq)
	if err != nil {
		return
	}
	res = &analysis.DeviceIndicatorTrendRes{
		Data: data,
	}
	return res, err
}

// GetDeviceIndicatorPolymerize 获取设备指标聚合
func (c *cTsdData) GetDeviceIndicatorPolymerize(ctx context.Context, req *analysis.DeviceIndicatorPolymerizeReq) (res *analysis.DeviceIndicatorPolymerizeRes, err error) {

	if req.DateRange == nil || len(req.DateRange) == 0 {
		searchDate := make([]string, 2)
		searchDate[0] = gtime.Now().Format("Y-m-d 00:00:00")
		searchDate[1] = gtime.Now().AddDate(0, 0, 1).Format("Y-m-d 23:59:59")
		req.DateRange = searchDate
	}
	polymerizeReq := model.DeviceIndicatorPolymerizeReq{
		ProductKey: req.ProductKey,
		DeviceKey:  req.DeviceKey,
		Properties: req.Properties,
		StartDate:  req.DateRange[0],
		EndDate:    req.DateRange[1],
		DateType:   req.DateType,
	}
	data, err := service.AnalysisTsdData().GetDeviceIndicatorPolymerize(ctx, polymerizeReq)
	if err != nil {
		return
	}
	res = &analysis.DeviceIndicatorPolymerizeRes{
		Data: data,
	}
	return res, err
}
