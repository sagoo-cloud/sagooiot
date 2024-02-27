package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
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

// 服务调用
func ServiceCall(ctx context.Context, funcKey string, request topicModel.TopicDownHandlerData) (map[string]interface{}, error) {
	if request.DeviceDetail == nil {
		return nil, fmt.Errorf("device detail is nil")
	}
	if request.DeviceDetail.TSL == nil {
		return nil, fmt.Errorf("device tsl is nil")
	}
	requestDataMap := make(map[string]interface{})
	originRequestData := make(map[string]interface{})

	if request.PayLoad != nil {
		if err := json.Unmarshal(request.PayLoad, &originRequestData); err != nil {
			return nil, err
		}
	}

	for _, service := range request.DeviceDetail.TSL.Functions {
		if service.Key == funcKey {
			for _, param := range service.Inputs {
				for key, value := range originRequestData {
					if key == param.Key {
						requestDataMap[param.Key] = param.ValueType.ConvertValue(value)
					}
				}
			}
		}
	}
	r := sagooProtocol.ServiceCallRequest{
		Id:      guid.S(),
		Version: "1.0.0",
		Params:  requestDataMap,
		Method:  fmt.Sprintf("thing.service.%s", funcKey),
	}

	requestData, _ := json.Marshal(r)

	// 记录服务下发日志
	baseLogic.InertTdLog(ctx, consts.MsgTypeFunctionSend, request.DeviceDetail.Key, r)
	time.Sleep(time.Second * 1)
	g.Log().Debug(ctx, "service call request: %s", string(requestData))

	// 产品定义的传输协议支持 tcp/udp/mqtt_server/http/websocket 后面定义为变量
	if request.DeviceDetail.Product.TransportProtocol == "mqtt_server" {
		if err := mqtt.Publish(fmt.Sprintf(strings.ReplaceAll(sagooProtocol.ServiceCallRegisterSubRequestTopic, "+", "%s"), request.DeviceDetail.Product.Key, request.DeviceDetail.Key, funcKey), requestData); err != nil {
			return nil, err
		}
	} else if request.DeviceDetail.Product.TransportProtocol == "udp" || request.DeviceDetail.Product.TransportProtocol == "tcp" {
		// 如果是udp或者tcp，查询出通道然后通过通道下发，需要注意的是，仅仅支持服务端建立的通道。
		//todo 暂时不支持多节点部署，支持多节点部署的话，需要一个有效的查找连接的方法

		reqData, _ := json.Marshal(requestDataMap)
		if err := WriteTunnel(ctx, funcKey, request, reqData); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("transport protocol %s not support", request.DeviceDetail.Product.TransportProtocol)
	}
	//北向服务调用请求消息
	north.WriteMessage(ctx, north.ServiceCallMessageTopic, nil, request.DeviceDetail.Product.Key, request.DeviceDetail.Key, iotModel.ServiceCallMessage{
		ServiceId: funcKey,
		Params:    requestDataMap,
		Timestamp: time.Now().UnixMilli(),
	})
	response, err := baseLogic.SyncRequest(ctx, r.Id, funcKey, r, 0)
	if err != nil {
		return nil, err
	} else if res, covertOk := response.(map[string]interface{}); !covertOk {
		return nil, fmt.Errorf("invoke service %s failed,response: %+v", funcKey, response)
	} else {
		code, _ := res["code"].(int)
		//北向服务调用响应请求消息
		north.WriteMessage(ctx, north.ServiceReplyMessageTopic, nil, request.DeviceDetail.Product.Key, request.DeviceDetail.Key, iotModel.ServiceCallReplyMessage{
			ServiceId: funcKey,
			Code:      code,
			Data:      res,
			Timestamp: time.Now().UnixMilli(),
		})
		return res, nil
	}

}
