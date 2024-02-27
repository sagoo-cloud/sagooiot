package set

import (
	"context"
	"encoding/json"
	"fmt"
	"sagooiot/internal/consts"
	"sagooiot/internal/mqtt"
	"sagooiot/network/core/logic/baseLogic"
	"sagooiot/pkg/iotModel"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/iotModel/sagooProtocol/north"
	"sagooiot/pkg/iotModel/topicModel"
	"strings"
	"time"

	"github.com/gogf/gf/v2/util/guid"
)

// 设备属性设置
func PropertySet(ctx context.Context, request topicModel.TopicDownHandlerData) (map[string]interface{}, error) {
	if request.DeviceDetail == nil {
		return nil, fmt.Errorf("device detail is nil")
	}
	if request.DeviceDetail.TSL == nil {
		return nil, fmt.Errorf("device tsl is nil")
	}

	requestDataMap := make(map[string]interface{})
	originRequestData := make(map[string]interface{})
	if err := json.Unmarshal(request.PayLoad, &originRequestData); err != nil {
		return nil, err
	}
	for k, v := range originRequestData {
		for _, property := range request.DeviceDetail.TSL.Properties {
			if property.Key == k {
				requestDataMap[k] = property.ValueType.ConvertValue(v)
			}
		}
	}
	r := sagooProtocol.PropertySetRequest{
		Id:      guid.S(),
		Version: "1.0.0",
		Params:  requestDataMap,
		Method:  "thing.service.property.set",
	}

	requestData, _ := json.Marshal(r)
	if err := mqtt.Publish(fmt.Sprintf(strings.ReplaceAll(sagooProtocol.PropertySetRegisterSubRequestTopic, "+", "%s"), request.DeviceDetail.Product.Key, request.DeviceDetail.Key), requestData); err != nil {
		return nil, err
	}
	//北向属性设置消息
	north.WriteMessage(ctx, north.PropertySetMessageTopic, nil, request.DeviceDetail.Product.Key, request.DeviceDetail.Key, iotModel.PropertySetMessage{
		Properties: requestDataMap,
		Timestamp:  time.Now().UnixMilli(),
	})
	baseLogic.InertTdLog(ctx, consts.MsgTypePropertySet, request.DeviceDetail.Key, r)
	response, err := baseLogic.SyncRequest(ctx, r.Id, "SetProperty", r, 0)
	if err != nil {
		return nil, err
	} else if res, covertOk := response.(map[string]interface{}); !covertOk {
		return nil, fmt.Errorf("set property  failed,response: %+v", response)
	} else {
		code, _ := res["code"].(int)
		//北向属性设置回复消息
		north.WriteMessage(ctx, north.PropertySetReplyMessageTopic, nil, request.DeviceDetail.Product.Key, request.DeviceDetail.Key, iotModel.PropertySetReplyMessage{
			Code:      code,
			Data:      res,
			Timestamp: time.Now().UnixMilli(),
		})
		return res, nil
	}
}
