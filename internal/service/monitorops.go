// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/sagoo-cloud/sagooiot/internal/model"
)

type (
	IMonitoropsRemoteconf interface {
		// GetRemoteconfList 获取列表数据
		GetRemoteconfList(ctx context.Context, in *model.GetRemoteconfListInput) (list []*model.RemoteconfOutput, err error)
		// GetRemoteconfById 获取指定ID数据
		GetRemoteconfById(ctx context.Context, id int) (out *model.RemoteconfOutput, err error)
		// AddRemoteconf 添加数据
		AddRemoteconf(ctx context.Context, in model.RemoteconfAddInput) (err error)
		// EditRemoteconf 修改数据
		EditRemoteconf(ctx context.Context, in model.RemoteconfEditInput) (err error)
		// DeleteRemoteconf 删除数据
		DeleteRemoteconf(ctx context.Context, Ids []int) (err error)
	}
)

var (
	localMonitoropsRemoteconf IMonitoropsRemoteconf
)

func MonitoropsRemoteconf() IMonitoropsRemoteconf {
	if localMonitoropsRemoteconf == nil {
		panic("implement not found for interface IMonitoropsRemoteconf, forgot register?")
	}
	return localMonitoropsRemoteconf
}

func RegisterMonitoropsRemoteconf(i IMonitoropsRemoteconf) {
	localMonitoropsRemoteconf = i
}
