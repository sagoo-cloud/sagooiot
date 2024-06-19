package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/mqtt"
	"sagooiot/network/core/logic/baseLogic"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/gpool"
	"sagooiot/pkg/iotModel/sagooProtocol"
	"sagooiot/pkg/iotModel/topicModel"
	"sagooiot/pkg/jsinterpreter"
	"sagooiot/pkg/plugins"
	"strings"
	"sync"
)

type filterMsgFunc struct {
	name string
	f    func(message MQTT.Message) MQTT.Message
}

type (
	SubMap struct {
		sync.RWMutex
		// 订阅主题
		subTopics map[string]handleFunc
		// 订阅处理前filter
		subTopicBeforeHandlerChain []filterMsgFunc
		// 订阅处理后filter
		subTopicAfterHandlerChain []filterMsgFunc
	}
	handleFunc struct {
		logType string
		f       func(context.Context, topicModel.TopicHandlerData) error
	}
)

var subMapInfo = SubMap{
	subTopics:                  make(map[string]handleFunc),
	subTopicBeforeHandlerChain: make([]filterMsgFunc, 0),
	subTopicAfterHandlerChain:  make([]filterMsgFunc, 0),
}

// 控制携程数量
var gPool = gpool.NewGPool(10000)

// HandleMessage 所有订阅的入口
func (s *SubMap) HandleMessage(ctx context.Context, handleF handleFunc) func(context.Context, MQTT.Client, MQTT.Message) {
	return func(ctx context.Context, client MQTT.Client, message MQTT.Message) {
		gPool.Go(func(ctx context.Context) error {
			// 根据topic拿到deviceKey和productKey
			topicInfo := strings.Split(message.Topic(), "/")
			if len(topicInfo) < 3 {
				//todo 是否入库，前端展示
				return errors.New(fmt.Sprintf("topic:%s is illegal, message(%s) ignored", message.Topic(), string(message.Payload())))
			}
			if topicInfo[1] != "sys" && topicInfo[1] != "ota" {
				//todo 非sys开头的topic不处理
				return errors.New(fmt.Sprintf("topic:%s is not supported,message(%s) ignored", message.Topic(), string(message.Payload())))
			}
			productKey, deviceKey := topicInfo[2], topicInfo[3]
			if topicInfo[1] == "ota" {
				productKey, deviceKey = topicInfo[4], topicInfo[5]
			}
			res := string(message.Payload())

			logType := handleF.logType
			if len(topicInfo) == 8 && topicInfo[6] == "property" {
				logType = consts.MsgTypePropertyReport
			}
			// 忽略一些不需要处理的消息
			if len(topicInfo) == 8 && logType == consts.MsgTypeFunctionReply && !strings.HasSuffix(topicInfo[6], "reply") {
				g.Log().Infof(ctx, "handleF: topic:%s, message:%s, message ignored", message.Topic(), string(message.Payload()))
				return nil
			}

			// 处理设备应答
			if strings.HasSuffix(topicInfo[6], "reply") {
				var msg sagooProtocol.ServiceCallOutputRes
				json.Unmarshal([]byte(res), &msg)

				if info, ok := baseLogic.AsyncMapInfo.Info[msg.Id]; ok {
					info.Response <- gconv.Map(msg)
				}
			}

			// 获取设备详情，拿出来消息协议，然后按照产品定义的消息协议解析消息
			deviceInfo, err := dcache.GetDeviceDetailInfo(deviceKey)
			if err != nil {
				g.Log().Debugf(ctx, "device info error: %v, topic:%s, message:%s, message ignored", err.Error(), message.Topic(), string(message.Payload()))
				return nil
			}
			if deviceInfo == nil {
				g.Log().Debugf(ctx, "device info is nil, topic:%s, message:%s, message ignored", message.Topic(), string(message.Payload()))
				return nil
			}

			// 设备禁用不处理
			if deviceInfo.Status == model.DeviceStatusNoEnable {
				g.Log().Debug(ctx, deviceKey, "device is no enable")
				return nil
			}

			dcache.UpdateStatus(ctx, deviceInfo) //更新设备状态

			messageProtocol := deviceInfo.Product.MessageProtocol
			if messageProtocol != consts.DefaultProtocol && messageProtocol != "" {
				if plugins.GetProtocolPlugin() == nil {
					return nil
				}
				pluginData, err := plugins.GetProtocolPlugin().GetProtocolDecodeData(deviceInfo.Product.MessageProtocol, message.Payload())
				if err != nil {
					return errors.New(fmt.Sprintf("get plugin error: %v, deviceKey:%s, data:%s, message ignored", err, deviceKey, string(message.Payload())))

				}
				if pluginData.Code != 0 {
					return errors.New(fmt.Sprintf("plugin parse error: code:%d message:%s, deviceKey:%s, data:%s, message ignored", pluginData.Code, pluginData.Message, deviceKey, string(message.Payload())))

				}
				pluginDataByte, _ := json.Marshal(pluginData.Data)
				res = string(pluginDataByte)
			}

			// 如果有js脚本，根据js脚本处理解析后的数据，处理后的数据数据格式为默认的消息协议格式
			if deviceInfo.Product.ScriptInfo != "" {
				var runScriptErr error
				res, runScriptErr = jsinterpreter.RunScript(res, deviceInfo.Product.ScriptInfo)
				if runScriptErr != nil {
					return errors.New(fmt.Sprintf("runScriptErr error: %v, topic:%s, message:%s, message ignored", runScriptErr, message.Topic(), string(message.Payload())))
				}
			}

			go baseLogic.InertTdLog(ctx, logType, deviceKey, res)

			// 前置处理器处理
			for _, filter := range s.subTopicBeforeHandlerChain {
				message = filter.f(message)
			}
			// 真正的topic处理方法
			if err := handleF.f(ctx, topicModel.TopicHandlerData{
				Topic:        message.Topic(),
				ProductKey:   productKey,
				DeviceKey:    deviceKey,
				PayLoad:      []byte(res),
				DeviceDetail: deviceInfo,
			}); err != nil {
				if err.Error() != "ignore" {
					return errors.New(fmt.Sprintf("handleF error: %s, topic:%s, message:%s ", err.Error(), message.Topic(), string(message.Payload())))

				} else {
					g.Log().Infof(ctx, "handleF: %s, topic:%s, message:%s, message ignored", err.Error(), message.Topic(), string(message.Payload()))
					return nil
				}
			}

			// 后置处理器处理
			for _, filter := range s.subTopicAfterHandlerChain {
				message = filter.f(message)
			}
			return nil
		})
	}
}

func RegisterSubTopicHandler(topic, logType string, handler func(context.Context, topicModel.TopicHandlerData) error) error {
	subMapInfo.Lock()
	defer subMapInfo.Unlock()
	if _, ok := subMapInfo.subTopics[topic]; ok {
		return fmt.Errorf("topic %s already registered", topic)
	}
	subMapInfo.subTopics[topic] = handleFunc{f: handler, logType: logType}
	return nil
}

func StartSubscriber(ctx context.Context) error {
	for topic := range subMapInfo.subTopics {
		if err := mqtt.Subscribe(ctx, topic, subMapInfo.HandleMessage(ctx, subMapInfo.subTopics[topic])); err != nil {
			return err
		}
	}
	return nil
}
