package product

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DeviceLogTypeReq struct {
	g.Meta `path:"/log/type" method:"get" summary:"日志类型" tags:"时序数据库管理"`
}
type DeviceLogTypeRes struct {
	List []string `json:"list" dc:"日志类型列表"`
}

type DeviceLogSearchReq struct {
	g.Meta `path:"/log/search" method:"get" summary:"日志搜索" tags:"时序数据库管理"`
	*model.DeviceLogSearchInput
}
type DeviceLogSearchRes struct {
	*model.DeviceLogSearchOutput
}
