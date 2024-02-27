package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/consts"
	"sagooiot/network/core"
	"sagooiot/network/core/logic/baseLogic"
	tunnelBase "sagooiot/network/core/tunnel/base"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/iotModel/topicModel"
	"strings"
)

func Init() (err error) {
	//  /sys/${productKey}/${devicekey}/thing/service/${tsl.service.identifier}_reply
	if err = core.RegisterSubTopicHandler("/sys/+/+/thing/service/+", consts.MsgTypeFunctionReply, ServiceCallOutput); err != nil {
		return err
	}
	tunnelBase.RegisterModelType(tunnelBase.UpServiceOutput, tunnelBase.ModelType{
		LogType:          consts.MsgTypeFunctionReply,
		GetTopicWithInfo: GetTopicWithInfo,
		Handle:           ServiceCallOutput,
	})
	return nil
}

// 服务调用
func ServiceCallOutput(ctx context.Context, data topicModel.TopicHandlerData) error {
	topicInfo := strings.Split(data.Topic, "/")
	if !strings.HasSuffix(topicInfo[6], "reply") {
		// 忽略非服务调用
		return errors.New("ignore")
	}
	var serviceCallResult sagooProtocol.ServiceCallOutputRes
	if reportDataErr := json.Unmarshal(data.PayLoad, &serviceCallResult); reportDataErr != nil {
		g.Log().Errorf(ctx, "parse data error: %v, topic:%s, message:%s, message ignored", reportDataErr, data.Topic, string(data.PayLoad))
		return reportDataErr
	}
	if serviceCallResult.Id == "" {
		// 忽略Id不存在的消息
		return errors.New("ignore")
	}
	funcKey, _, responseChan, err := baseLogic.GetCallInfoById(ctx, serviceCallResult.Id)
	if err != nil {
		g.Log().Errorf(ctx, "get call info for service(id:%s) call error:%s", serviceCallResult.Id, err.Error())
		return err
	}
	var responseData = make(map[string]interface{})
	for _, function := range data.DeviceDetail.TSL.Functions {
		if function.Key == funcKey {
			for _, node := range function.Outputs {
				for key, value := range serviceCallResult.Data {
					if key == node.Key {
						responseData[key] = node.ValueType.ConvertValue(value)
					}
				}
			}
		}
	}
	responseChan <- responseData
	return nil
}

func GetTopicWithInfo(productKey, deviceKey, identity string) string {
	str := strings.ReplaceAll(sagooProtocol.ServiceCallRegisterSubResponseTopic, "+", "%s") + "_reply"
	return fmt.Sprintf(str, productKey, deviceKey, identity)
}
