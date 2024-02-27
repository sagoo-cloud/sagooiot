package plugins

import (
	"sagooiot/pkg/plugins/consts/PluginType"
)

// GetNoticePlugin 获取通知插件
func GetNoticePlugin() *SysPlugin {
	ins := GetPlugin(PluginType.Notice)
	return ins
}
