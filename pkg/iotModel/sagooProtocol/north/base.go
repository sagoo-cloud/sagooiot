package north

import (
	"context"
)

// Message 北向消息通用结构体
type Message struct {
	Meta       map[string]string `json:"meta"`       //消息元数据，里面的字段暂时为空
	MessageId  string            `json:"messageId"`  // 消息id，`string`类型,消息唯一标识
	ProductKey string            `json:"productKey"` // 产品`key`，`string`类型
	DeviceKey  string            `json:"deviceKey"`  //设备`key`，`string`类型
	Data       interface{}       `json:"data"`       // 消息体，里面的字段根据不同的消息类型会有不同的结构体
}

// WriteMessage 北向消息发送
func WriteMessage(ctx context.Context, topic string, meta map[string]string, productKey, deviceKey string, data interface{}) {
	//m := Message{
	//	Meta:       meta,
	//	MessageId:  guid.S(),
	//	ProductKey: productKey,
	//	DeviceKey:  deviceKey,
	//	Data:       data,
	//}
	//payload, _ := json.Marshal(m)
	//if err := mqtt.Publish(topic, payload); err != nil {
	//	g.Log().Errorf(ctx, "publish north-message error: %s, topic:%s, message:%s", err, topic, string(payload))
	//}
}
