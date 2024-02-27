package plugins

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hashicorp/go-plugin"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/module"
)

var pluginsFilePath = "../plugins/built"

// NewManager 创建插件管理器
func NewManager(ptype, glob, dir string, pluginImpl plugin.Plugin) *Manager {
	manager := &Manager{
		Type:       ptype,
		Glob:       glob,
		Path:       dir,
		Plugins:    map[string]*PluginInfo{},
		pluginImpl: pluginImpl,
	}
	return manager
}

// pluginInit 初始化插件管理器
func pluginInit(sysPluginType string) (pm *Manager, err error) {
	// 静态目录设置
	pluginsPath := g.Cfg().MustGet(context.TODO(), "system.pluginsPath").String()
	if pluginsPath != "" {
		pluginsFilePath = pluginsPath
	}

	switch sysPluginType {
	case PluginType.Notice:
		pm = NewManager(sysPluginType, PluginType.Notice+"-*", pluginsFilePath, &module.NoticePlugin{})
		defer pm.Dispose()

		break
	case PluginType.Protocol:
		pm = NewManager(sysPluginType, PluginType.Protocol+"-*", pluginsFilePath, &module.ProtocolPlugin{})
		defer pm.Dispose()

		break
	default:
		err = gerror.New("无效的插件类型")
		return
	}

	//初始化管理器
	err = pm.Init()
	//启动所有插件
	err = pm.Launch()

	return
}
