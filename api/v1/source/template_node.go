package source

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DataTemplateNodeAddReq struct {
	g.Meta `path:"/template/node/add" method:"post" summary:"添加数据模型节点" tags:"数据建模"`
	*model.DataTemplateNodeAddInput
}
type DataTemplateNodeAddRes struct{}

type DataTemplateNodeEditReq struct {
	g.Meta `path:"/template/node/edit" method:"put" summary:"编辑数据模型节点" tags:"数据建模"`
	*model.DataTemplateNodeEditInput
}
type DataTemplateNodeEditRes struct{}

type DataTemplateNodeDelReq struct {
	g.Meta `path:"/template/node/del" method:"delete" summary:"删除数据模型节点" tags:"数据建模"`
	Id     uint64 `json:"id" dc:"数据模型节点ID" v:"required#数据模型节点ID不能为空"`
}
type DataTemplateNodeDelRes struct{}

type DataTemplateNodeListReq struct {
	g.Meta `path:"/template/node/list" method:"get" summary:"数据模型节点列表" tags:"数据建模"`
	Tid    uint64 `json:"tid" dc:"数据模型ID" v:"required#数据模型ID不能为空"`
}
type DataTemplateNodeListRes struct {
	List []*model.DataTemplateNodeOutput `json:"list" dc:"数据模型节点列表"`
}
