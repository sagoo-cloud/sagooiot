package main

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	gplugin "github.com/hashicorp/go-plugin"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	extend "github.com/sagoo-cloud/sagooiot/extend/module"
	"github.com/sagoo-cloud/sagooiot/extend/sdk"
	"github.com/sagoo-cloud/sagooiot/plugins/notice/wework/internal"
	"net/rpc"
)

type Options struct {
	PayloadURL string
	Secret     string
	Subject    string
	Body       string
}

//NoticeWework 实现
type NoticeWework struct{}

func (NoticeWework) Info() model.ModuleInfo {
	var res = model.ModuleInfo{}
	res.Name = "wework"
	res.Title = "企业微信通知"
	res.Author = "Microrain"
	res.Intro = "通过企业微信方式发送通知"
	res.Version = "0.01"
	return res
}

func (NoticeWework) Send(data []byte) (res model.JsonRes) {
	//解析通知数据
	nd, err := sdk.DecodeNoticeData(data)
	if err != nil {
		res.Code = 2
		res.Message = "插件数据解析失败"
		res.Data = err.Error()
		return res
	}

	var sendObjectList []model.NoticeSendObject
	err = gjson.DecodeTo(nd.Msg.Totag, &sendObjectList)
	if err != nil {
		g.Log().Error(context.TODO(), err)
		return
	}
	corpid := gconv.String(nd.Config["Corpid"])
	agentID := gconv.String(nd.Config["AgentID"])
	secret := gconv.String(nd.Config["Secret"])
	token := gconv.String(nd.Config["Token"])
	encodingAESKey := gconv.String(nd.Config["EncodingAESKey"])

	alarmService := internal.GetInstance(corpid, agentID, secret, token, encodingAESKey)
	for _, object := range sendObjectList {
		if object.Name == "wework" {
			toUser := object.Value
			content := nd.Msg.MsgBody
			g.Log().Debug(context.TODO(), toUser, content)
			data, err := alarmService.SendMessage(toUser, content)
			if err != nil {
				g.Log().Error(context.TODO(), err)
			}
			g.Log().Debug(context.TODO(), data)
		}
	}

	return
}

//WeworkPlugin 插件接口实现
type WeworkPlugin struct{}

//Server 此方法由插件进程延迟调
func (WeworkPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &extend.NoticeRPCServer{Impl: new(NoticeWework)}, nil
}

// Client 此方法由宿主进程调用
func (WeworkPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &extend.NoticeRPC{Client: c}, nil
}

func main() {
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: extend.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

var pluginMap = map[string]gplugin.Plugin{
	"wework": new(WeworkPlugin),
}
