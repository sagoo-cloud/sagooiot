package module

import (
	"fmt"
	gplugin "github.com/hashicorp/go-plugin"
	"github.com/sagoo-cloud/sagooiot/extend/model"
	"net/rpc"
)

// Protocol 协议解析插件接口
type Protocol interface {
	Info() model.ModuleInfo
	Encode(args interface{}) (string, error)
	Decode(data []byte, dataIdent string) (string, error)
}

// ProtocolRPC 基于RPC实现
type ProtocolRPC struct {
	Client *rpc.Client
}

func (g *ProtocolRPC) Info() model.ModuleInfo {
	var resp model.ModuleInfo
	err := g.Client.Call("Plugin.Info", new(interface{}), &resp)
	if err != nil {
		//希望接口返回错误
		//这里没有太多其他选择。
		fmt.Println("==========")
		panic(err)
	}
	return resp
}
func (g *ProtocolRPC) Encode(args interface{}) (string, error) {
	var resp string
	err := g.Client.Call("Plugin.Encode", args, &resp)
	if err != nil {
		//希望接口返回错误
		//这里没有太多其他选择。
		//panic(err)
	}

	return resp, nil
}
func (g *ProtocolRPC) Decode(data []byte, dataIdent string) string {
	var resp string
	err := g.Client.Call("Plugin.Decode", data, &resp)
	if err != nil {
		//希望接口返回错误
		//这里没有太多其他选择。
		panic(err)
	}
	return resp
}

// ProtocolRPCServer  GreeterRPC的RPC服务器，符合 net/rpc的要求
type ProtocolRPCServer struct {
	// 内嵌业务接口
	// 插件进程会将实现业务接口的对象赋值给Impl
	Impl Protocol
}

func (s *ProtocolRPCServer) Info(args interface{}, resp *model.ModuleInfo) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ProtocolRPCServer) Encode(args interface{}) (string, error) {
	return s.Impl.Encode(args)
}
func (s *ProtocolRPCServer) Decode(args []byte, dataIdent string) (string, error) {
	return s.Impl.Decode(args, dataIdent)
}

// ProtocolPlugin 插件的虚拟实现。用于PluginMap的插件接口。在运行时，来自插件实现的实际实现会覆盖
type ProtocolPlugin struct{}

// Server 此方法由插件进程延迟的调用
func (ProtocolPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	return &ProtocolRPCServer{}, nil
	//return interface{}, nil
}

// Client 此方法由宿主进程调用
func (ProtocolPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ProtocolRPC{Client: c}, nil
	//return interface{}, nil
}
