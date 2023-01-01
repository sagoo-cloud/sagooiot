package core

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	networkModel "github.com/sagoo-cloud/sagooiot/network/model"
	"time"
)

type MessagePropertyReporter struct {
	r messageRouter
}

func (m MessagePropertyReporter) DeviceDataHandle() {
	if err := service.TdLogTable().Insert(m.r.ctx, &model.TdLogAddInput{
		Ts:      gtime.Now(),
		Device:  m.r.deviceDetail.Key,
		Type:    consts.GetTopicType(consts.DataBusPropertyReport),
		Content: string(m.r.msg),
	}); err != nil {
		g.Log().Errorf(m.r.ctx, "insert deviceLog failed, err: %w, message:%s, message ignored", err, string(m.r.msg))
	}

	if m.r.data == nil || len(m.r.data) == 0 {
		g.Log().Printf(m.r.ctx, "report data is empty, message:%s, message ignored\n", string(m.r.msg))
		return
	}

	var reportFormatData = make(map[string]interface{})
	for k, v := range m.r.data {
		for _, property := range m.r.deviceDetail.TSL.Properties {
			if property.Key == k {
				reportFormatData[k] = property.ValueType.ConvertValue(v)
				break
			}
		}
	}
	nowTime := time.Now().Unix()
	reportMessage := networkModel.ReportPropertyMessage{
		Common: networkModel.Common{
			DeviceKey: m.r.deviceDetail.Key,
			MessageId: fmt.Sprintf("%s-%s-%d", m.r.deviceDetail.Product.Key, m.r.deviceDetail.Key, nowTime),
			Timestamp: nowTime,
		},
		Properties: reportFormatData,
	}
	reportFormatDataByte, _ := json.Marshal(reportMessage)
	if propertyReportErr := mqtt.Publish(consts.GetDataBusWrapperTopic(m.r.deviceDetail.Product.Key, m.r.deviceDetail.Key, consts.DataBusPropertyReport), reportFormatDataByte); propertyReportErr != nil {
		g.Log().Errorf(m.r.ctx, "publish  data error: %w,message:%s, message ignored", propertyReportErr, string(m.r.msg))
		return
	}
}
