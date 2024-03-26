// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/pkg/general"
)

type (
	IAnalysisAlarm interface {
		// GetDeviceAlertCountByYearMonth 按年度每月设备告警数统计
		GetDeviceAlertCountByYearMonth(ctx context.Context, year string) (res []model.CountData, err error)
		// GetDeviceAlertCountByMonthDay 按月度每日设备告警数统计
		GetDeviceAlertCountByMonthDay(ctx context.Context, month string) (res []model.CountData, err error)
		// GetDeviceAlertCountByDayHour 按日每小时设备告警数统计
		GetDeviceAlertCountByDayHour(ctx context.Context, day string) (res []model.CountData, err error)
		// GetAlarmTotalCount 告警总数统计（当年、当月、当日）,dataType :day,month,year ,date:2021 or 01 or21
		GetAlarmTotalCount(ctx context.Context, dataType, date string) (number int64, err error)
		// GetAlarmLevelCount 告警级别统计
		GetAlarmLevelCount(ctx context.Context, dataType, date string) (res []model.CountData, err error)
	}
	IAnalysisDevice interface {
		// GetDeviceDataTotalCount 获取设备消息总数统计,dataType :day,month,year
		GetDeviceDataTotalCount(ctx context.Context, dataType string) (number int64, err error)
		// GetDeviceOnlineOfflineCount 获取设备在线离线统计
		GetDeviceOnlineOfflineCount(ctx context.Context) (res model.DeviceOnlineOfflineCount, err error)
		// GetDeviceDataCountList 按年度每月设备消息统计，dataType 为统计数据类型 year:按年度,统计每个月的，month:按月份，统计每天的。当前的年与月
		GetDeviceDataCountList(ctx context.Context, dateType string) (res []model.CountData, err error)
	}
	IAnalysisDeviceData interface {
		// GetDeviceData 获取设备数据
		GetDeviceData(ctx context.Context, reqData model.DeviceDataReq) (res []interface{}, err error)
		// GetDeviceDataForProductByLatest 获取产品下的所有设备最新一条数据
		GetDeviceDataForProductByLatest(ctx context.Context, productKey string) (res []model.DeviceDataRes, err error)
		// GetDeviceHistoryData 获取设备历史数据（来自TSD的数据）
		GetDeviceHistoryData(ctx context.Context, reqData model.DeviceDataReq) (res []interface{}, err error)
		// GetDeviceAlarmLogData 获取设备告警数据
		GetDeviceAlarmLogData(ctx context.Context, reqData *general.SelectReq) (res interface{}, err error)
	}
	IAnalysisDeviceDataTsd interface {
		GetDeviceData(ctx context.Context, reqData general.SelectReq) (rs []interface{}, err error)
	}
	IAnalysisProduct interface {
		// GetDeviceCountForProduct 获取产品下的设备数量
		GetDeviceCountForProduct(ctx context.Context, productKey string) (number int, err error)
		// GetProductCount 获取产品数量统计
		GetProductCount(ctx context.Context) (res model.ProductCountRes, err error)
	}
	IAnalysisTsdData interface {
		// GetDeviceIndicatorTrend 获取指标趋势
		GetDeviceIndicatorTrend(ctx context.Context, req model.DeviceIndicatorTrendReq) (rs []*model.DeviceIndicatorTrendRes, err error)
		// GetDeviceIndicatorPolymerize 获取指标聚合
		GetDeviceIndicatorPolymerize(ctx context.Context, req model.DeviceIndicatorPolymerizeReq) (rs []*model.DeviceIndicatorPolymerizeRes, err error)
	}
)

var (
	localAnalysisAlarm         IAnalysisAlarm
	localAnalysisDevice        IAnalysisDevice
	localAnalysisDeviceData    IAnalysisDeviceData
	localAnalysisDeviceDataTsd IAnalysisDeviceDataTsd
	localAnalysisProduct       IAnalysisProduct
	localAnalysisTsdData       IAnalysisTsdData
)

func AnalysisAlarm() IAnalysisAlarm {
	if localAnalysisAlarm == nil {
		panic("implement not found for interface IAnalysisAlarm, forgot register?")
	}
	return localAnalysisAlarm
}

func RegisterAnalysisAlarm(i IAnalysisAlarm) {
	localAnalysisAlarm = i
}

func AnalysisDevice() IAnalysisDevice {
	if localAnalysisDevice == nil {
		panic("implement not found for interface IAnalysisDevice, forgot register?")
	}
	return localAnalysisDevice
}

func RegisterAnalysisDevice(i IAnalysisDevice) {
	localAnalysisDevice = i
}

func AnalysisDeviceData() IAnalysisDeviceData {
	if localAnalysisDeviceData == nil {
		panic("implement not found for interface IAnalysisDeviceData, forgot register?")
	}
	return localAnalysisDeviceData
}

func RegisterAnalysisDeviceData(i IAnalysisDeviceData) {
	localAnalysisDeviceData = i
}

func AnalysisDeviceDataTsd() IAnalysisDeviceDataTsd {
	if localAnalysisDeviceDataTsd == nil {
		panic("implement not found for interface IAnalysisDeviceDataTsd, forgot register?")
	}
	return localAnalysisDeviceDataTsd
}

func RegisterAnalysisDeviceDataTsd(i IAnalysisDeviceDataTsd) {
	localAnalysisDeviceDataTsd = i
}

func AnalysisProduct() IAnalysisProduct {
	if localAnalysisProduct == nil {
		panic("implement not found for interface IAnalysisProduct, forgot register?")
	}
	return localAnalysisProduct
}

func RegisterAnalysisProduct(i IAnalysisProduct) {
	localAnalysisProduct = i
}

func AnalysisTsdData() IAnalysisTsdData {
	if localAnalysisTsdData == nil {
		panic("implement not found for interface IAnalysisTsdData, forgot register?")
	}
	return localAnalysisTsdData
}

func RegisterAnalysisTsdData(i IAnalysisTsdData) {
	localAnalysisTsdData = i
}
