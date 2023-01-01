package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/product"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var TSLProperty = cTSLProperty{}

type cTSLProperty struct{}

func (c *cTSLProperty) ListProperty(ctx context.Context, req *product.ListTSLPropertyReq) (res *product.ListTSLPropertyRes, err error) {
	out, err := service.DevTSLProperty().ListProperty(ctx, req.ListTSLPropertyInput)
	res = &product.ListTSLPropertyRes{
		ListTSLPropertyOutput: out,
	}
	return
}

func (c *cTSLProperty) AllProperty(ctx context.Context, req *product.AllTSLPropertyReq) (res *product.AllTSLPropertyRes, err error) {
	list, err := service.DevTSLProperty().AllProperty(ctx, req.Key)
	res = &product.AllTSLPropertyRes{
		Data: list,
	}
	return
}

func (c *cTSLProperty) AddProperty(ctx context.Context, req *product.AddTSLPropertyReq) (res *product.AddTSLPropertyRes, err error) {
	err = service.DevTSLProperty().AddProperty(ctx, req.TSLPropertyInput)
	return
}

func (c *cTSLProperty) EditProperty(ctx context.Context, req *product.EditTSLPropertyReq) (res *product.EditTSLPropertyRes, err error) {
	err = service.DevTSLProperty().EditProperty(ctx, req.TSLPropertyInput)
	return
}

func (c *cTSLProperty) DelProperty(ctx context.Context, req *product.DelTSLPropertyReq) (res *product.DelTSLPropertyRes, err error) {
	err = service.DevTSLProperty().DelProperty(ctx, req.DelTSLPropertyInput)
	return
}
