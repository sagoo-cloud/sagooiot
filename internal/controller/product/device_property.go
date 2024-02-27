package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var DeviceProperty = cDeviceProperty{}

type cDeviceProperty struct{}

// Set 设备属性设置
func (c *cDeviceProperty) Set(ctx context.Context, req *product.DevicePropertyReq) (res *product.DevicePropertyRes, err error) {
	out, err := service.DevDeviceProperty().Set(ctx, req.DevicePropertyInput)
	res = &product.DevicePropertyRes{
		DevicePropertyOutput: out,
	}
	return
}
