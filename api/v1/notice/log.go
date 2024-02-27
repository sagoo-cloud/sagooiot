package notice

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type LogSearchReq struct {
	g.Meta `path:"/log/search" method:"get" summary:"通知日志搜索" tags:"通知服务管理"`
	*model.NoticeLogSearchInput
}
type LogSearchRes struct {
	*model.NoticeLogSearchOutput
}

type LogDelReq struct {
	g.Meta `path:"/log/del" method:"delete" summary:"通知日志删除" tags:"通知服务管理"`
	Ids    []uint64 `json:"ids" dc:"日志Ids" v:"required#日志ID不能为空"`
}
type LogDelRes struct{}
