package proxy

import (
	"context"
	"sagooiot/internal/service"
	proxyModel "sagooiot/pkg/proxy/model"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetProductCount 获取产品数量统计
func GetProductCount(ctx context.Context) (out proxyModel.ProductCount, err error) {
	res, err := service.AnalysisProduct().GetProductCount(ctx)
	if err != nil {
		return
	}
	err = gconv.Scan(res, &out)
	return
}

// GetDeviceOnlineOfflineCount 获取设备在线离线统计
func GetDeviceOnlineOfflineCount(ctx context.Context) (out proxyModel.DeviceOnlineOfflineCount, err error) {
	res, err := service.AnalysisDevice().GetDeviceOnlineOfflineCount(ctx)
	if err != nil {
		return
	}
	err = gconv.Scan(res, &out)
	return
}

// GetDeviceDataTotalCount 获取设备消息总数统计,dataType :day,month,year
func GetDeviceDataTotalCount(ctx context.Context, dataType string) (number int64, err error) {
	return service.AnalysisDevice().GetDeviceDataTotalCount(ctx, dataType)
}

// GetAlarmLevelCount 告警级别统计
func GetAlarmLevelCount(ctx context.Context, dataType, date string) (out []proxyModel.CountData, err error) {
	res, err := service.AnalysisAlarm().GetAlarmLevelCount(ctx, dataType, date)
	if err != nil {
		return
	}
	err = gconv.Scan(res, &out)
	return
}
