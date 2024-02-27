package set

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
	//  /sys/${productKey}/${deviceKey}/thing/service/property/set_reply
	if err = core.RegisterSubTopicHandler(sagooProtocol.PropertySetRegisterSubResponseTopic, consts.MsgTypePropertySetReply, PropertySet); err != nil {
		return err
	}
	tunnelBase.RegisterModelType(tunnelBase.UpSetProperty, tunnelBase.ModelType{
		LogType:          consts.MsgTypePropertySetReply,
		GetTopicWithInfo: GetTopicWithInfo,
		Handle:           PropertySet,
	})
	return nil
}

// 属性设置
func PropertySet(ctx context.Context, data topicModel.TopicHandlerData) error {
	var propertySet sagooProtocol.PropertySetRes
	if propertySetErr := json.Unmarshal(data.PayLoad, &propertySet); propertySetErr != nil {
		g.Log().Debugf(ctx, "parse data error: %v, topic:%s, message:%s, message ignored", propertySetErr, data.Topic, string(data.PayLoad))
		return propertySetErr
	}
	var setPropertyRes sagooProtocol.PropertySetRes
	if reportDataErr := json.Unmarshal(data.PayLoad, &setPropertyRes); reportDataErr != nil {
		g.Log().Debugf(ctx, "parse data error: %v, topic:%s, message:%s, message ignored", reportDataErr, data.Topic, string(data.PayLoad))
		return reportDataErr
	}
	if setPropertyRes.Id == "" {
		// 忽略Id不存在的消息
		return errors.New("ignore")
	}
	funcKey, _, responseChan, err := baseLogic.GetCallInfoById(ctx, setPropertyRes.Id)
	if err != nil {
		g.Log().Debugf(ctx, "get call info for service(id:%s) call error:%s", setPropertyRes.Id, err.Error())
		return err
	}
	var responseData = make(map[string]interface{})
	for _, function := range data.DeviceDetail.TSL.Functions {
		if function.Key == funcKey {
			for _, node := range function.Outputs {
				for key, value := range setPropertyRes.Data {
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

func GetTopicWithInfo(deviceKey, productKey, identity string) string {
	return fmt.Sprintf(strings.ReplaceAll(sagooProtocol.PropertySetRegisterSubResponseTopic, "+", "%s"), productKey, deviceKey)
}
