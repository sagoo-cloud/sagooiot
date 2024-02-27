package product

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/product"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var TSLFunction = cTSLFunction{}

type cTSLFunction struct{}

func (c *cTSLFunction) ListFunction(ctx context.Context, req *product.ListTSLFunctionReq) (res *product.ListTSLFunctionRes, err error) {
	var reqData = new(model.ListTSLFunctionInput)
	err = gconv.Scan(req, &reqData)
	if err != nil {
		return nil, err
	}
	out, err := service.DevTSLFunction().ListFunction(ctx, reqData)
	res = &product.ListTSLFunctionRes{
		ListTSLFunctionOutput: out,
	}
	return
}

func (c *cTSLFunction) AllFunction(ctx context.Context, req *product.AllTSLFunctionReq) (res *product.AllTSLFunctionRes, err error) {
	list, err := service.DevTSLFunction().AllFunction(ctx, req.ProductKey, req.InputsValueTypes)
	res = &product.AllTSLFunctionRes{
		Data: list,
	}
	return
}

func (c *cTSLFunction) AddFunction(ctx context.Context, req *product.AddTSLFunctionReq) (res *product.AddTSLFunctionRes, err error) {
	err = service.DevTSLFunction().AddFunction(ctx, req.TSLFunctionAddInput)
	return
}

func (c *cTSLFunction) EditFunction(ctx context.Context, req *product.EditTSLFunctionReq) (res *product.EditTSLFunctionRes, err error) {
	err = service.DevTSLFunction().EditFunction(ctx, req.TSLFunctionAddInput)
	return
}

func (c *cTSLFunction) DelFunction(ctx context.Context, req *product.DelTSLFunctionReq) (res *product.DelTSLFunctionRes, err error) {
	err = service.DevTSLFunction().DelFunction(ctx, req.DelTSLFunctionInput)
	return
}
