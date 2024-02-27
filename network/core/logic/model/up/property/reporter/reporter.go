package reporter

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/service"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/iotModel/sagooProtocol/north"
	"sagooiot/pkg/iotModel/topicModel"
	"strings"
)

func Init() (err error) {
	return nil
}

// ReportProperty 属性上报
func ReportProperty(ctx context.Context, data topicModel.TopicHandlerData) error {
	// 记录开始时间
	//start := time.Now()

	//对存在转义字符的进行全量替换
	payLoad := strings.ReplaceAll(string(data.PayLoad), "\\", "")

	var reportData sagooProtocol.ReportPropertyReq
	if reportDataErr := json.Unmarshal([]byte(payLoad), &reportData); reportDataErr != nil {
		g.Log().Errorf(ctx, "parse data error: %s, topic:%s, message:%s, message ignored", reportDataErr, data.Topic, string(data.PayLoad))
		return reportDataErr
	}

	//解析数据
	reportDataInfo, err := service.DevTSLParse().ParseData(ctx, data.DeviceKey, []byte(payLoad))
	if err != nil {
		return err
	}
	//北向属性上报消息
	north.WriteMessage(ctx, north.PropertyReportMessageTopic, nil, data.ProductKey, data.DeviceDetail.Key, iotModel.PropertyReportMessage{
		Properties: reportDataInfo,
	})

	//告警处理
	err = service.AlarmRule().Check(ctx, data.ProductKey, data.DeviceKey, consts.AlarmTriggerTypeProperty, reportDataInfo)
	if err != nil {
		g.Log().Errorf(ctx, "告警检测失败: %s", err.Error())
	}

	//记录结束时间
	//end := time.Now()
	//计算运行时间
	//taken := end.Sub(start)
	//g.Log().Debugf(ctx, "属性上报, deviceKey:%s - ReportProperty，程序用时：%s", data.DeviceKey, taken)
	if reportData.Sys.Ack == sagooProtocol.NeedAck {
		return mqtt.PublishWithInterface(
			fmt.Sprintf(strings.ReplaceAll(sagooProtocol.PropertyRegisterPubResponseTopic, "+", "%s"), data.ProductKey, data.DeviceKey),
			sagooProtocol.ReportPropertyReply{
				Code:    200,
				Data:    struct{}{},
				Id:      reportData.Id,
				Message: "success",
				Method:  "thing.event.property.post",
				Version: "1.0",
			},
		)
	} else {
		return nil
	}
}

func GetTopicWithInfo(deviceKey, productKey, identity string) string {
	return fmt.Sprintf(strings.ReplaceAll(sagooProtocol.PropertyRegisterSubRequestTopic, "+", "%s"), productKey, deviceKey)
}
