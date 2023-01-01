package core

import (
	"context"
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	networkModel "github.com/sagoo-cloud/sagooiot/network/model"
	"strings"
)

func onlineHandler(ctx context.Context, client MQTT.Client, message MQTT.Message) {
	msg := networkModel.DeviceOnlineMessage{}
	if err := json.Unmarshal(message.Payload(), &msg); err != nil {
		g.Log().Errorf(ctx, "onlineHandler error: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
	}
	if err := service.TdLogTable().Insert(ctx, &model.TdLogAddInput{
		Ts:      gtime.Now(),
		Device:  msg.DeviceKey,
		Type:    consts.GetTopicType(consts.DataBusOnline),
		Content: string(message.Payload()),
	}); err != nil {
		g.Log().Errorf(ctx, "insert deviceLog failed, err: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
	}
	if err := service.DevDevice().Online(ctx, msg.DeviceKey); err != nil {
		g.Log().Errorf(ctx, "onlineHandler set device online error: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
	}
}

func offlineHandler(ctx context.Context, client MQTT.Client, message MQTT.Message) {
	msg := networkModel.DeviceOfflineMessage{}
	if err := json.Unmarshal(message.Payload(), &msg); err != nil {
		g.Log().Errorf(ctx, "offlineHandler error: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
	}
	if err := service.TdLogTable().Insert(ctx, &model.TdLogAddInput{
		Ts:      gtime.Now(),
		Device:  msg.DeviceKey,
		Type:    consts.GetTopicType(consts.DataBusOffline),
		Content: string(message.Payload()),
	}); err != nil {
		g.Log().Errorf(ctx, "insert deviceLog failed, err: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
	}
	if err := service.DevDevice().Offline(ctx, msg.DeviceKey); err != nil {
		g.Log().Errorf(ctx, "onlineHandler set device online error: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
	}
}

func reportPropertiesHandler(ctx context.Context, client MQTT.Client, message MQTT.Message) {
	var reportMessage networkModel.ReportPropertyMessage
	if err := json.Unmarshal(message.Payload(), &reportMessage); err != nil {
		g.Log().Errorf(ctx, "unmarshal error: %w, topic:%s, message:%s, message ignored", err, message.Topic(), string(message.Payload()))
		return
	}
	if insertErr := service.TSLTable().Insert(ctx, reportMessage.DeviceKey, reportMessage.Properties); insertErr != nil {
		g.Log().Errorf(ctx, "insert error: %w, topic:%s, message:%s, message ignored", insertErr, message.Topic(), string(message.Payload()))
		return
	}
	topicNodes := strings.Split(message.Topic(), "/")
	if len(topicNodes) < 3 {
		return
	}
	if alarmCheckErr := service.AlarmRule().Check(ctx, topicNodes[2], reportMessage.DeviceKey, reportMessage.Properties); alarmCheckErr != nil {
		g.Log().Errorf(ctx, "alarm check error: %w, topic:%s, message:%s, message ignored", alarmCheckErr, message.Topic(), string(message.Payload()))
		return
	}
}
