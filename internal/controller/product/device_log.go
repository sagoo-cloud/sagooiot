package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/product"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var DeviceLog = cDeviceLog{}

type cDeviceLog struct{}

// 日志类型
func (c *cDeviceLog) LogType(ctx context.Context, req *product.DeviceLogTypeReq) (res *product.DeviceLogTypeRes, err error) {
	res = new(product.DeviceLogTypeRes)
	res.List = service.DevDeviceLog().LogType(ctx)
	return
}

// 日志搜索
func (c *cDeviceLog) Search(ctx context.Context, req *product.DeviceLogSearchReq) (res *product.DeviceLogSearchRes, err error) {
	out, err := service.DevDeviceLog().Search(ctx, req.DeviceLogSearchInput)
	res = &product.DeviceLogSearchRes{
		DeviceLogSearchOutput: out,
	}
	return
}
