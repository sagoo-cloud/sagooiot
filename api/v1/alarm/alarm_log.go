package alarm

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

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
	*model.AlarmLogListInput
}
type AlarmLogListRes struct {
	*model.AlarmLogListOutput
}

type AlarmLogHandleReq struct {
	g.Meta `path:"/log/handle" method:"post" summary:"告警处理" tags:"告警"`
	model.AlarmLogHandleInput
}
type AlarmLogHandleRes struct{}
