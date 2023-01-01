package source

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DataSourceDeviceAddReq struct {
	g.Meta `path:"/device/add" method:"post" summary:"添加设备数据源" tags:"数据源"`
	*model.DataSourceDeviceAddInput
}
type DataSourceDeviceAddRes struct{}

type DataSourceDeviceEditReq struct {
	g.Meta `path:"/device/edit" method:"put" summary:"编辑设备数据源" tags:"数据源"`
	*model.DataSourceDeviceEditInput
}
type DataSourceDeviceEditRes struct{}

type DataSourceDeviceGetReq struct {
	g.Meta   `path:"/device/get" method:"get" summary:"获取设备数据" tags:"数据源"`
	SourceId uint64 `json:"sourceId" dc:"数据源ID" v:"required#数据源ID不能为空"`
}
type DataSourceDeviceGetRes struct {
	Data string `json:"data" dc:"设备源数据"`
}
