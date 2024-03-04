package batch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/service"
	"sagooiot/network/core"
	"sagooiot/network/core/logic/baseLogic"
	"sagooiot/network/core/tunnel/base"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/iotModel/sagooProtocol/north"
	"sagooiot/pkg/iotModel/topicModel"
	"strings"
	"time"
)

func Init() (err error) {
	// /sys/<span class="math-inline">\{productKey\}/</span>{devicekey}/thing/event/property/pack/post
	if err = core.RegisterSubTopicHandler(sagooProtocol.BatchRegisterSubRequestTopic, consts.MsgTypeGatewayBatch, GatewayBatchReportProperty); err != nil {
		return err
	}
	base.RegisterModelType(base.UpBatch, base.ModelType{
		LogType:          consts.MsgTypeGatewayBatch,
		GetTopicWithInfo: GetTopicWithInfo,
		Handle:           GatewayBatchReportProperty,
	})
	return nil
}

// GatewayBatchReportProperty 网关批量属性上报
func GatewayBatchReportProperty(ctx context.Context, data topicModel.TopicHandlerData) error {
	var gatewayBatchReport sagooProtocol.GatewayBatchReq
	if err := json.Unmarshal(data.PayLoad, &gatewayBatchReport); err != nil {
		return logError(ctx, "parse data error", err, data)
	}

	//网关子设备处理
	for _, sub := range gatewayBatchReport.Params.SubDevices {

		subDevice, err := dcache.GetDeviceDetailInfo(sub.Identity.DeviceKey)
		if err != nil {
			continue
		}
		dcache.UpdateStatus(ctx, subDevice) //更新子设备状态

		if len(sub.Properties) > 0 {
			if err := handleProperties(ctx, data, subDevice, sub.Properties); err != nil {
				return err
			}
		}
		if len(sub.Events) > 0 {
			if err := handleEvents(ctx, data, subDevice, sub.Events); err != nil {
				return err
			}
		}
	}

	//网关确认响应处理
	return handleAcknowledgment(ctx, gatewayBatchReport, data)
}

// logError 记录错误日志并返回错误
func logError(ctx context.Context, message string, err error, data topicModel.TopicHandlerData) error {
	g.Log().Errorf(ctx, "%s: %v, topic:%s, message:%s, message ignored", message, err, data.Topic, string(data.PayLoad))
	return err
}

// handleProperties 处理属性上报
func handleProperties(ctx context.Context, data topicModel.TopicHandlerData, subDevice *model.DeviceOutput, properties map[string]interface{}) error {
	reportDataInfo, err := service.DevTSLParse().HandleProperties(ctx, subDevice, properties)
	if err != nil {
		return logError(ctx, "parse property error", err, data)
	}

	// 上报处理结果
	if len(reportDataInfo) > 0 {

		var reportData = new(sagooProtocol.ReportPropertyReq)
		reportData.Id = guid.S()
		reportData.Version = "1.0"
		reportData.Sys = sagooProtocol.SysInfo{
			Ack: 0,
		}
		reportData.Params = properties
		reportData.Method = "thing.event.property.post"
		// 上报数据存入日志库
		go baseLogic.InertTdLog(ctx, consts.MsgTypePropertyReport, subDevice.Key, reportData)

		north.WriteMessage(ctx, north.PropertyReportMessageTopic, nil, subDevice.ProductKey, subDevice.Key, iotModel.PropertyReportMessage{
			Properties: reportDataInfo,
		})

		// 检查报警规则
		if err := service.AlarmRule().Check(ctx, subDevice.ProductKey, subDevice.Key, consts.AlarmTriggerTypeProperty, reportDataInfo); err != nil {
			return logError(ctx, "handleProperties alarm check error", err, data)
		}
	}

	return nil
}

// handleEvents 处理事件上报
func handleEvents(ctx context.Context, data topicModel.TopicHandlerData, subDevice *model.DeviceOutput, events map[string]sagooProtocol.EventNode) error {
	for eventName, eventData := range events {
		var reportData = new(sagooProtocol.ReportPropertyReq)
		reportData.Id = guid.S()
		reportData.Version = "1.0"
		reportData.Sys = sagooProtocol.SysInfo{
			Ack: 0,
		}
		reportData.Params = eventData.Value
		reportData.Method = "thing.event." + eventName + ".post"

		// 上报数据存入日志库
		go baseLogic.InertTdLog(ctx, consts.MsgTypeEvent, subDevice.Key, reportData)

		north.WriteMessage(ctx, north.EventReportMessageTopic, nil, subDevice.ProductKey, subDevice.Key, iotModel.EventReportMessage{
			EventId:   eventName,
			Events:    eventData.Value,
			Timestamp: time.Now().UnixMilli(),
		})

	}
	return nil
}

// handleAcknowledgment 处理确认响应
func handleAcknowledgment(ctx context.Context, gatewayBatchReport sagooProtocol.GatewayBatchReq, data topicModel.TopicHandlerData) error {
	if gatewayBatchReport.Sys.Ack == sagooProtocol.NeedAck {
		return mqtt.PublishWithInterface(
			fmt.Sprintf(strings.ReplaceAll(sagooProtocol.BatchRegisterPubResponseTopic, "+", "%s"), data.ProductKey, data.DeviceKey),
			sagooProtocol.GatewayBatchReply{
				Code:    200,
				Data:    struct{}{},
				Id:      gatewayBatchReport.Id,
				Message: "success",
				Method:  "thing.event.property.pack.post",
				Version: "1.0",
			},
		)
	}
	return nil
}

// GetTopicWithInfo 获取带有设备信息的主题
func GetTopicWithInfo(productKey, deviceKey, identity string) string {
	return fmt.Sprintf(strings.ReplaceAll(sagooProtocol.BatchRegisterSubRequestTopic, "+", "%s"), productKey, deviceKey)
}
