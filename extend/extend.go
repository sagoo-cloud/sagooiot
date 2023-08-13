package extend

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	"github.com/sagoo-cloud/sagooiot/extend/module"
	"sync"
)

type SysPlugin struct {
	pluginManager *Manager
}

const (
	// NoticePluginTypeName 通知类插件
	NoticePluginTypeName = "notice"
	// ProtocolPluginTypeName 协议类插件
	ProtocolPluginTypeName = "protocol"
)

var ins *SysPlugin
var once sync.Once

// GetNoticePlugin 构造方法
func GetNoticePlugin() *SysPlugin {
	once.Do(func() {
		ins = &SysPlugin{}
		pm, err := pluginInit(NoticePluginTypeName)
		if err != nil {
			g.Log().Error(context.TODO(), err.Error())
		}
		ins.pluginManager = pm
	})

	return ins
}

// GetProtocolPlugin 构造方法
func GetProtocolPlugin() *SysPlugin {
	once.Do(func() {
		ins = &SysPlugin{}
		pm, err := pluginInit(ProtocolPluginTypeName)
		if err != nil {
			g.Log().Error(context.TODO(), err.Error())
		}
		ins.pluginManager = pm
	})

	return ins
}

// 初始化处理协议插件
func pluginInit(sysPluginType string) (pm *Manager, err error) {
	// 静态目录设置
	pluginsPath := g.Cfg().MustGet(context.TODO(), "system.pluginsPath").String()
	//pluginsPath := "../plugins/built"
	switch sysPluginType {
	case NoticePluginTypeName:
		pm = NewManager(sysPluginType, NoticePluginTypeName+"-*", pluginsPath, &module.NoticePlugin{})
		defer pm.Dispose()

		break
	case ProtocolPluginTypeName:
		pm = NewManager(sysPluginType, ProtocolPluginTypeName+"-*", pluginsPath, &module.ProtocolPlugin{})
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

	return
}

// GetProtocolPlugin 获取指协议名称的插件
func (pm *SysPlugin) GetProtocolPlugin(protocolName string) (obj module.Protocol, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(protocolName)
	if err != nil {
		return
	}
	obj = p.(module.Protocol)
	return
}

// GetProtocolUnpackData 通过协议解析插件处理后，获取解析数据。protocolType 为协议名称
// todo 需要标记数据协议子类型
func (pm *SysPlugin) GetProtocolUnpackData(protocolType string, data []byte) (res string, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(protocolType)
	if err != nil {
		return
	}

	var rd = model.DataReq{}
	rd.Data = data
	resData := p.(module.Protocol).Decode(rd)
	return gconv.String(resData), err
}

// NoticeSend 通过插件发送通知信息。noticeName 为通知插件名称；msg为通知内容
func (pm *SysPlugin) NoticeSend(noticeName string, msg model.NoticeInfoData) (res string, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(noticeName)
	if err != nil {
		return
	}

	var nd = new(model.NoticeData)
	nd.Msg = msg
	cfgData, err := getPluginsConfigData(NoticePluginTypeName, noticeName)
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
