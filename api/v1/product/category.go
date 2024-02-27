package product

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CategoryListForPageReq struct {
	g.Meta `path:"/category/page_list" method:"get" summary:"产品分类列表(分页)" tags:"产品分类"`
	Name   string `json:"name" dc:"分类名称"`
	*common.PaginationReq
}
type CategoryListForPageRes struct {
	Category []*model.ProductCategoryTreeOutput `json:"category" dc:"产品分类列表"`
	common.PaginationRes
}

type CategoryListReq struct {
	g.Meta `path:"/category/list" method:"get" summary:"产品分类列表" tags:"产品分类"`
	Name   string `json:"name" dc:"分类名称"`
}
type CategoryListRes struct {
	Category []*model.ProductCategoryTreeOutput `json:"category" dc:"产品分类列表"`
}

type AddCategoryReq struct {
	g.Meta `path:"/category/add" method:"post" summary:"添加产品分类" tags:"产品分类"`
	*model.AddProductCategoryInput
}
type AddCategoryRes struct{}

type EditCategoryReq struct {
	g.Meta `path:"/category/edit" method:"put" summary:"编辑产品分类" tags:"产品分类"`
	*model.EditProductCategoryInput
}
type EditCategoryRes struct{}

type DelCategoryReq struct {
	g.Meta `path:"/category/del" method:"delete" summary:"删除产品分类" tags:"产品分类"`
	Id     uint `json:"id" dc:"分类ID" v:"required#分类ID不能为空"`
}
type DelCategoryRes struct{}
