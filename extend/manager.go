package extend

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gyaml"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/extend/module"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hashicorp/go-plugin"
)

type PluginInfo struct {
	ID     string
	Path   string
	Client *plugin.Client
}

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

// Manager 为不同类型的插件，管理的生命周期
type Manager struct {
	Type        string                 // 管理器处理的插件类型的id
	Glob        string                 // 全局的插件文件名
	Path        string                 // 插件路径
	Plugins     map[string]*PluginInfo // 插件信息列表
	initialized bool                   // 是否初始化
	pluginImpl  plugin.Plugin          // 插件实现虚拟接口
}

func (m *Manager) Init() error {

	//发现插件绝对路径
	plugins, err := plugin.Discover(m.Glob, m.Path)
	if err != nil {
		return err
	}

	//获取所有插件信息
	for _, p := range plugins {
		_, file := filepath.Split(p)
		globAsterix := strings.LastIndex(m.Glob, "*")
		trim := m.Glob[0:globAsterix]
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

func (m *Manager) Launch() error {

	for id, info := range m.Plugins {

		fmt.Printf("注册插件 type=%s, id=%s, impl=%s \n", m.Type, id, info.Path)
		// 创建新的客户端
		// 两种方式选其一
		// 以exec.Command方式启动插件进程，并创建宿主机进程和插件进程的连接
		// 或者使用Reattach连接到现有进程，需提供Reattach信息
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: module.HandshakeConfig,
			Plugins:         m.pluginMap(id),
			//创建新进程，或使用Reattach连接到现有进程中
			Cmd: exec.Command(info.Path),
		})

		if _, ok := m.Plugins[id]; !ok {
			// 如果没有找到，忽略？
			continue
		}
		pinfo := m.Plugins[id]
		pinfo.Client = client

	}

	return nil
}

func (m *Manager) Dispose() {
	var wg sync.WaitGroup
	for _, pinfo := range m.Plugins {
		wg.Add(1)

		go func(client *plugin.Client) {
			// 关闭client，释放相关资源，终止插件子程序的运行
			client.Kill()
			wg.Done()
		}(pinfo.Client)
	}

	wg.Wait()

}

func (m *Manager) GetInterface(id string) (interface{}, error) {

	if _, ok := m.Plugins[id]; !ok {
		return nil, errors.New("在注册的插件中找不到插件ID！")
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

func getPluginsConfigData(pluginType, pluginName string) (res map[interface{}]interface{}, err error) {
	key := "plugins" + pluginType + pluginName
	fmt.Println(key)
	pcgData, err := g.Redis().Do(context.TODO(), "GET", key)

	err = gyaml.DecodeTo([]byte(pcgData.String()), &res)

	return
}
