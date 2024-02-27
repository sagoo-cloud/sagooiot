package baseLogic

import (
	"context"
	"sagooiot/network/model"
)

func Online(ctx context.Context, msg model.DeviceOnlineMessage) error {
	// todo 设备状态更改不再使用状态机处理更改为逻辑层超时时间控制
	return nil
}

func Offline(ctx context.Context, msg model.DeviceOfflineMessage) error {
	// todo 设备状态更改不再使用状态机处理更改为逻辑层超时时间控制
	return nil
}
