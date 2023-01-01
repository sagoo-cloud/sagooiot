package core

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	interModel "github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"sync"
)

var allServers sync.Map

func startServer(ctx context.Context, server *model.Server) error {
	svr, err := NewServer(server)
	if err != nil {
		return err
	}
	allServers.Store(server.Id, &Server{
		Server:   *server,
		Instance: svr,
	})
	err = svr.Open(ctx)
	if err != nil {
		return err
	}

	return nil
}

// LoadServers 启动时候加载服务
func LoadServers(ctx context.Context) error {
	// 这里限制10000，后面需要注意这里的参数显示
	total, networkServerListRes, networkServerListResErr := service.NetworkServer().GetServerList(ctx, &interModel.GetNetworkServerListInput{
		PaginationInput: interModel.PaginationInput{
			PageNum:  1,
			PageSize: consts.ServerListLimit,
		},
	})
	if networkServerListResErr != nil {
		return networkServerListResErr
	}
	if total >= consts.ServerListLimit {
		return fmt.Errorf("server限制数量为%d实际已有%d超出限制，请联系管理员处理", consts.ServerListLimit, total)
	}
	var allServerModels = make([]model.Server, len(networkServerListRes))
	for index, node := range networkServerListRes {
		allServerModels[index] = MapperServer(ctx, *node)
	}

	for index := range allServerModels {
		go func(server model.Server) error {
			if server.Disabled {
				return nil
			}
			go func() {
				err := startServer(ctx, &server)
				if err != nil {
					g.Log().Errorf(ctx, "服务启动失败！服务名称：%s，无法启动。错误信息：%s", server.Name, server.Addr, err)
				}
			}()
			return nil
		}(allServerModels[index])
	}
	return nil
}

// LoadServer 启动时候加载服务
func LoadServer(ctx context.Context, id int) error {
	networkServerRes, networkServerResErr := service.NetworkServer().GetServerById(ctx, id)
	if networkServerResErr != nil {
		return networkServerResErr
	}
	server := MapperServer(ctx, *networkServerRes)
	if server.Disabled {
		return nil
	}
	err := startServer(ctx, &server)
	if err != nil {
		return err
	}
	return nil
}

// GetServer 获取通道
func GetServer(id int) *Server {
	d, ok := allServers.Load(id)
	if ok {
		return d.(*Server)
	}
	return nil
}

func RemoveServer(id int) error {
	d, ok := allServers.LoadAndDelete(id)
	if ok {
		tnl := d.(*Server)
		return tnl.Instance.Close()
	}
	return nil //error
}
