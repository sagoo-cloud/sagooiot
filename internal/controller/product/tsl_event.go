package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/product"
	"github.com/sagoo-cloud/sagooiot/internal/service"
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

func (c *cTSLEvent) AddEvent(ctx context.Context, req *product.AddTSLEventReq) (res *product.AddTSLEventRes, err error) {
	err = service.DevTSLEvent().AddEvent(ctx, req.TSLEventInput)
	return
}

func (c *cTSLEvent) EditEvent(ctx context.Context, req *product.EditTSLEventReq) (res *product.EditTSLEventRes, err error) {
	err = service.DevTSLEvent().EditEvent(ctx, req.TSLEventInput)
	return
}

func (c *cTSLEvent) DelEvent(ctx context.Context, req *product.DelTSLEventReq) (res *product.DelTSLEventRes, err error) {
	err = service.DevTSLEvent().DelEvent(ctx, req.DelTSLEventInput)
	return
}
