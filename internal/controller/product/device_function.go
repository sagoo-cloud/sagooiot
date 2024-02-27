package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var DeviceFunction = cDeviceFunction{}

type cDeviceFunction struct{}

// Do 执行设备功能
func (c *cDeviceFunction) Do(ctx context.Context, req *product.DeviceFunctionReq) (res *product.DeviceFunctionRes, err error) {
	out, err := service.DevDeviceFunction().Do(ctx, req.DeviceFunctionInput)
	res = &product.DeviceFunctionRes{
		DeviceFunctionOutput: out,
	}
	return
}
