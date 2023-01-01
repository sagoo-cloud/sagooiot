package core

import (
	"context"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/extend"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	networkModel "github.com/sagoo-cloud/sagooiot/network/model"
	"strings"
	"time"
)

var subMap = map[consts.Topic]func(context.Context, MQTT.Client, MQTT.Message){
	//device 处理的topic
	consts.TopicDeviceData: deviceDataHandler,

	//dataBus  处理的topic
	consts.CommonDataBusPrefix + consts.DataBusOnline:         onlineHandler,
	consts.CommonDataBusPrefix + consts.DataBusOffline:        offlineHandler,
	consts.CommonDataBusPrefix + consts.DataBusPropertyReport: reportPropertiesHandler,
}

func StartSubscriber(ctx context.Context) error {
	for topic, f := range subMap {
		if err := mqtt.Subscribe(ctx, string(topic), f); err != nil {
			return err
		}
	}
	return nil
}

func deviceDataHandler(ctx context.Context, client MQTT.Client, message MQTT.Message) {
	topicInfo := strings.Split(message.Topic(), "/")
	if len(topicInfo) != 3 {
		g.Log().Error(ctx, fmt.Sprintf("topic:%s is illegal, message(%s) ignored", message.Topic(), string(message.Payload())))
		return
	}
	productKey, deviceKey := topicInfo[1], topicInfo[2]
	productDetail, productDetailErr := service.DevProduct().Get(ctx, productKey)
	if productDetailErr != nil || productDetail == nil {
		g.Log().Errorf(ctx, "find product info error: %w, topic:%s, message:%s, message ignored", productDetailErr, message.Topic(), string(message.Payload()))
		return
	}
	deviceDetail, err := service.DevDevice().Get(ctx, deviceKey)
	if err != nil {
		return
	}
	if deviceDetail == nil {
		g.Log().Errorf(ctx, "fail to find device, topic:%s, message:%s, message ignored", message.Topic(), string(message.Payload()))
		return
	}
	if deviceDetail != nil && deviceDetail.Status != consts.DeviceStatueOnline {
		onlineMessage, _ := json.Marshal(networkModel.DeviceOnlineMessage{
			DeviceKey: deviceKey,
			Timestamp: time.Now().Unix(),
		})
		if propertyReportErr := mqtt.Publish(consts.GetDataBusWrapperTopic(productKey, deviceKey, consts.DataBusOnline), onlineMessage); propertyReportErr != nil {
			g.Log().Errorf(ctx, "publish formate data error: %w, topic:%s, message:%s, message ignored", propertyReportErr, message.Topic(), string(message.Payload()))
			return
		}
	}
	messageProtocol := productDetail.MessageProtocol
	res := string(message.Payload())
	if messageProtocol != "" {
		if extend.GetProtocolPlugin() == nil {
			return
		}
		res, err = extend.GetProtocolPlugin().GetProtocolUnpackData(messageProtocol, message.Payload())
		if err != nil {
			g.Log().Errorf(ctx, "get plugin error: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
			return
		}
	}
	var reportData networkModel.DefaultMessageType
	if reportDataErr := json.Unmarshal([]byte(res), &reportData); reportDataErr != nil {
		g.Log().Errorf(ctx, "parse data error: %w, topic:%s, message:%s, message ignored", reportDataErr, message.Topic(), string(message.Payload()))
		return
	}
	messageRouter{
		ctx:          ctx,
		data:         reportData.Data,
		msgType:      reportData.DataType,
		deviceDetail: deviceDetail,
		msg:          message.Payload(),
	}.router()
}
