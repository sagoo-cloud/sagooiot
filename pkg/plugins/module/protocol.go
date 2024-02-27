package module

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	gplugin "github.com/hashicorp/go-plugin"
	"net/rpc"
	"os"
	"sagooiot/pkg/plugins/model"
	"time"
)

var parentPid int

func init() {
	parentPid = os.Getppid()
}

// Protocol 协议解析插件接口
type Protocol interface {
	Info() model.PluginInfo
	Encode(args model.DataReq) model.JsonRes
	Decode(data model.DataReq) model.JsonRes
}

// ProtocolRPC 基于RPC实现
type ProtocolRPC struct {
	Client *rpc.Client
}

func (p *ProtocolRPC) Info() model.PluginInfo {
	var resp model.PluginInfo
	err := p.Client.Call("Plugin.Info", new(interface{}), &resp)
	if err != nil {
		//希望接口返回错误
		//这里没有太多其他选择。
		g.Log().Debug(context.Background(), err.Error())
		//panic(err)
	}
	return resp
}
func (p *ProtocolRPC) Encode(args model.DataReq) model.JsonRes {
	var resp model.JsonRes
	err := p.Client.Call("Plugin.Encode", args, &resp)
	//fmt.Println(args)
	if err != nil {
		//希望接口返回错误
		//这里没有太多其他选择。
		//panic(err)
		resp.Code = 1
		resp.Message = "protocol.go Encode " + err.Error()
		return resp
	}

	return resp
}
func (p *ProtocolRPC) Decode(data model.DataReq) model.JsonRes {
	var resp model.JsonRes
	err := p.Client.Call("Plugin.Decode", data, &resp)
	if err != nil {
		//希望接口返回错误
		//这里没有太多其他选择。
		//panic(err)
		resp.Code = 1
		resp.Message = err.Error()
		return resp
	}
	return resp
}

// ProtocolRPCServer  GreeterRPC的RPC服务器，符合 net/rpc的要求
type ProtocolRPCServer struct {
	// 内嵌业务接口
	// 插件进程会将实现业务接口的对象赋值给Impl
	Impl Protocol
}

func (s *ProtocolRPCServer) Info(args interface{}, resp *model.PluginInfo) error {
	*resp = s.Impl.Info()
	return nil
}
func (s *ProtocolRPCServer) Encode(args model.DataReq, resp *model.JsonRes) error {
	*resp = s.Impl.Encode(args)
	return nil
}
func (s *ProtocolRPCServer) Decode(args model.DataReq, resp *model.JsonRes) error {
	*resp = s.Impl.Decode(args)
	return nil
}

// ProtocolPlugin 插件的虚拟实现。用于PluginMap的插件接口。在运行时，来自插件实现的实际实现会覆盖
type ProtocolPlugin struct{}

// Server 此方法由插件进程延迟的调用
func (p *ProtocolPlugin) Server(*gplugin.MuxBroker) (interface{}, error) {
	checkParentAlive()
	return &ProtocolRPCServer{}, nil
	//return interface{}, nil
}

// Client 此方法由宿主进程调用
func (p *ProtocolPlugin) Client(b *gplugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &ProtocolRPC{Client: c}, nil
	//return interface{}, nil
}

// checkParentAlive 检查父进程(也就是 client )是否退出，如果退出了，自己也需要退出。
func checkParentAlive() {
	go func() {
		for {
			if parentPid == 1 || os.Getppid() != parentPid {
				fmt.Println("parent no alive, exit")
				os.Exit(0)
			}
			_, err := os.FindProcess(parentPid)
			if err != nil {
				fmt.Println("parent no alive, exit")
				os.Exit(0)
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
