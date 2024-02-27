package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var DeviceTag = cDeviceTag{}

type cDeviceTag struct{}

func (c *cDeviceTag) Add(ctx context.Context, req *product.AddTagDeviceReq) (res *product.AddTagDeviceRes, err error) {
	err = service.DevDeviceTag().Add(ctx, req.AddTagDeviceInput)
	return
}

func (c *cDeviceTag) Edit(ctx context.Context, req *product.EditTagDeviceReq) (res *product.EditTagDeviceRes, err error) {
	err = service.DevDeviceTag().Edit(ctx, req.EditTagDeviceInput)
	return
}

func (c *cDeviceTag) Del(ctx context.Context, req *product.DelTagDeviceReq) (res *product.DelTagDeviceRes, err error) {
	err = service.DevDeviceTag().Del(ctx, req.Id)
	return
}
