package mqttclient

import (
	"context"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
	"sync"
	"time"
)

var systemMqttClient *MqttWrapperClient

type MqttConf struct {
	Addr     string
	ClientId string
	UserName string
	Password string
}

type MqttWrapperClient struct {
	sync.RWMutex
	subTopic map[string]mqtt.MessageHandler
	c        mqtt.Client
}

func InitMqtt(c *MqttConf) (err error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(c.Addr)
	opts.SetClientID(c.ClientId + strconv.Itoa(time.Now().Nanosecond()))
	opts.SetAutoReconnect(true)
	opts.SetConnectRetryInterval(time.Second * 30)
	opts.SetUsername(c.UserName)
	opts.SetPassword(c.Password)
	opts.SetConnectRetry(true)
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		if systemMqttClient != nil {
			for topic, handler := range systemMqttClient.subTopic {
				if subErr := systemMqttClient.c.Subscribe(topic, 2, handler); subErr.Error() != nil {
					g.Log().Errorf(context.TODO(), "reconnect subscribe error: %s", subErr.Error())
				}
			}
		}
	})
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		g.Log().Debugf(context.TODO(), "MQTT连接断开... %s", err.Error())
	})
	err = getNewClient(opts)
	return
}

// getNewClient 获取新的mqtt客户端
func getNewClient(opts *mqtt.ClientOptions) (err error) {
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		systemMqttClient = &MqttWrapperClient{
			c:        mqttClient,
			subTopic: make(map[string]mqtt.MessageHandler),
		}
		return nil
	}
}

func Close() {
	systemMqttClient.c.Disconnect(uint(250))
}

func Publish(topic string, payload []byte) error {
	pubToken := systemMqttClient.c.Publish(topic, 2, false, payload)
	return pubToken.Error()
}

func Subscribe(ctx context.Context, topic string, h mqtt.MessageHandler) (err error) {
	qos := g.Cfg().MustGet(ctx, "mqtt.qos").Int()
	if subErr := systemMqttClient.c.Subscribe(topic, byte(qos), h); subErr.Error() != nil {
		g.Log().Errorf(ctx, "subscribe error: %s", subErr.Error())
		return subErr.Error()
	} else {
		systemMqttClient.Lock()
		systemMqttClient.subTopic[topic] = h
		systemMqttClient.Unlock()
		return nil
	}
}
