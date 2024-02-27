package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var TSLDataType = cTSLDataType{}

type cTSLDataType struct{}

func (c *cTSLDataType) DataTypeValueList(ctx context.Context, req *product.DateTypeReq) (res *product.DateTypeRes, err error) {
	res = new(product.DateTypeRes)
	res.DataType, err = service.DevTSLDataType().DataTypeValueList(ctx)
	return
}
