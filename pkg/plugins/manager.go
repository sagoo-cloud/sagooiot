package plugins

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hashicorp/go-plugin"
	"os/exec"
	"path/filepath"
	"sagooiot/internal/consts"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/plugins/consts/PluginType"
	"sagooiot/pkg/plugins/module"
	"strings"
	"sync"
)

// PluginInfo 插件信息
type PluginInfo struct {
	ID     string
	Path   string
	Stats  bool
	Client *plugin.Client
}

// Manager 插件管理器， 为不同类型的插件，管理的生命周期
type Manager struct {
	Type        string                 // 管理器处理的插件类型的id, 如protocol, notice
	Glob        string                 // 全局的插件文件名
	Path        string                 // 插件路径
	Plugins     map[string]*PluginInfo // 插件信息列表
	initialized bool                   // 是否初始化
	pluginImpl  plugin.Plugin          // 插件实现虚拟接口
}

// Init 初始化插件管理器
func (m *Manager) Init() error {

	//发现插件绝对路径
	plugins, err := plugin.Discover(m.Glob, m.Path)
	if err != nil {
		return err
	}

	//获取所有插件信息
	for _, p := range plugins {
		_, file := filepath.Split(p)
		globAster := strings.LastIndex(m.Glob, "*")
		trim := m.Glob[0:globAster]
		id := strings.TrimPrefix(file, trim)

		//添加到插件信息
		m.Plugins[id] = &PluginInfo{
			ID:   id,
			Path: p,
		}
	}

	m.initialized = true

	return nil
}

// Launch 启动所有插件
func (m *Manager) Launch() error {
	for id, info := range m.Plugins {
		g.Log().Debugf(context.Background(), "注册插件 type=%s, id=%s, impl=%s", m.Type, id, info.Path)
		// 创建新的客户端
		// 以exec.Command方式启动插件进程，并创建宿主机进程和插件进程的连接
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: module.HandshakeConfig,
			Plugins:         m.pluginMap(id),
			Cmd:             exec.Command(info.Path),
		})

		if _, ok := m.Plugins[id]; !ok {
			// 如果没有找到，忽略？
			continue
		}
		m.Plugins[id].Client = client

	}

	return nil
}

// Dispose 释放插件资源
func (m *Manager) Dispose() {
	var wg sync.WaitGroup
	for _, p := range m.Plugins {
		wg.Add(1)
		go func(client *plugin.Client) {
			// 关闭client，释放相关资源，终止插件子程序的运行
			client.Kill()
			wg.Done()
		}(p.Client)
	}
	wg.Wait()
}

// GetInterface 获取插件接口
func (m *Manager) GetInterface(id string) (interface{}, error) {

	if _, ok := m.Plugins[id]; !ok {
		return nil, fmt.Errorf("在注册的插件中找不到插件ID:%s", id)
	}
	//获取注册插件客户端 plugin.Client
	client := m.Plugins[id].Client

	// 返回协议客户端，如rpc客户端或grpc客户端，用于后续通信
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	// 根据指定插件名称分配新实例
	raw, err := rpcClient.Dispense(id)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// pluginMap //插件名称到插件对象的映射关系
func (m *Manager) pluginMap(id string) map[string]plugin.Plugin {
	pMap := map[string]plugin.Plugin{}
	pMap[id] = m.pluginImpl
	return pMap
}

// GetPluginsConfig 获取插件配置
func getPluginsConfigData(pluginType, pluginName string) (res map[interface{}]interface{}, err error) {
	key := fmt.Sprintf(consts.PluginsTypeName, pluginType, pluginName)
	pcgData, err := cache.Instance().Get(context.Background(), key)
	if err != nil {
		return
	}
	if pcgData == nil {
		return nil, errors.New("插件缓存配置不存在")
	}
	err = gyaml.DecodeTo([]byte(pcgData.String()), &res)
	return
}

// StartPlugin 启用一个插件
func (m *Manager) StartPlugin(id string) (err error) {
	p := m.Plugins[id]
	if p != nil {
		p.Client.Kill()
	}
	path := g.Cfg().MustGet(context.Background(), "system.pluginsPath").String() + "/protocol-" + id
	//path := "./plugins/built/protocol-" + id
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: module.HandshakeConfig,
		Plugins:         m.pluginMap(id),
		Cmd:             exec.Command(path),
	})
	if err != nil {
		g.Log().Error(context.Background(), id, "插件加载出错", err.Error())
	}
	p.Client = client
	g.Log().Debugf(context.Background(), "注册插件 type=%s, id=%s, impl=%s", m.Type, id, p.Path)

	return
}

// StopPlugin 停用一个插件
func (m *Manager) StopPlugin(id string) (err error) {
	p := m.Plugins[id]
	if p != nil {
		p.Client.Kill()
	}
	return
}

// AddPlugin 添加一个插件
func (m *Manager) AddPlugin(Type string, id string) (err error) {
	p := m.Plugins[id]
	if p != nil {
		err = errors.New("插件已存在")
		return
	}
	path := ""
	switch Type {
	case PluginType.Protocol:
		path = pluginsFilePath + "/protocol-" + id
		break
	case PluginType.Notice:
		path = pluginsFilePath + "/notice-" + id
		break
	default:
		err = errors.New("无效的插件类型")
	}

	//添加到插件列
	m.Plugins[id] = &PluginInfo{
		ID:   id,
		Path: path,
	}

	return
}

// RemovePlugin 移除一个插件
func (m *Manager) RemovePlugin(id string) (err error) {
	p := m.Plugins[id]
	if p == nil {
		err = errors.New("插件不存在")
		return
	}
	// 关闭client，释放相关资源，终止插件子程序的运行
	p.Client.Kill()
	// 从插件列表中移除
	delete(m.Plugins, id)
	return
}
