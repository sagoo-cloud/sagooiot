package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/product"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var Device = cDevice{}

type cDevice struct{}

func (c *cDevice) Get(ctx context.Context, req *product.GetDeviceReq) (res *product.GetDeviceRes, err error) {
	p, err := service.DevDevice().Get(ctx, req.Key)
	res = &product.GetDeviceRes{
		Data: p,
	}
	return
}

func (c *cDevice) Detail(ctx context.Context, req *product.DetailDeviceReq) (res *product.DetailDeviceRes, err error) {
	p, err := service.DevDevice().Detail(ctx, req.Id)
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
	out, err := service.DevDevice().List(ctx, req.ListDeviceInput)
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

func (c *cDevice) Del(ctx context.Context, req *product.DelDeviceReq) (res *product.DelDeviceRes, err error) {
	err = service.DevDevice().Del(ctx, req.Ids)
	return
}

func (c *cDevice) Deploy(ctx context.Context, req *product.DeployDeviceReq) (res *product.DeployDeviceRes, err error) {
	err = service.DevDevice().Deploy(ctx, req.Id)
	return
}

func (c *cDevice) Undeploy(ctx context.Context, req *product.UndeployDeviceReq) (res *product.UndeployDeviceRes, err error) {
	err = service.DevDevice().Undeploy(ctx, req.Id)
	return
}

func (c *cDevice) Online(ctx context.Context, req *product.OnlineDeviceReq) (res *product.OnlineDeviceRes, err error) {
	err = service.DevDevice().Online(ctx, req.Key)
	return
}

func (c *cDevice) Offline(ctx context.Context, req *product.OfflineDeviceReq) (res *product.OfflineDeviceRes, err error) {
	err = service.DevDevice().Offline(ctx, req.Key)
	return
}

func (c *cDevice) RunStatus(ctx context.Context, req *product.DeviceRunStatusReq) (res *product.DeviceRunStatusRes, err error) {
	out, err := service.DevDevice().RunStatus(ctx, req.Id)
	if err != nil {
		return
	}
	res = &product.DeviceRunStatusRes{
		DeviceRunStatusOutput: out,
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
	out, err := service.DevDevice().GetPropertyList(ctx, req.DeviceGetPropertyListInput)
	res = &product.DeviceGetPropertyListRes{
		DeviceGetPropertyListOutput: out,
	}
	return
}

func (c *cDevice) Statistics(ctx context.Context, req *product.DeviceStatisticsReq) (res product.DeviceStatisticsRes, err error) {
	res.DeviceTotal, err = service.DevDevice().Total(ctx)
	return
}

func (c *cDevice) StatisticsForMonths(ctx context.Context, req *product.DeviceStatisticsForMonthsReq) (res product.DeviceStatisticsForMonthsRes, err error) {
	res.MsgTotal, err = service.DevDevice().TotalForMonths(ctx)
	if err != nil {
		return
	}
	res.AlarmTotal, err = service.DevDevice().AlarmTotalForMonths(ctx)
	return
}
