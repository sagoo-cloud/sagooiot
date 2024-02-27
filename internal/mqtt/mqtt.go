package mqtt

import (
	"context"
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"sagooiot/pkg/mqttclient"
)

func InitSystemMqtt() error {
	var ctx = gctx.New()
	var err error
	err = mqttclient.InitMqtt(&mqttclient.MqttConf{
		Addr:     g.Cfg().MustGet(ctx, "mqtt.addr").String(),
		ClientId: g.Cfg().MustGet(ctx, "mqtt.clientId").String(),
		UserName: g.Cfg().MustGet(ctx, "mqtt.auth.userName").String(),
		Password: g.Cfg().MustGet(ctx, "mqtt.auth.userPassWorld").String(),
	})
	return err
}

func Close() {
	mqttclient.Close()
}

func Publish(topic string, payload []byte) error {
	return mqttclient.Publish(topic, payload)
}

func PublishWithInterface(topic string, payload interface{}) error {
	data, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		return marshalErr
	}
	return mqttclient.Publish(topic, data)
}

// Subscribe todo 全局应该只有一个处理，需要设置消息隔离级别
func Subscribe(ctx context.Context, topic string, f func(context.Context, MQTT.Client, MQTT.Message)) error {
	return mqttclient.Subscribe(ctx, topic, func(client MQTT.Client, message MQTT.Message) {
		f(ctx, client, message)
	})
}
