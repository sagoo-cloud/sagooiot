package main

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	gplugin "github.com/hashicorp/go-plugin"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	extend "github.com/sagoo-cloud/sagooiot/extend/module"
	"github.com/sagoo-cloud/sagooiot/extend/sdk"
	"github.com/sagoo-cloud/sagooiot/plugins/notice/dingding/internal"
	"net/rpc"
)

type Options struct {
	PayloadURL string
	Secret     string
	Subject    string
	Body       string
}

//NoticeDingding 实现
type NoticeDingding struct{}

func (NoticeDingding) Info() model.ModuleInfo {
	var res = model.ModuleInfo{}
	res.Name = "dingding"
	res.Title = "Ding Ding"
	res.Author = "Microrain"
	res.Intro = "通过钉钉方式发送通知"
	res.Version = "0.01"
	return res
}

func (NoticeDingding) Send(data []byte) (res model.JsonRes) {
	//解析通知数据
	nd, err := sdk.DecodeNoticeData(data)
	if err != nil {
		res.Code = 2
		res.Message = "插件数据解析失败"
		res.Data = err.Error()
		return res
	}

	// 设定相关参数
	opts := []internal.Option{
		internal.AppKey(gconv.String(nd.Config["AppKey"])),
		internal.AppSecret(gconv.String(nd.Config["AppSecret"])),
		internal.AgentID(gconv.String(nd.Config["AgentID"])),
	}
	ding := internal.GetDingdingChannel(opts...)
	accessToken, err := ding.GetAccessToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ding.SendTextMessage(accessToken, "", "", "", "Hello, world!", "")
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

//DingdingPlugin 插件接口实现
type DingdingPlugin struct{}

//Server 此方法由插件进程延迟调
func (DingdingPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &extend.NoticeRPCServer{Impl: new(NoticeDingding)}, nil
}

// Client 此方法由宿主进程调用
func (DingdingPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &extend.NoticeRPC{Client: c}, nil
}

func main() {
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: extend.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

var pluginMap = map[string]gplugin.Plugin{
	"dingding": new(DingdingPlugin),
}
