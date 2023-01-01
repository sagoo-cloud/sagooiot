package source

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DataSourceApiAddReq struct {
	g.Meta `path:"/api/add" method:"post" summary:"添加API数据源" tags:"数据源"`
	*model.DataSourceApiAddInput
}
type DataSourceApiAddRes struct{}

type DataSourceApiEditReq struct {
	g.Meta `path:"/api/edit" method:"put" summary:"编辑API数据源" tags:"数据源"`
	*model.DataSourceApiEditInput
}
type DataSourceApiEditRes struct{}

type DataSourceApiGetReq struct {
	g.Meta   `path:"/api/get" method:"get" summary:"获取API数据" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceApiGetRes struct {
	Data string `json:"data" dc:"api源数据"`
}

type DataSourceDelReq struct {
	g.Meta `path:"/del" method:"delete" summary:"删除数据源" tags:"数据源"`
	Ids    []uint64 `json:"ids" dc:"数据源Ids" v:"required#数据源ID不能为空"`
}
type DataSourceDelRes struct{}

type DataSourceSearchReq struct {
	g.Meta `path:"/search" method:"get" summary:"搜索数据源" tags:"数据源"`
	*model.DataSourceSearchInput
}
type DataSourceSearchRes struct {
	*model.DataSourceSearchOutput
}

type DataSourceListReq struct {
	g.Meta `path:"/list" method:"get" summary:"数据源列表" tags:"数据源"`
}
type DataSourceListRes struct {
	List []*entity.DataSource `json:"list" dc:"数据源列表"`
}

type DataSourceReq struct {
	g.Meta   `path:"/detail" method:"get" summary:"数据源详情" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceRes struct {
	Data *model.DataSourceOutput `json:"data" dc:"数据源详情"`
}

type DataSourceDeployReq struct {
	g.Meta   `path:"/deploy" method:"post" summary:"发布" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceDeployRes struct{}

type DataSourceUndeployReq struct {
	g.Meta   `path:"/undeploy" method:"post" summary:"停用" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceUndeployRes struct{}

type DataSourceDataReq struct {
	g.Meta `path:"/getdata" method:"get" summary:"获取源数据记录" tags:"数据源"`
	*model.DataSourceDataInput
}
type DataSourceDataRes struct {
	*model.DataSourceDataOutput
}

type DataSourceCopyReq struct {
	g.Meta   `path:"/copy" method:"post" summary:"复制数据源" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceCopyRes struct{}
