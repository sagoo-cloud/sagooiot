package product

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/product"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var TSLProperty = cTSLProperty{}

type cTSLProperty struct{}

func (c *cTSLProperty) ListProperty(ctx context.Context, req *product.ListTSLPropertyReq) (res *product.ListTSLPropertyRes, err error) {
	var reqData = new(model.ListTSLPropertyInput)
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return nil, err
	}
	out, err := service.DevTSLProperty().ListProperty(ctx, reqData)
	res = &product.ListTSLPropertyRes{
		ListTSLPropertyOutput: out,
	}
	return
}

func (c *cTSLProperty) AllProperty(ctx context.Context, req *product.AllTSLPropertyReq) (res *product.AllTSLPropertyRes, err error) {
	list, err := service.DevTSLProperty().AllProperty(ctx, req.ProductKey)
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
	var reqData = new(model.DelTSLPropertyInput)
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return nil, err
	}
	err = service.DevTSLProperty().DelProperty(ctx, reqData)
	return
}
