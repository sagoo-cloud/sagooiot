package mqtt

import (
	"context"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/sagoo-cloud/sagooiot/network/pkg/mqttclient"
)

var systemMqttClient *mqttclient.MqttWrapperClient

func InitSystemMqtt() error {
	var ctx = gctx.New()
	var err error
	systemMqttClient, err = mqttclient.InitMqtt(&mqttclient.MqttConf{
		Addr:     g.Cfg().MustGet(ctx, "mqtt.addr").String(),
		ClientId: g.Cfg().MustGet(ctx, "mqtt.clientId").String(),
		UserName: g.Cfg().MustGet(ctx, "mqtt.auth.userName").String(),
		Password: g.Cfg().MustGet(ctx, "mqtt.auth.userPassWorld").String(),
	})
	return err
}

func Close() {
	systemMqttClient.Close()
}

func Publish(topic string, payload []byte) error {
	return systemMqttClient.Publish(topic, payload)
}

func Subscribe(ctx context.Context, topic string, f func(context.Context, MQTT.Client, MQTT.Message)) error {
	return systemMqttClient.Subscribe(ctx, topic, func(client MQTT.Client, message MQTT.Message) {
		f(ctx, client, message)
	})
}
