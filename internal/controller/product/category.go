package product

import (
	"context"
	"sagooiot/api/v1/product"
	"sagooiot/internal/service"
)

var Category = cCategory{}

type cCategory struct{}

func (c *cCategory) ListForPage(ctx context.Context, req *product.CategoryListForPageReq) (res *product.CategoryListForPageRes, err error) {
	list, total, err := service.DevCategory().ListForPage(ctx, req.PageNum, req.PageSize, req.Name)
	res = &product.CategoryListForPageRes{
		Category: list,
	}
	res.Total = total
	res.CurrentPage = req.PageNum
	return
}

func (c *cCategory) List(ctx context.Context, req *product.CategoryListReq) (res *product.CategoryListRes, err error) {
	list, err := service.DevCategory().List(ctx, req.Name)
	res = &product.CategoryListRes{
		Category: list,
	}
	return
}

func (c *cCategory) Add(ctx context.Context, req *product.AddCategoryReq) (res *product.AddCategoryRes, err error) {
	err = service.DevCategory().Add(ctx, req.AddProductCategoryInput)
	return
}

func (c *cCategory) Edit(ctx context.Context, req *product.EditCategoryReq) (res *product.EditCategoryRes, err error) {
	err = service.DevCategory().Edit(ctx, req.EditProductCategoryInput)
	return
}

func (c *cCategory) Del(ctx context.Context, req *product.DelCategoryReq) (res *product.DelCategoryRes, err error) {
	err = service.DevCategory().Del(ctx, req.Id)
	return
}
