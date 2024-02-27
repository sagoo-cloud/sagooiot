package plugins

import (
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/model"
	"sagooiot/pkg/plugins/module"
	"sync"
)

type SysPlugin struct {
	pluginManager *Manager
}

var (
	onceMap = make(map[string]*sync.Once)
	insMap  = make(map[string]*SysPlugin) // 用于存放插件类型的单例
	mu      sync.Mutex
)

// GetPlugin 构造方法
func GetPlugin(pluginType string) *SysPlugin {
	mu.Lock()
	once, ok := onceMap[pluginType]
	if !ok {
		once = &sync.Once{}
		onceMap[pluginType] = once
	}
	mu.Unlock()

	once.Do(func() {
		instance := &SysPlugin{}
		pm, err := pluginInit(pluginType)
		if err != nil {
			g.Log().Error(context.TODO(), err.Error())
		}
		instance.pluginManager = pm

		mu.Lock()
		insMap[pluginType] = instance
		mu.Unlock()
	})

	return insMap[pluginType]
}

// GetProtocolByName 获取指定协议名称的插件
func (pm *SysPlugin) GetProtocolByName(protocolName string) (obj module.Protocol, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(protocolName)
	if err != nil {
		g.Log().Debug(context.Background(), err.Error())
		return
	}
	obj = p.(module.Protocol)
	return
}

// GetNoticeByName 获取指定通知名称的插件
func (pm *SysPlugin) GetNoticeByName(noticeName string) (obj module.Notice, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(noticeName)
	if err != nil {
		g.Log().Error(context.Background(), err.Error())
		return
	}
	obj = p.(module.Notice)
	return
}

// GetProtocolDecodeData 通过协议解析插件处理后，获取解析数据。protocolType 为协议名称
func (pm *SysPlugin) GetProtocolDecodeData(protocolName string, data []byte) (res model.JsonRes, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(protocolName)
	if err != nil {
		return
	}

	var rd = model.DataReq{}
	rd.Data = data
	resData := p.(module.Protocol).Decode(rd)
	return resData, err
}

// GetProtocolEncodeData 通过协议插件进行编码Encode处理后，获取下发的编码后的数据。protocolType 为协议名称
func (pm *SysPlugin) GetProtocolEncodeData(protocolName string, reqData model.DataReq) (res model.JsonRes, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(protocolName)
	if err != nil {
		return
	}

	//var rd = model.DataReq{}
	//rd.Data = data
	resData := p.(module.Protocol).Encode(reqData)
	return resData, err
}

// StartPlugin 启动插件
func (pm *SysPlugin) StartPlugin(pluginId string) (err error) {
	return pm.pluginManager.StartPlugin(pluginId)
}

// StopPlugin 启动插件
func (pm *SysPlugin) StopPlugin(pluginId string) (err error) {
	return pm.pluginManager.StopPlugin(pluginId)
}

// NoticeSend 通过插件发送通知信息。noticeName 为通知插件名称；msg为通知内容
func (pm *SysPlugin) NoticeSend(noticeName string, msg model.NoticeInfoData) (res model.JsonRes, err error) {
	//获取插件
	p, err := pm.pluginManager.GetInterface(noticeName)
	if err != nil {
		return
	}

	var nd = new(model.NoticeData)
	nd.Msg = msg
	cfgData, err := getPluginsConfigData(PluginType.Notice, noticeName)
	if err != nil {
		return
	}
	nd.Config = cfgData
	ndJson := gjson.New(nd)
	//转为byte
	byteData := ndJson.MustToJson()

	//g.Log().Debug(context.Background(), noticeName, "通知插件发送数据：", byteData)

	res = p.(module.Notice).Send(byteData)
	//res, err = gjson.New(sendRes).ToJsonString()
	//g.Log().Debug(context.Background(), "发送结果：", res)

	return
}
