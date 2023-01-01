package internal

import (
	"context"
	"encoding/json"
	"github.com/fastwego/wxwork/corporation"
	"github.com/fastwego/wxwork/corporation/apis/message"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/plugins/notice/wework/model"
)

type Alarm struct {
	Corp      *corporation.Corporation
	CorpApp   *corporation.App
	AppConfig corporation.AppConfig
}

/**
 * 建立私有变量
 */
var instance *Alarm

func GetInstance(corpid, agentID, secret, token, encodingAESKey string) *Alarm {
	instance = new(Alarm)
	// 加载应用的配置
	appConfigInfo, _ := g.Cfg().Get(context.TODO(), "wework.alarm", corporation.AppConfig{})
	p := gjson.New(appConfigInfo)

	if err := p.Scan(&instance.AppConfig); err != nil {
		g.Log().Error(context.TODO(), err)
	}

	instance.AppConfig.Token = token
	instance.AppConfig.AgentId = agentID
	instance.AppConfig.Secret = secret
	instance.AppConfig.EncodingAESKey = encodingAESKey

	instance.Corp = corporation.New(corporation.Config{Corpid: corpid})
	instance.CorpApp = instance.Corp.NewApp(instance.AppConfig)

	return instance
}

func (e *Alarm) SendMessage(toUser, content string) (interface{}, error) {
	if content == "" {
		return nil, gerror.New("发送的内容为空值")
	}

	sendMsg := new(model.Text)
	sendMsg.Agentid = e.AppConfig.AgentId
	sendMsg.Touser = toUser //"@all"
	sendMsg.Msgtype = "text"
	sendMsg.Text.Content = content
	sendMsg.DuplicateCheckInterval = 1800

	jsonByte := ToJson(sendMsg)

	resp, err := message.Send(e.CorpApp, jsonByte)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ToJson(data interface{}) []byte {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		g.Log().Error(context.TODO(), err)
	}
	return jsonByte
}
