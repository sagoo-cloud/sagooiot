package alarm

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AlarmLogDetailReq struct {
	g.Meta `path:"/log/detail" method:"get" summary:"告警日志详情" tags:"告警"`
	Id     uint64 `json:"id" dc:"告警日志ID" v:"required#告警日志ID不能为空"`
}
type AlarmLogDetailRes struct {
	Data *model.AlarmLogOutput `json:"data" dc:"告警日志详情"`
}

type AlarmLogListReq struct {
	g.Meta `path:"/log/list" method:"get" summary:"告警日志" tags:"告警"`
	Status string `p:"status"` //告警状态
	common.PaginationReq
}
type AlarmLogListRes struct {
	*model.AlarmLogListOutput
}

type AlarmLogHandleReq struct {
	g.Meta  `path:"/log/handle" method:"post" summary:"告警处理" tags:"告警"`
	Id      uint64 `json:"id" dc:"告警日志ID" v:"required#告警日志ID不能为空"`
	Status  int    `json:"status" d:"1" dc:"处理状态" v:"required|in:1,2#请选择处理状态|未知的处理状态，请正确选择"`
	Content string `json:"content" dc:"处理意见"`
}
type AlarmLogHandleRes struct{}
