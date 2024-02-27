package product

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/product"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var Device = cDevice{}

type cDevice struct{}

func (c *cDevice) Detail(ctx context.Context, req *product.DetailDeviceReq) (res *product.DetailDeviceRes, err error) {
	p, err := service.DevDevice().Detail(ctx, req.DeviceKey)
	res = &product.DetailDeviceRes{
		Data: p,
	}
	return
}

func (c *cDevice) ListForPage(ctx context.Context, req *product.ListDeviceForPageReq) (res *product.ListDeviceForPageRes, err error) {
	out, err := service.DevDevice().ListForPage(ctx, req.ListDeviceForPageInput)
	res = &product.ListDeviceForPageRes{
		ListDeviceForPageOutput: out,
	}
	return
}

func (c *cDevice) List(ctx context.Context, req *product.ListDeviceReq) (res *product.ListDeviceRes, err error) {
	out, err := service.DevDevice().List(ctx, req.ProductKey, req.KeyWord)
	res = &product.ListDeviceRes{
		Device: out,
	}
	return
}

func (c *cDevice) Add(ctx context.Context, req *product.AddDeviceReq) (res *product.AddDeviceRes, err error) {
	_, err = service.DevDevice().Add(ctx, req.AddDeviceInput)
	return
}

func (c *cDevice) Edit(ctx context.Context, req *product.EditDeviceReq) (res *product.EditDeviceRes, err error) {
	err = service.DevDevice().Edit(ctx, req.EditDeviceInput)
	return
}

func (c *cDevice) UpdateExtend(ctx context.Context, req *product.UpdateDeviceExtendReq) (res *product.UpdateDeviceExtendRes, err error) {
	err = service.DevDevice().UpdateExtend(ctx, req.DeviceExtendInput)
	return
}

func (c *cDevice) Del(ctx context.Context, req *product.DelDeviceReq) (res *product.DelDeviceRes, err error) {
	err = service.DevDevice().Del(ctx, req.Keys)
	return
}

func (c *cDevice) Deploy(ctx context.Context, req *product.DeployDeviceReq) (res *product.DeployDeviceRes, err error) {
	err = service.DevDevice().Deploy(ctx, req.DeviceKey)
	return
}

func (c *cDevice) Undeploy(ctx context.Context, req *product.UndeployDeviceReq) (res *product.UndeployDeviceRes, err error) {
	err = service.DevDevice().Undeploy(ctx, req.DeviceKey)
	return
}

func (c *cDevice) RunStatus(ctx context.Context, req *product.DeviceRunStatusReq) (res *product.DeviceRunStatusRes, err error) {
	out, err := service.DevDevice().RunStatus(ctx, req.DeviceKey)
	if err != nil {
		return
	}
	res = &product.DeviceRunStatusRes{
		DeviceRunStatusOutput: out,
	}
	return
}

func (c *cDevice) GetLatestProperty(ctx context.Context, req *product.DeviceGetLatestPropertyReq) (res *product.DeviceGetLatestPropertyRes, err error) {
	list, err := service.DevDevice().GetLatestProperty(ctx, req.DeviceKey)
	if err != nil {
		return
	}
	res = &product.DeviceGetLatestPropertyRes{
		List: list,
	}
	return
}

func (c *cDevice) GetProperty(ctx context.Context, req *product.DeviceGetPropertyReq) (res *product.DeviceGetPropertyRes, err error) {
	out, err := service.DevDevice().GetProperty(ctx, req.DeviceGetPropertyInput)
	res = &product.DeviceGetPropertyRes{
		DevicePropertiy: out,
	}
	return
}

func (c *cDevice) GetPropertyList(ctx context.Context, req *product.DeviceGetPropertyListReq) (res *product.DeviceGetPropertyListRes, err error) {
	var input *model.DeviceGetPropertyListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}
	out, err := service.DevDevice().GetPropertyList(ctx, input)
	res = &product.DeviceGetPropertyListRes{
		DeviceGetPropertyListOutput: out,
	}
	return
}

func (c *cDevice) BindSubDevice(ctx context.Context, req *product.DeviceBindReq) (res *product.DeviceBindRes, err error) {
	err = service.DevDevice().BindSubDevice(ctx, req.DeviceBindInput)
	return
}

func (c *cDevice) UnBindSubDevice(ctx context.Context, req *product.DeviceUnBindReq) (res *product.DeviceUnBindRes, err error) {
	err = service.DevDevice().UnBindSubDevice(ctx, req.DeviceBindInput)
	return
}

func (c *cDevice) BindList(ctx context.Context, req *product.BindListReq) (res *product.BindListRes, err error) {
	out, err := service.DevDevice().BindList(ctx, req.DeviceBindListInput)
	res = &product.BindListRes{
		DeviceBindListOutput: out,
	}
	return
}

func (c *cDevice) ListForSub(ctx context.Context, req *product.ListForSubReq) (res *product.ListForSubRes, err error) {
	out, err := service.DevDevice().ListForSub(ctx, req.ListForSubInput)
	res = &product.ListForSubRes{
		ListDeviceForPageOutput: out,
	}
	return
}

func (c *cDevice) DelSub(ctx context.Context, req *product.DelSubDeviceReq) (res *product.DelSubDeviceRes, err error) {
	err = service.DevDevice().DelSub(ctx, req.DeviceKey)
	return
}

func (c *cDevice) ImportDevices(ctx context.Context, req *product.ImportDevicesReq) (res product.ImportDevicesRes, err error) {
	return service.DevDevice().ImportDevices(ctx, req)
}

func (c *cDevice) ExportDevices(ctx context.Context, req *product.ExportDevicesReq) (res product.ExportDevicesRes, err error) {
	return service.DevDevice().ExportDevices(ctx, req)
}
func (c *cDevice) SetDevicesStatus(ctx context.Context, req *product.SetDeviceStatusReq) (res product.SetDeviceStatusRes, err error) {
	return service.DevDevice().SetDevicesStatus(ctx, req)
}

func (c *cDevice) GetDeviceDataList(ctx context.Context, req *product.DeviceDataListReq) (res *product.DeviceDataListRes, err error) {
	out, err := service.DevDevice().GetDeviceDataList(ctx, req.DeviceDataListInput)
	res = &product.DeviceDataListRes{
		DeviceDataListOutput: out,
	}
	return
}
