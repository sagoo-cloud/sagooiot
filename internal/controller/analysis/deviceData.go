package analysis

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/analysis"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/general"
)

var DeviceData = cDeviceData{}

type cDeviceData struct{}

// GetDeviceData 获取设备最近的数据
func (c *cDeviceData) GetDeviceData(ctx context.Context, req *analysis.DeviceDataReq) (res *analysis.DeviceDataRes, err error) {
	var reqData model.DeviceDataReq
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return nil, err
	}

	data, err := service.AnalysisDeviceData().GetDeviceData(ctx, reqData)
	if err != nil {
		return
	}
	res = &analysis.DeviceDataRes{
		Data: data,
	}
	return
}

// GetDeviceDataForProductByLatest 获取产品下的所有设备最新一条数据
func (c *cDeviceData) GetDeviceDataForProductByLatest(ctx context.Context, req *analysis.DeviceDataForProductByLatestReq) (res *analysis.DeviceDataForProductByLatestRes, err error) {
	data, err := service.AnalysisDeviceData().GetDeviceDataForProductByLatest(ctx, req.ProductKey)
	if err != nil {
		return
	}
	res = &analysis.DeviceDataForProductByLatestRes{
		Data: data,
	}
	return
}

// GetDeviceDataForTsd 获取设备的时序数据
func (c *cDeviceData) GetDeviceDataForTsd(ctx context.Context, req *analysis.DeviceDataForTsdReq) (res *analysis.DeviceDataForTsdRes, err error) {
	var reqData general.SelectReq
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return nil, err
	}

	data, err := service.AnalysisDeviceDataTsd().GetDeviceData(ctx, reqData)
	if err != nil {
		return
	}
	res = &analysis.DeviceDataForTsdRes{
		Data: data,
	}
	return

}

// GetDeviceAlarmLogData 获取设备告警日志数据
func (c *cDeviceData) GetDeviceAlarmLogData(ctx context.Context, req *analysis.DeviceAlarmLogDataReq) (res *analysis.DeviceAlarmLogDataRes, err error) {
	var reqData = new(general.SelectReq)
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return nil, err
	}

	data, err := service.AnalysisDeviceData().GetDeviceAlarmLogData(ctx, reqData)
	if err != nil {
		return
	}
	res = &analysis.DeviceAlarmLogDataRes{
		Data: data,
	}
	return
}
