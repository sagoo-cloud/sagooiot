package source

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DataNodeAddReq struct {
	g.Meta `path:"/node/add" method:"post" summary:"添加数据节点" tags:"数据源"`
	*model.DataNodeAddInput
}
type DataNodeAddRes struct{}

type DataNodeEditReq struct {
	g.Meta `path:"/node/edit" method:"put" summary:"编辑数据节点" tags:"数据源"`
	*model.DataNodeEditInput
}
type DataNodeEditRes struct{}

type DataNodeDelReq struct {
	g.Meta `path:"/node/del" method:"delete" summary:"删除数据节点" tags:"数据源"`
	NodeId uint64 `json:"nodeId" dc:"数据节点ID" v:"required#数据节点ID不能为空"`
}
type DataNodeDelRes struct{}

type DataNodeListReq struct {
	g.Meta   `path:"/node/list" method:"get" summary:"数据节点列表" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataNodeListRes struct {
	List []*model.DataNodeOutput `json:"list" dc:"数据节点列表"`
}
