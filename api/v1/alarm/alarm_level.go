package alarm

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AlarmLevelReq struct {
	g.Meta `path:"/level/all" method:"get" summary:"告警级别" tags:"告警"`
}
type AlarmLevelRes struct {
	*model.AlarmLevelListOutput
}

type AlarmLevelEditReq struct {
	g.Meta `path:"/level/edit" method:"put" summary:"告警配置" tags:"告警"`
	List   []*model.AlarmLevelEditInput `json:"list" dc:"告警级别列表" v:"required#告警级别不能为空"`
}
type AlarmLevelEditRes struct{}
