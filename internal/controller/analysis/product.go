package analysis

import (
	"context"
	"sagooiot/api/v1/analysis"
	"sagooiot/internal/service"
)

var Product = cProduct{}

type cProduct struct{}

// GetProductCount 获取产品数量统计
func (c *cProduct) GetProductCount(ctx context.Context, req *analysis.ProductCountReq) (res *analysis.ProductCountRes, err error) {
	data, err := service.AnalysisProduct().GetProductCount(ctx)
	if err != nil {
		return
	}
	res = &analysis.ProductCountRes{
		Total:   data.Total,
		Enable:  data.Enable,
		Disable: data.Disable,
		Added:   data.Added,
	}
	return
}

// GetDeviceCountForProduct 获取属于该产品下的设备数量
func (c *cProduct) GetDeviceCountForProduct(ctx context.Context, req *analysis.DeviceCountForProductReq) (res *analysis.DeviceCountForProductRes, err error) {
	data, err := service.AnalysisProduct().GetDeviceCountForProduct(ctx, req.ProductKey)
	if err != nil {
		return
	}
	res = &analysis.DeviceCountForProductRes{
		Number: data,
	}
	return
}
