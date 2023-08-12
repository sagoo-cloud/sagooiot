package extend

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	"github.com/sagoo-cloud/sagooiot/extend/module"
	"sync"
)

type SysPlugin struct {
	pluginManager *Manager
}

const (
	NoticePluginName   = "notice"
	ProtocolPluginName = "protocol"
)

var ins *SysPlugin
var once sync.Once

//GetNoticePlugin 构造方法
func GetNoticePlugin() *SysPlugin {
	once.Do(func() {
		ins = &SysPlugin{}
		pm, err := pluginInit("notice")
		if err != nil {
			g.Log().Error(context.TODO(), err.Error())
		}
		ins.pluginManager = pm
	})

	return ins
}

//GetProtocolPlugin 构造方法
func GetProtocolPlugin() *SysPlugin {
	once.Do(func() {
		ins = &SysPlugin{}
		pm, err := pluginInit("protocol")
		if err != nil {
			g.Log().Error(context.TODO(), err.Error())
		}
		ins.pluginManager = pm
	})

	return ins
}

//初始化处理协议插件
func pluginInit(sysPluginType string) (pm *Manager, err error) {
	// 静态目录设置
	pluginsPath := g.Cfg().MustGet(context.TODO(), "system.pluginsPath").String()
	//pluginsPath := "../plugins/built"
	switch sysPluginType {
	case "notice":
		pm = NewManager(sysPluginType, "notice-*", pluginsPath, &module.NoticePlugin{})
		defer pm.Dispose()

		break
	case "protocol":
		pm = NewManager(sysPluginType, "protocol-*", pluginsPath, &module.ProtocolPlugin{})
		defer pm.Dispose()

		break
	default:
		err = gerror.New("无效的插件类型")
		return
	}

	//defer ProtocolPlugin.Dispose()
	//初始化管理器
	err = pm.Init()
	//重启所有插件
	err = pm.Launch()
	//if len(pm.Plugins) > 0 {
	//	for key, _ := range pm.Plugins {
	//		data, e := pm.GetInterface(key)
	//		if e != nil {
	//			return
	//		}
	//		//将插件启动的信息存入数据库
	//		res := data.(module.Notice).Info()
	//		var pluginData sysModel.SysPluginsAddInput
	//		err = gconv.Scan(res, &pluginData)
	//		pluginData.Status = 1
	//		pluginData.Types = sysPluginType
	//		pluginData.StartTime = gtime.Datetime()
	//		//go service.SysPlugins().SaveSysPlugins(context.TODO(), pluginData)
	//	}
	//}

	return
}

// GetProtocolUnpackData 通过协议解析插件处理后，获取解析数据。protocolType 为协议名称
func (pm *SysPlugin) GetProtocolUnpackData(protocolType string, data []byte) (res string, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(protocolType)
	if err != nil {
		return
	}
	res, err = p.(module.Protocol).Decode(data, "")
	return
}

// NoticeSend 通过插件发送通知信息。noticeType 为通知插件名称；msg为通知内容
func (pm *SysPlugin) NoticeSend(noticeType string, msg model.NoticeInfoData) (res string, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(noticeType)
	if err != nil {
		return
	}

	var nd = new(model.NoticeData)
	nd.Msg = msg
	cfgData, err := getPluginsConfigData("notice", noticeType)
	if err != nil {
		return
	}
	nd.Config = cfgData
	ndJson := gjson.New(nd)
	//转为byte
	byteData := ndJson.MustToJson()

	sendRes := p.(module.Notice).Send(byteData)
	res, err = gjson.New(sendRes).ToJsonString()
	g.Log().Debug(context.TODO(), "通知发送结果：", res)
	return
}
