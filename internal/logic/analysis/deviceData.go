package analysis

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/general"
)

type sAnalysisDeviceData struct {
}

func init() {
	service.RegisterAnalysisDeviceData(analysisDeviceDataNew())
}

func analysisDeviceDataNew() *sAnalysisDeviceData {
	return &sAnalysisDeviceData{}
}

// GetDeviceData 获取设备数据
func (s *sAnalysisDeviceData) GetDeviceData(ctx context.Context, reqData model.DeviceDataReq) (res []interface{}, err error) {
	deviceLogList := new(model.DeviceLogSearchOutput)
	result, total, currentPage, err := dcache.GetDataByPage(ctx, reqData.DeviceKey, reqData.PageNum, reqData.PageSize, nil, reqData.DateRange)
	if err != nil {
		return
	}

	deviceLogList.Total = total
	deviceLogList.CurrentPage = currentPage
	var logs []model.TdLog
	if err := gconv.Scan(result, &logs); err != nil {
		return nil, err
	}
	deviceLogList.List = logs

	// 基于物模型解析数据
	for _, log := range logs {
		tmp, err := service.DevTSLParse().ParseData(ctx, reqData.DeviceKey, []byte(log.Content))
		if err != nil {
			continue
		}
		res = append(res, tmp)
	}

	return
}

// GetDeviceDataForProductByLatest 获取产品下的所有设备最新一条数据
func (s *sAnalysisDeviceData) GetDeviceDataForProductByLatest(ctx context.Context, productKey string) (res []model.DeviceDataRes, err error) {
	deviceKeys := getDeviceKeys(ctx, productKey)
	for _, key := range deviceKeys {
		var deviceData model.DeviceDataRes
		data := dcache.GetDeviceDetailDataByLatest(ctx, key)
		if data != nil {
			deviceData.DeviceKey = key
			deviceData.DeviceData = data
			res = append(res, deviceData)
		}
	}
	return
}

// GetDeviceHistoryData 获取设备历史数据（来自TSD的数据）
func (s *sAnalysisDeviceData) GetDeviceHistoryData(ctx context.Context, reqData model.DeviceDataReq) (res []interface{}, err error) {

	return
}

// GetDeviceAlarmLogData 获取设备告警数据
func (s *sAnalysisDeviceData) GetDeviceAlarmLogData(ctx context.Context, reqData *general.SelectReq) (res interface{}, err error) {
	m := dao.AlarmLog.Ctx(ctx)
	data, err := general.ListByPage(ctx, m, reqData, []string{"device_key", "rule_name", "product_key"})
	res = data
	return
}
