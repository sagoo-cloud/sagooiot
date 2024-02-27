package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var TSLTag = cTSLTag{}

type cTSLTag struct{}

func (c *cTSLTag) ListTag(ctx context.Context, req *product.ListTSLTagReq) (res *product.ListTSLTagRes, err error) {
	out, err := service.DevTSLTag().ListTag(ctx, req.ListTSLTagInput)
	res = &product.ListTSLTagRes{
		ListTSLTagOutput: out,
	}
	return
}

func (c *cTSLTag) AddTag(ctx context.Context, req *product.AddTSLTagReq) (res *product.AddTSLTagRes, err error) {
	err = service.DevTSLTag().AddTag(ctx, req.TSLTagInput)
	return
}

func (c *cTSLTag) EditTag(ctx context.Context, req *product.EditTSLTagReq) (res *product.EditTSLTagRes, err error) {
	err = service.DevTSLTag().EditTag(ctx, req.TSLTagInput)
	return
}

func (c *cTSLTag) DelTag(ctx context.Context, req *product.DelTSLTagReq) (res *product.DelTSLTagRes, err error) {
	err = service.DevTSLTag().DelTag(ctx, req.DelTSLTagInput)
	return
}
