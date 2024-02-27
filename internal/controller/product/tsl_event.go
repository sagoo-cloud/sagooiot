package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var TSLEvent = cTSLEvent{}

type cTSLEvent struct{}

func (c *cTSLEvent) ListEvent(ctx context.Context, req *product.ListTSLEventReq) (res *product.ListTSLEventRes, err error) {
	out, err := service.DevTSLEvent().ListEvent(ctx, req.ListTSLEventInput)
	res = &product.ListTSLEventRes{
		ListTSLEventOutput: out,
	}
	return
}

func (c *cTSLFunction) AllEvent(ctx context.Context, req *product.AllTSLEventReq) (res *product.AllTSLEventRes, err error) {
	list, err := service.DevTSLEvent().AllEvent(ctx, req.ProductKey)
	res = &product.AllTSLEventRes{
		Data: list,
	}
	return
}

func (c *cTSLEvent) AddEvent(ctx context.Context, req *product.AddTSLEventReq) (res *product.AddTSLEventRes, err error) {
	err = service.DevTSLEvent().AddEvent(ctx, req.TSLEventAddInput)
	return
}

func (c *cTSLEvent) EditEvent(ctx context.Context, req *product.EditTSLEventReq) (res *product.EditTSLEventRes, err error) {
	err = service.DevTSLEvent().EditEvent(ctx, req.TSLEventAddInput)
	return
}

func (c *cTSLEvent) DelEvent(ctx context.Context, req *product.DelTSLEventReq) (res *product.DelTSLEventRes, err error) {
	err = service.DevTSLEvent().DelEvent(ctx, req.DelTSLEventInput)
	return
}
