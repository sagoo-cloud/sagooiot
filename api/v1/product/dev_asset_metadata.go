package product

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
)

// GetDevAssetMetadataListReq 获取数据列表
type GetDevAssetMetadataListReq struct {
	g.Meta     `path:"/dev_asset_metadata/list" method:"get" summary:"获取档案属性列表" tags:"档案管理"`
	KeyWord    string `json:"keyWord" dc:"搜索关键字"`      //搜索关键字
	ProductKey string `json:"productKey" dc:"对应产品key"` // 产品key
	common.PaginationReq
}
type GetDevAssetMetadataListRes struct {
	Data []GetDevAssetMetadataByIdRes
	common.PaginationRes
}

// GetDevAssetMetadataByIdReq 获取指定ID的数据
type GetDevAssetMetadataByIdReq struct {
	g.Meta `path:"/dev_asset_metadata/get" method:"get" summary:"获取档案属性" tags:"档案管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetDevAssetMetadataByIdRes struct {
	ProductKey string `json:"productKey"          description:"产品标识"`
	Name       string `json:"name"          description:"字段名称"`
	Desc       string `json:"desc"          description:"字段描述"`
	Types      string `json:"types"          description:"字段类型"`
	CreatedAt  string `json:"createdAt"          description:"创建时间"`
	DeletedAt  string `json:"deletedAt"          description:"删除时间"`
	Id         string `json:"id"          description:""`
	Title      string `json:"title"          description:"字段标题"`
	FieldName  string `json:"fieldName"          description:"关联字段名称"`
	UpdatedAt  string `json:"updatedAt"          description:"更新时间"`
}

// GetDevAssetMetadataByIdReq 获取指定ID的数据
type GetDevAssetMetadataByProductKeyReq struct {
	g.Meta     `path:"/dev_asset_metadata/key" method:"get" summary:"获取档案属性" tags:"档案管理"`
	ProductKey string `json:"productKey"        description:"productKey" v:"required#productKey不能为空"`
}
type GetDevAssetMetadataByProductKeyRes struct {
	ProductKey string `json:"productKey"          description:"产品标识"`
	Name       string `json:"name"          description:"字段名称"`
	Desc       string `json:"desc"          description:"字段描述"`
	Types      string `json:"types"          description:"字段类型"`
	CreatedAt  string `json:"createdAt"          description:"创建时间"`
	DeletedAt  string `json:"deletedAt"          description:"删除时间"`
	Id         string `json:"id"          description:""`
	Title      string `json:"title"          description:"字段标题"`
	FieldName  string `json:"fieldName"          description:"关联字段名称"`
	UpdatedAt  string `json:"updatedAt"          description:"更新时间"`
}

// AddDevAssetMetadataReq 添加数据
type AddDevAssetMetadataReq struct {
	g.Meta `path:"/dev_asset_metadata/add" method:"post" summary:"添加档案属性" tags:"档案管理"`

	ProductKey string `json:"productKey"       v:"required#产品标识不能为空"      description:"产品标识"`
	Name       string `json:"name"      v:"required#字段名称不能为空"      description:"字段名称"`
	Desc       string `json:"desc"          description:"字段描述"`
	Types      string `json:"types"          description:"字段类型"`
	Title      string `json:"title"      v:"required#字段标题不能为空"      description:"字段标题"`
}
type MetaData struct {
	ProductKey string `json:"productKey"    description:"产品标识"`
	Name       string `json:"name"          description:"字段名称"`
	Desc       string `json:"desc"          description:"字段描述"`
	Types      string `json:"types"          description:"字段类型"`
	Title      string `json:"title"          description:"字段标题"`
}
type AddDevAssetMetadataRes struct{}

// EditDevAssetMetadataReq 编辑数据api
type EditDevAssetMetadataReq struct {
	g.Meta     `path:"/dev_asset_metadata/edit" method:"put" summary:"编辑档案属性" tags:"档案管理"`
	Title      string `json:"title"       v:"required#字段标题不能为空"     description:"字段标题"`
	Id         string `json:"id"          description:""`
	Name       string `json:"name"         v:"required#字段名称不能为空"      description:"字段名称"`
	Desc       string `json:"desc"          description:"字段描述"`
	Types      string `json:"types"          description:"字段类型"`
	ProductKey string `json:"productKey"    v:"required#产品标识不能为空"      description:"产品标识"`
}
type EditDevAssetMetadataRes struct{}

// DeleteDevAssetMetadataReq 删除数据
type DeleteDevAssetMetadataReq struct {
	g.Meta `path:"/dev_asset_metadata/delete" method:"delete" summary:"删除档案属性" tags:"档案管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteDevAssetMetadataRes struct{}
