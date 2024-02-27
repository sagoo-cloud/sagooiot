package alarm

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type AlarmRuleListReq struct {
	g.Meta `path:"/rule/list" method:"get" summary:"告警规则列表" tags:"告警"`
	common.PaginationReq
}
type AlarmRuleListRes struct {
	*model.AlarmRuleListOutput
}

type AlarmRuleAddReq struct {
	g.Meta `path:"/rule/add" method:"post" summary:"新增告警规则" tags:"告警"`
	*model.AlarmRuleAddInput
}
type AlarmRuleAddRes struct{}

type AlarmRuleEditReq struct {
	g.Meta `path:"/rule/edit" method:"put" summary:"编辑告警规则" tags:"告警"`
	*model.AlarmRuleEditInput
}
type AlarmRuleEditRes struct{}

type AlarmRuleDeployReq struct {
	g.Meta `path:"/rule/deploy" method:"post" summary:"启用" tags:"告警"`
	Id     uint64 `json:"id" dc:"告警规则ID" v:"required#告警规则ID不能为空"`
}
type AlarmRuleDeployRes struct{}

type AlarmRuleUndeployReq struct {
	g.Meta `path:"/rule/undeploy" method:"post" summary:"禁用" tags:"告警"`
	Id     uint64 `json:"id" dc:"告警规则ID" v:"required#告警规则ID不能为空"`
}
type AlarmRuleUndeployRes struct{}

type AlarmRuleDelReq struct {
	g.Meta `path:"/rule/del" method:"delete" summary:"删除" tags:"告警"`
	Id     uint64 `json:"id" dc:"告警规则ID" v:"required#告警规则ID不能为空"`
}
type AlarmRuleDelRes struct{}

type AlarmRuleDetailReq struct {
	g.Meta `path:"/rule/detail" method:"get" summary:"详情" tags:"告警"`
	Id     uint64 `json:"id" dc:"告警规则ID" v:"required#告警规则ID不能为空"`
}
type AlarmRuleDetailRes struct {
	Data *model.AlarmRuleOutput `json:"data" dc:"告警详情"`
}

type AlarmRuleOperatorReq struct {
	g.Meta `path:"/rule/operator" method:"get" summary:"操作符" tags:"告警"`
}
type AlarmRuleOperatorRes struct {
	List []model.OperatorOutput `json:"list" dc:"操作符列表"`
}

type AlarmRuleTriggerTypeReq struct {
	g.Meta     `path:"/rule/trigger_type" method:"get" summary:"触发类型" tags:"告警"`
	ProductKey string `json:"productKey" dc:"产品标识"`
}
type AlarmRuleTriggerTypeRes struct {
	List []model.TriggerTypeOutput `json:"list" dc:"触发类型列表"`
}

type AlarmRuleTriggerParamReq struct {
	g.Meta      `path:"/rule/trigger_param" method:"get" summary:"触发条件参数" tags:"告警"`
	ProductKey  string `json:"productKey" dc:"产品标识"`
	TriggerType int    `json:"triggerType" dc:"触发类型:1=上线,2=离线,3=属性上报,4=事件上报"`
	EventKey    string `json:"eventKey" dc:"事件标识"`
}
type AlarmRuleTriggerParamRes struct {
	List []model.TriggerParamOutput `json:"list" dc:"触发条件参数列表"`
}

type AlarmCronRuleAddReq struct {
	g.Meta `path:"/rule/cron/add" method:"post" summary:"新增定时触发告警规则" tags:"告警"`
	*model.AlarmCronRuleAddInput
}
type AlarmCronRuleAddRes struct{}

type AlarmCronRuleEditReq struct {
	g.Meta `path:"/rule/cron/edit" method:"put" summary:"编辑定时触发告警规则" tags:"告警"`
	*model.AlarmCronRuleEditInput
}
type AlarmCronRuleEditRes struct{}
