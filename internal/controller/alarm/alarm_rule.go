package alarm

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/api/v1/alarm"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

var AlarmRule = cAlarmRule{}

type cAlarmRule struct{}

func (c *cAlarmRule) List(ctx context.Context, req *alarm.AlarmRuleListReq) (res *alarm.AlarmRuleListRes, err error) {
	out, err := service.AlarmRule().List(ctx, req.AlarmRuleListInput)
	res = &alarm.AlarmRuleListRes{
		AlarmRuleListOutput: out,
	}
	return
}

func (c *cAlarmRule) Add(ctx context.Context, req *alarm.AlarmRuleAddReq) (res *alarm.AlarmRuleAddRes, err error) {
	err = service.AlarmRule().Add(ctx, req.AlarmRuleAddInput)
	return
}

func (c *cAlarmRule) Edit(ctx context.Context, req *alarm.AlarmRuleEditReq) (res *alarm.AlarmRuleEditRes, err error) {
	err = service.AlarmRule().Edit(ctx, req.AlarmRuleEditInput)
	return
}

func (c *cAlarmRule) Deploy(ctx context.Context, req *alarm.AlarmRuleDeployReq) (res *alarm.AlarmRuleDeployRes, err error) {
	err = service.AlarmRule().Deploy(ctx, req.Id)
	return
}

func (c *cAlarmRule) Undeploy(ctx context.Context, req *alarm.AlarmRuleUndeployReq) (res *alarm.AlarmRuleUndeployRes, err error) {
	err = service.AlarmRule().Undeploy(ctx, req.Id)
	return
}

func (c *cAlarmRule) Del(ctx context.Context, req *alarm.AlarmRuleDelReq) (res *alarm.AlarmRuleDelRes, err error) {
	err = service.AlarmRule().Del(ctx, req.Id)
	return
}

func (c *cAlarmRule) Detail(ctx context.Context, req *alarm.AlarmRuleDetailReq) (res *alarm.AlarmRuleDetailRes, err error) {
	out, err := service.AlarmRule().Detail(ctx, req.Id)
	if err != nil || out == nil {
		return
	}
	res = new(alarm.AlarmRuleDetailRes)
	res.Data = out
	return
}

func (c *cAlarmRule) Operator(ctx context.Context, req *alarm.AlarmRuleOperatorReq) (res *alarm.AlarmRuleOperatorRes, err error) {
	out, err := service.AlarmRule().Operator(ctx)
	if err != nil || out == nil {
		return
	}
	res = new(alarm.AlarmRuleOperatorRes)
	res.List = out
	return
}

func (c *cAlarmRule) TriggerType(ctx context.Context, req *alarm.AlarmRuleTriggerTypeReq) (res *alarm.AlarmRuleTriggerTypeRes, err error) {
	out, err := service.AlarmRule().TriggerType(ctx, req.ProductKey)
	if err != nil || out == nil {
		return
	}
	res = new(alarm.AlarmRuleTriggerTypeRes)
	res.List = out
	return
}

func (c *cAlarmRule) TriggerParam(ctx context.Context, req *alarm.AlarmRuleTriggerParamReq) (res *alarm.AlarmRuleTriggerParamRes, err error) {
	out, err := service.AlarmRule().TriggerParam(ctx, req.ProductKey)
	if err != nil || out == nil {
		return
	}
	res = new(alarm.AlarmRuleTriggerParamRes)
	res.List = out
	return
}
