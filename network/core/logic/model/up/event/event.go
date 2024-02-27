package event

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/service"
	"sagooiot/network/core"
	"sagooiot/network/core/logic/model/up/property/reporter"
	"sagooiot/network/core/tunnel/base"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/iotModel/sagooProtocol/north"
	"sagooiot/pkg/iotModel/topicModel"
	"strings"
	"time"
)

func Init() (err error) {
	//  /sys/${productKey}/${devicekey}/thing/event/${tsl.event.identifier}/post
	if err = core.RegisterSubTopicHandler(sagooProtocol.EventRegisterSubRequestTopic, consts.MsgTypeEvent, ReportEvent); err != nil {
		return err
	}
	base.RegisterModelType(base.UpEvent, base.ModelType{
		LogType:          consts.MsgTypeEvent,
		GetTopicWithInfo: GetTopicWithInfo,
		Handle:           ReportEvent,
	})
	return nil
}

// 事件上报
func ReportEvent(ctx context.Context, data topicModel.TopicHandlerData) error {
	if strings.HasSuffix(data.Topic, "property/post") {
		// 这段特殊的逻辑是为了处理topic通配符的问题
		return reporter.ReportProperty(ctx, data)
	}
	topicInfo := strings.Split(data.Topic, "/")
	eventKey := topicInfo[6]
	if eventKey == "property" {
		// 忽略属性上报信息
		return errors.New("ignore")
	}
	var reportData sagooProtocol.ReportEventReq
	if reportDataErr := json.Unmarshal(data.PayLoad, &reportData); reportDataErr != nil {
		g.Log().Errorf(ctx, "parse data error: %v, topic:%s, message:%s, message ignored", reportDataErr, data.Topic, string(data.PayLoad))
		return reportDataErr
	}
	if reportData.Params.Value == nil || len(reportData.Params.Value) == 0 {
		g.Log().Printf(ctx, "event data is empty, topic:%s, message:%s, message ignored\n", data.Topic, string(data.PayLoad))
		return errors.New("event data is empty")
	}
	var reportEventData = model.ReportEventData{
		Key: eventKey,
		Param: model.ReportEventParam{
			Value:      map[string]any{},
			CreateTime: reportData.Params.CreateAt,
		},
	}
	for _, event := range data.DeviceDetail.TSL.Events {
		if event.Key == eventKey {
			for _, o := range event.Outputs {
				for k, v := range reportData.Params.Value {
					if k == o.Name {
						reportEventData.Param.Value[k] = o.ValueType.ConvertValue(v)
					}
				}
			}
		}
	}
	if reportEventErr := service.DevDataReport().Event(ctx, data.DeviceKey, reportEventData); reportEventErr != nil {
		g.Log().Errorf(ctx, "report event error: %v, topic:%s, message:%s, message ignored", reportEventErr, data.Topic, string(data.PayLoad))
		return reportEventErr
	}
	//北向事件上报消息
	north.WriteMessage(ctx, north.EventReportMessageTopic, nil, data.DeviceDetail.Product.Key, data.DeviceDetail.Key, iotModel.EventReportMessage{
		EventId:   eventKey,
		Events:    reportEventData.Param.Value,
		Timestamp: time.Now().UnixMilli(),
	})
	if alarmCheckErr := service.AlarmRule().Check(ctx, data.DeviceKey, data.DeviceKey, consts.AlarmTriggerTypeProperty, reportEventData); alarmCheckErr != nil {
		g.Log().Errorf(ctx, "alarm check error: %v, topic:%s, message:%s, message ignored", alarmCheckErr, data.Topic, string(data.PayLoad))
		return alarmCheckErr
	}
	if reportData.Sys.Ack == sagooProtocol.NeedAck {
		return mqtt.PublishWithInterface(
			fmt.Sprintf(strings.ReplaceAll(sagooProtocol.EventRegisterPubResponseTopic, "+", "%s")+"_reply", data.ProductKey, data.DeviceKey, eventKey),
			sagooProtocol.ReportEventReply{
				Code:    200,
				Data:    struct{}{},
				Id:      reportData.Id,
				Message: "success",
				Method:  fmt.Sprintf("thing.event.%s.post_reply", eventKey),
				Version: "1.0",
			},
		)
	}
	return nil
}

func GetTopicWithInfo(deviceKey, productKey, identity string) string {
	return fmt.Sprintf(strings.ReplaceAll(sagooProtocol.EventRegisterSubRequestTopic, "+", "%s"), productKey, deviceKey, identity)
}
