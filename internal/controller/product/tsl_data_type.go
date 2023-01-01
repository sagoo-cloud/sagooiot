package product

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/product"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var TSLDataType = cTSLDataType{}

type cTSLDataType struct{}

func (c *cTSLDataType) DataTypeValueList(ctx context.Context, req *product.DateTypeReq) (res *product.DateTypeRes, err error) {
	res = new(product.DateTypeRes)
	res.DataType, err = service.DevTSLDataType().DataTypeValueList(ctx)
	return
}
