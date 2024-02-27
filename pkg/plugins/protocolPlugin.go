package plugins

import "sagooiot/pkg/plugins/consts/PluginType"

// GetProtocolPlugin 获取协议插件
func GetProtocolPlugin() *SysPlugin {
	ins := GetPlugin(PluginType.Protocol)
	return ins
}
