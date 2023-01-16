package alarm

import (
	"context"
	"encoding/json"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/logic/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAlarmRule struct{}

func init() {
	service.RegisterAlarmRule(alarmRuleNew())
}

func alarmRuleNew() *sAlarmRule {
	return &sAlarmRule{}
}

func (s *sAlarmRule) List(ctx context.Context, in *model.AlarmRuleListInput) (out *model.AlarmRuleListOutput, err error) {
	out = new(model.AlarmRuleListOutput)
	c := dao.AlarmRule.Columns()
	m := dao.AlarmRule.Ctx(ctx).WithAll().OrderDesc(c.Id)

	if len(in.DateRange) > 0 {
		m = m.WhereBetween(c.CreatedAt, in.DateRange[0], in.DateRange[1])
	}

	out.Total, _ = m.Count()
	out.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Scan(&out.List)
	if err != nil {
		return
	}
	for i, v := range out.List {
		out.List[i].TriggerTypeName = model.AlarmTriggerType[v.TriggerType]
	}

	return
}

func (s *sAlarmRule) Cache(ctx context.Context) (rs map[string][]model.AlarmRuleOutput, err error) {
	key := "alarm:rule"
	tag := "alarm"
	value := common.Cache().GetOrSetFunc(ctx, key, func(ctx context.Context) (value interface{}, err error) {
		var list []model.AlarmRuleOutput
		err = dao.AlarmRule.Ctx(ctx).WithAll().
			Where(dao.AlarmRule.Columns().Status, 1).
			OrderDesc(dao.AlarmRule.Columns().Id).
			Scan(&list)
		if err != nil || len(list) == 0 {
			return
		}

		rs := make(map[string][]model.AlarmRuleOutput)
		for _, v := range list {
			if v.TriggerCondition != "" {
				err = json.Unmarshal([]byte(v.TriggerCondition), &v.Condition)
			}
			if v.Action != "" {
				err = json.Unmarshal([]byte(v.Action), &v.PerformAction)
			}
			rs[v.ProductKey] = append(rs[v.ProductKey], v)
		}
		value = rs
		return
	}, 0, tag)

	data := gconv.Map(value)
	rs = make(map[string][]model.AlarmRuleOutput, len(data))
	for k, v := range data {
		var t []model.AlarmRuleOutput
		if err = gconv.Scan(v, &t); err == nil {
			rs[k] = t
		}
	}
	return
}
func (s *sAlarmRule) delCache(ctx context.Context) {
	key := "alarm:rule"
	common.Cache().Remove(ctx, key)
}

func (s *sAlarmRule) Detail(ctx context.Context, id uint64) (out *model.AlarmRuleOutput, err error) {
	err = dao.AlarmRule.Ctx(ctx).WithAll().Where(dao.AlarmRule.Columns().Id, id).Scan(&out)
	if err != nil || out == nil {
		return
	}
	out.TriggerTypeName = model.AlarmTriggerType[out.TriggerType]

	if out.TriggerCondition != "" {
		err = json.Unmarshal([]byte(out.TriggerCondition), &out.Condition)
	}
	if out.Action != "" {
		err = json.Unmarshal([]byte(out.Action), &out.PerformAction)
	}
	return
}

func (s *sAlarmRule) Add(ctx context.Context, in *model.AlarmRuleAddInput) (err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.AlarmRule
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.CreateBy = uint(loginUserId)
	param.Status = 0
	param.TriggerCondition, _ = json.Marshal(in.AlarmTriggerCondition)
	param.Action, _ = json.Marshal(in.AlarmPerformAction)

	_, err = dao.AlarmRule.Ctx(ctx).Data(param).Insert()
	if err != nil {
		return
	}
	s.delCache(ctx)

	return
}

func (s *sAlarmRule) Edit(ctx context.Context, in *model.AlarmRuleEditInput) (err error) {
	p, err := s.Detail(ctx, in.Id)
	if err != nil {
		return err
	}
	if p == nil {
		err = gerror.New("告警规则不存在")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	var param *do.AlarmRule
	err = gconv.Scan(in, &param)
	if err != nil {
		return
	}
	param.UpdateBy = uint(loginUserId)
	param.Id = nil
	param.TriggerCondition, _ = json.Marshal(in.AlarmTriggerCondition)
	param.Action, _ = json.Marshal(in.AlarmPerformAction)

	_, err = dao.AlarmRule.Ctx(ctx).Data(param).Where(dao.AlarmRule.Columns().Id, in.Id).Update()
	if err != nil {
		return
	}
	s.delCache(ctx)

	return
}

func (s *sAlarmRule) Deploy(ctx context.Context, id uint64) (err error) {
	var p *entity.AlarmRule
	err = dao.AlarmRule.Ctx(ctx).Where(dao.AlarmRule.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}
	if p == nil || p.Status == model.AlarmRuleStatusOn {
		err = gerror.New("告警规则不存在，或已启用")
		return
	}

	_, err = dao.AlarmRule.Ctx(ctx).
		Data(g.Map{dao.AlarmRule.Columns().Status: model.AlarmRuleStatusOn}).
		Where(dao.AlarmRule.Columns().Id, id).
		Update()
	if err != nil {
		return
	}
	s.delCache(ctx)

	return
}

func (s *sAlarmRule) Undeploy(ctx context.Context, id uint64) (err error) {
	var p *entity.AlarmRule
	err = dao.AlarmRule.Ctx(ctx).Where(dao.AlarmRule.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}
	if p == nil || p.Status == model.AlarmRuleStatusOff {
		err = gerror.New("告警规则不存在，或已禁用")
		return
	}

	_, err = dao.AlarmRule.Ctx(ctx).
		Data(g.Map{dao.AlarmRule.Columns().Status: model.AlarmRuleStatusOff}).
		Where(dao.AlarmRule.Columns().Id, id).
		Update()
	if err != nil {
		return
	}
	s.delCache(ctx)

	return
}

func (s *sAlarmRule) Del(ctx context.Context, id uint64) (err error) {
	var p *entity.AlarmRule
	err = dao.AlarmRule.Ctx(ctx).Where(dao.AlarmRule.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		err = gerror.New("告警规则不存在")
		return
	}
	if p.Status == model.AlarmRuleStatusOn {
		err = gerror.New("告警规则已启用，请先禁用，再删除")
		return
	}

	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	_, err = dao.AlarmRule.Ctx(ctx).
		Data(do.AlarmRule{
			DeletedBy: uint(loginUserId),
			DeletedAt: gtime.Now(),
		}).
		Where(dao.AlarmRule.Columns().Id, id).
		Where(dao.AlarmRule.Columns().Status, model.AlarmRuleStatusOff).
		Unscoped().
		Update()
	if err != nil {
		return
	}
	s.delCache(ctx)

	return
}

func (s *sAlarmRule) Operator(ctx context.Context) (out []model.OperatorOutput, err error) {
	out = []model.OperatorOutput{
		{Title: "等于", Type: model.OperatorEq},
		{Title: "不等于", Type: model.OperatorNe},
		{Title: "大于", Type: model.OperatorGt},
		{Title: "大于等于", Type: model.OperatorGte},
		{Title: "小于", Type: model.OperatorLt},
		{Title: "小于等于", Type: model.OperatorLte},
		{Title: "在...之间", Type: model.OperatorBet},
		{Title: "不在...之间", Type: model.OperatorNbet},
	}
	return
}

func (s *sAlarmRule) TriggerType(ctx context.Context, productKey string) (out []model.TriggerTypeOutput, err error) {
	out = []model.TriggerTypeOutput{
		{Title: model.AlarmTriggerType[model.AlarmTriggerTypeOnline], Type: model.AlarmTriggerTypeOnline},
		{Title: model.AlarmTriggerType[model.AlarmTriggerTypeOffline], Type: model.AlarmTriggerTypeOffline},
	}

	product, err := service.DevProduct().Get(ctx, productKey)
	if err != nil || product == nil {
		return
	}
	if product.TSL != nil {
		if len(product.TSL.Properties) > 0 {
			out = append(out, model.TriggerTypeOutput{
				Title: model.AlarmTriggerType[model.AlarmTriggerTypeProperty], Type: model.AlarmTriggerTypeProperty,
			})
		}
		if len(product.TSL.Events) > 0 {
			out = append(out, model.TriggerTypeOutput{
				Title: model.AlarmTriggerType[model.AlarmTriggerTypeEvent], Type: model.AlarmTriggerTypeEvent,
			})
		}
	}

	return
}

func (s *sAlarmRule) TriggerParam(ctx context.Context, productKey string, triggerType int) (out []model.TriggerParamOutput, err error) {
	out = []model.TriggerParamOutput{

		{Title: "上报时间", ParamKey: "sysReportTime"},
	}

	product, err := service.DevProduct().Get(ctx, productKey)
	if err != nil || product == nil {
		return
	}
	if product.TSL != nil {
		switch {
		case triggerType == model.AlarmTriggerTypeProperty && len(product.TSL.Properties) > 0:
			for _, v := range product.TSL.Properties {
				out = append(out, model.TriggerParamOutput{
					Title: v.Name, ParamKey: v.Key,
				})
			}
		case triggerType == model.AlarmTriggerTypeEvent && len(product.TSL.Events) > 0:
			for _, v := range product.TSL.Events {
				out = append(out, model.TriggerParamOutput{
					Title: v.Name, ParamKey: v.Key,
				})
			}
		}
	}

	return
}
