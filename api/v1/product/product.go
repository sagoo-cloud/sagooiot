package product

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type DetailProductReq struct {
	g.Meta     `path:"/detail" method:"get" summary:"产品详情" tags:"产品"`
	ProductKey string `json:"productKey" dc:"产品标识" v:"required#产品标识不能为空"`
}
type DetailProductRes struct {
	Data *model.DetailProductOutput `json:"data" dc:"产品详情"`
}

type ListForPageReq struct {
	g.Meta `path:"/page_list" method:"get" summary:"产品搜索列表（分页）" tags:"产品"`
	model.ListForPageInput
}
type ListForPageRes struct {
	model.ListForPageOutput
}

type ListReq struct {
	g.Meta `path:"/list" method:"get" summary:"产品列表" tags:"产品"`
}
type ListRes struct {
	Product []*model.ProductOutput `json:"product" dc:"产品列表"`
}

type AddProductReq struct {
	g.Meta `path:"/add" method:"post" summary:"添加产品" tags:"产品"`
	*model.AddProductInput
}
type AddProductRes struct{}

type EditProductReq struct {
	g.Meta `path:"/edit" method:"put" summary:"编辑产品" tags:"产品"`
	*model.EditProductInput
}
type EditProductRes struct{}

type UpdateExtendReq struct {
	g.Meta `path:"/extend/update" method:"put" summary:"更新产品扩展信息" tags:"产品"`
	*model.ExtendInput
}
type UpdateExtendRes struct{}

type DelProductReq struct {
	g.Meta `path:"/del" method:"delete" summary:"删除产品" tags:"产品"`
	Keys   []string `json:"keys" dc:"产品Key组" v:"required#产品KEY不能为空"`
}
type DelProductRes struct{}

type DeployProductReq struct {
	g.Meta     `path:"/deploy" method:"post" summary:"发布产品" tags:"产品"`
	ProductKey string `json:"productKey" dc:"产品标识" v:"required#产品标识不能为空"`
}
type DeployProductRes struct{}

type UndeployProductReq struct {
	g.Meta     `path:"/undeploy" method:"post" summary:"停用产品" tags:"产品"`
	ProductKey string `json:"productKey" dc:"产品标识" v:"required#产品标识不能为空"`
}
type UndeployProductRes struct{}

type UploadIconReq struct {
	g.Meta `path:"/icon/upload" method:"post" mime:"multipart/form-data" summary:"图标上传" tags:"产品"`
	Icon   *ghttp.UploadFile `json:"icon" type:"file" dc:"选择上传图片"`
}
type UploadIconRes struct {
	IconPath string `json:"name" dc:"图标地址"`
}

type ListForSubProductReq struct {
	g.Meta `path:"/sub_list" method:"get" summary:"子设备类型产品列表" tags:"产品"`
}
type ListForSubProductRes struct {
	Product []*model.ProductOutput `json:"product" dc:"子设备类型产品列表"`
}

type UpdateScriptInfoReq struct {
	g.Meta `path:"/script/update" method:"put" summary:"脚本更新" tags:"产品"`
	*model.ScriptInfoInput
}
type UpdateScriptInfoRes struct{}

type ConnectIntroReq struct {
	g.Meta     `path:"/connect_intro" method:"get" summary:"获取设备接入信息" tags:"产品"`
	ProductKey string `json:"productKey" dc:"产品标识" v:"required#产品标识不能为空"`
}
type ConnectIntroRes struct {
	Data *model.DeviceConnectIntroOutput `json:"data" dc:"设备接入信息"`
}
