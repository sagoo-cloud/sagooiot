package source

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DataSourceDbAddReq struct {
	g.Meta `path:"/db/add" method:"post" summary:"添加数据库数据源" tags:"数据源"`
	*model.DataSourceDbAddInput
}
type DataSourceDbAddRes struct{}

type DataSourceDbEditReq struct {
	g.Meta `path:"/db/edit" method:"put" summary:"编辑数据库数据源" tags:"数据源"`
	*model.DataSourceDbEditInput
}
type DataSourceDbEditRes struct{}

type DataSourceDbGetReq struct {
	g.Meta   `path:"/db/get" method:"get" summary:"获取数据库数据" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceDbGetRes struct {
	Data string `json:"data" dc:"数据库源数据"`
}

type DataSourceDbFieldsReq struct {
	g.Meta   `path:"/db/fields" method:"get" summary:"获取数据表字段" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceDbFieldsRes struct {
	Data g.MapStrAny `json:"data" dc:"数据表字段"`
}
