package main

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	gplugin "github.com/hashicorp/go-plugin"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	extend "github.com/sagoo-cloud/sagooiot/extend/module"
	"github.com/sagoo-cloud/sagooiot/extend/sdk"
	"github.com/sagoo-cloud/sagooiot/plugins/notice/mail/internal"
	"net/rpc"
)

var logger *glog.Logger

//NoticeMail 实现
type NoticeMail struct{}

func (NoticeMail) Info() model.ModuleInfo {
	var res = model.ModuleInfo{}
	res.Name = "mail"
	res.Title = "电子邮件通知"
	res.Author = "Microrain"
	res.Intro = "通过电子邮件发送通知"
	res.Version = "0.01"
	return res
}

func (NoticeMail) Send(data []byte) (res model.JsonRes) {

	//解析通知数据
	nd, err := sdk.DecodeNoticeData(data)
	if err != nil {
		res.Code = 2
		res.Message = "邮件插件数据解析失败"
		res.Data = err.Error()
		return res
	}

	// 设定相关参数
	opts := []internal.Option{
		internal.MailHost(gconv.String(nd.Config["MailHost"])),
		internal.MailPort(gconv.Int(nd.Config["MailPort"])),
		internal.MailUser(gconv.String(nd.Config["MailUser"])),
		internal.MailPass(gconv.String(nd.Config["MailPass"])),
	}
	m := internal.GetMailChannel(opts...)
	if err := m.Send(nd.Msg); err != nil {
		res.Code = 2
		res.Message = "邮件发送失败"
		res.Data = err.Error()
		return res
	}

	tmpData := gjson.New(nd)
	res.Code = 0
	res.Message = "邮件发送成功"
	res.Data = tmpData.MustToJsonString()
	return res
}

//MailPlugin 插件接口实现
type MailPlugin struct{}

//Server 此方法由插件进程延迟调
func (MailPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &extend.NoticeRPCServer{Impl: new(NoticeMail)}, nil
}

// Client 此方法由宿主进程调用
func (MailPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &extend.NoticeRPC{Client: c}, nil
}

func main() {
	gplugin.Serve(&gplugin.ServeConfig{
		HandshakeConfig: extend.HandshakeConfig,
		Plugins:         pluginMap,
	})
}

var pluginMap = map[string]gplugin.Plugin{
	"mail": new(MailPlugin),
}
