package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var DeviceTree = cDeviceTree{}

type cDeviceTree struct{}

func (c *cDeviceTree) List(ctx context.Context, req *product.DeviceTreeListReq) (res *product.DeviceTreeListRes, err error) {
	list, err := service.DevDeviceTree().List(ctx)
	if err != nil || len(list) == 0 {
		return
	}
	res = &product.DeviceTreeListRes{
		List: list,
	}
	return
}

func (c *cDeviceTree) Change(ctx context.Context, req *product.DeviceTreeChangeReq) (res *product.DeviceTreeChangeRes, err error) {
	err = service.DevDeviceTree().Change(ctx, req.InfoId, req.ParentInfoId)
	return
}

func (c *cDeviceTree) Detail(ctx context.Context, req *product.DetailDeviceTreeInfoReq) (res *product.DetailDeviceTreeInfoRes, err error) {
	out, err := service.DevDeviceTree().Detail(ctx, req.InfoId)
	if err != nil || out == nil {
		return
	}
	res = &product.DetailDeviceTreeInfoRes{
		Data: out,
	}
	return
}

func (c *cDeviceTree) Add(ctx context.Context, req *product.AddDeviceTreeInfoReq) (res *product.AddDeviceTreeInfoRes, err error) {
	err = service.DevDeviceTree().Add(ctx, req.AddDeviceTreeInfoInput)
	return
}

func (c *cDeviceTree) Edit(ctx context.Context, req *product.EditDeviceTreeInfoReq) (res *product.EditDeviceTreeInfoRes, err error) {
	err = service.DevDeviceTree().Edit(ctx, req.EditDeviceTreeInfoInput)
	return
}

func (c *cDeviceTree) Del(ctx context.Context, req *product.DelDeviceTreeInfoReq) (res *product.DelDeviceTreeInfoRes, err error) {
	err = service.DevDeviceTree().Del(ctx, req.Id)
	return
}
