package module

import (
	gplugin "github.com/hashicorp/go-plugin"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	"net/rpc"
)

// Notice 通知服务插件接口
type Notice interface {
	Info() model.ModuleInfo
	Send(data []byte) model.JsonRes
}

// NoticeRPC 基于RPC实现
type NoticeRPC struct {
	Client *rpc.Client
}

func (g *NoticeRPC) Info() model.ModuleInfo {
	var resp model.ModuleInfo
	err := g.Client.Call("Plugin.Info", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
func (g *NoticeRPC) Send(data []byte) model.JsonRes {
	var resp model.JsonRes
	err := g.Client.Call("Plugin.Send", data, &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

// NoticeRPCServer  GreeterRPC的RPC服务器，符合 net/rpc的要求
type NoticeRPCServer struct {
	Impl Notice
}

func (s *NoticeRPCServer) Info(args interface{}, resp *model.ModuleInfo) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *NoticeRPCServer) Send(data []byte, resp *model.JsonRes) error {
	*resp = s.Impl.Send(data)
	return nil
}

// NoticePlugin 插件的虚拟实现。用于PluginMap的插件接口。在运行时，来自插件实现的实际实现会覆盖
type NoticePlugin struct{}

// Server 此方法由插件进程延迟的调用
func (NoticePlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	checkParentAlive()
	return &NoticeRPCServer{}, nil
}

// Client 此方法由宿主进程调用
func (NoticePlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &NoticeRPC{Client: c}, nil
}
