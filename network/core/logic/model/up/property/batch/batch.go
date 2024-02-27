package batch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/service"
	"sagooiot/network/core"
	"sagooiot/network/core/tunnel/base"
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

	//网关属性处理
	if len(gatewayBatchReport.Params.Properties) > 0 {
		if err := handleProperties(ctx, data, gatewayBatchReport.Params.Properties); err != nil {
			return err
		}
	}

	//网关事件处理
	if len(gatewayBatchReport.Params.Events) > 0 {
		if err := handleEvents(ctx, data, gatewayBatchReport.Params.Events, ""); err != nil {
			return err
		}
	}

	//网关子设备处理
	for _, sub := range gatewayBatchReport.Params.SubDevices {
		if len(sub.Properties) > 0 {
			if err := handleProperties(ctx, data, sub.Properties); err != nil {
				return err
			}
		}
		if len(sub.Events) > 0 {
			if err := handleEvents(ctx, data, sub.Events, sub.Identity.DeviceKey); err != nil {
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
func handleProperties(ctx context.Context, data topicModel.TopicHandlerData, properties map[string]interface{}) error {
	reportDataInfo, err := service.DevTSLParse().HandleProperties(ctx, data.DeviceDetail, properties)
	if err != nil {
		return logError(ctx, "parse property error", err, data)
	}

	// 上报处理结果
	if len(reportDataInfo) > 0 {
		//if err := service.DevDataReport().Property(ctx, data.DeviceKey, reportDataInfo, subDeviceKey); err != nil {
		//	return logError(ctx, "report property error", err, data)
		//}
		north.WriteMessage(ctx, north.PropertyReportMessageTopic, nil, data.ProductKey, data.DeviceKey, iotModel.PropertyReportMessage{
			Properties: reportDataInfo,
		})

		// 检查报警规则
		if err := service.AlarmRule().Check(ctx, data.DeviceKey, data.DeviceKey, consts.AlarmTriggerTypeProperty, reportDataInfo); err != nil {
			return logError(ctx, "alarm check error", err, data)
		}
	}

	return nil
}

// handleEvents 处理事件上报
func handleEvents(ctx context.Context, data topicModel.TopicHandlerData, events map[string]sagooProtocol.EventNode, subDeviceKey string) error {
	resList, err := service.DevTSLParse().HandleEvents(ctx, data.DeviceDetail, events)
	if err != nil {
		return logError(ctx, "parse event error", err, data)
	}

	for _, e := range resList {
		// 上报事件
		if len(e.Param.Value) > 0 {
			//if err := service.DevDataReport().Event(ctx, data.DeviceKey, reportEventData, subDeviceKey); err != nil {
			//	return logError(ctx, "report event error", err, data)
			//}

			north.WriteMessage(ctx, north.EventReportMessageTopic, nil, data.ProductKey, data.DeviceDetail.Key, iotModel.EventReportMessage{
				EventId:   e.Key,
				Events:    e.Param.Value,
				Timestamp: time.Now().UnixMilli(),
			})

			// 检查报警规则
			if err := service.AlarmRule().Check(ctx, data.DeviceKey, data.ProductKey, consts.AlarmTriggerTypeProperty, e); err != nil {
				return logError(ctx, "alarm check error", err, data)
			}
		}
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
