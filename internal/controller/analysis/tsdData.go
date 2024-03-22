package analysis

import (
	"context"
	"sagooiot/api/v1/analysis"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var TsdData = cTsdData{}

type cTsdData struct {
}

// getDeviceIndicatorTrend 获取设备指标趋势
func (c *cTsdData) getDeviceIndicatorTrend(ctx context.Context, req *analysis.DeviceIndicatorTrendReq) (res *analysis.DeviceIndicatorTrendRes, err error) {

	trendReq := model.DeviceIndicatorTrendReq{
		ProductKey:       req.ProductKey,
		DeviceCode:       req.DeviceCode,
		DeviceProperties: req.DeviceProperties,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
	}
	data, err := service.AnalysisTsdData().GetDeviceIndicatorTrend(ctx, trendReq)
	if err != nil {
		return
	}
	res.Data = data

	return res, err
}

// getDeviceIndicatorTrend 获取设备指标聚合
func (c *cTsdData) getDeviceIndicatorPolymerize(ctx context.Context, req *analysis.DeviceIndicatorPolymerizeReq) (res *analysis.DeviceIndicatorPolymerizeRes, err error) {

	polymerizeReq := model.DeviceIndicatorPolymerizeReq{
		ProductKey:       req.ProductKey,
		DeviceCode:       req.DeviceCode,
		DeviceProperties: req.DeviceProperties,
		StartDate:        req.StartDate,
		EndDate:          req.EndDate,
		DateType:         req.DateType,
	}
	data, err := service.AnalysisTsdData().GetDeviceIndicatorPolymerize(ctx, polymerizeReq)
	if err != nil {
		return
	}
	res.Data = data

	return res, err
}
