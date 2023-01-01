package source

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DataTemplateAddReq struct {
	g.Meta `path:"/template/add" method:"post" summary:"添加数据模型" tags:"数据建模"`
	*model.DataTemplateAddInput
}
type DataTemplateAddRes struct{}

type DataTemplateEditReq struct {
	g.Meta `path:"/template/edit" method:"put" summary:"编辑数据模型" tags:"数据建模"`
	*model.DataTemplateEditInput
}
type DataTemplateEditRes struct{}

type DataTemplateDelReq struct {
	g.Meta `path:"/template/del" method:"delete" summary:"删除数据模型" tags:"数据建模"`
	Ids    []uint64 `json:"ids" dc:"数据模型Ids" v:"required#数据模型ID不能为空"`
}
type DataTemplateDelRes struct{}

type DataTemplateSearchReq struct {
	g.Meta `path:"/template/search" method:"get" summary:"搜索数据模型" tags:"数据建模"`
	*model.DataTemplateSearchInput
}
type DataTemplateSearchRes struct {
	*model.DataTemplateSearchOutput
}

type DataTemplateListReq struct {
	g.Meta `path:"/template/list" method:"get" summary:"已发布数据模型列表" tags:"数据建模"`
}
type DataTemplateListRes struct {
	List []*entity.DataTemplate `json:"list" dc:"数据模型列表"`
}

type DataTemplateReq struct {
	g.Meta `path:"/template/detail" method:"get" summary:"数据模型详情" tags:"数据建模"`
	Id     uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
}
type DataTemplateRes struct {
	Data *model.DataTemplateOutput `json:"data" dc:"数据模型详情"`
}

type DataTemplateDeployReq struct {
	g.Meta `path:"/template/deploy" method:"post" summary:"发布" tags:"数据建模"`
	Id     uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
}
type DataTemplateDeployRes struct{}

type DataTemplateUndeployReq struct {
	g.Meta `path:"/template/undeploy" method:"post" summary:"停用" tags:"数据建模"`
	Id     uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
}
type DataTemplateUndeployRes struct{}

type DataTemplateDataReq struct {
	g.Meta `path:"/template/getdata" method:"get" summary:"获取模型数据" tags:"数据建模"`
	*model.DataTemplateDataInput
}
type DataTemplateDataRes struct {
	*model.DataTemplateDataOutput
}

type DataTemplateCopyReq struct {
	g.Meta `path:"/template/copy" method:"post" summary:"复制模型" tags:"数据建模"`
	Id     uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
}
type DataTemplateCopyRes struct{}

type DataTemplateCheckRelationReq struct {
	g.Meta `path:"/template/relation_check" method:"get" summary:"检测数据模型是否需要设置关联" tags:"数据建模"`
	Id     uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
}
type DataTemplateCheckRelationRes struct {
	Yes bool `json:"yes" dc:"是否需要设置关联: true:需要设置, false:不需要"`
}

type DataTemplateRelationReq struct {
	g.Meta `path:"/template/relation" method:"post" summary:"设置主源、关联字段" tags:"数据建模"`
	*model.TemplateDataRelationInput
}
type DataTemplateRelationRes struct{}

type TemplateSourceListReq struct {
	g.Meta `path:"/template/source_list" method:"get" summary:"数据模型源列表" tags:"数据建模"`
	Id     uint64 `json:"id" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
}
type TemplateSourceListRes struct {
	List []*model.DataSourceOutput
}
