package alarm

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/database/gdb"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/dcache"

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

// TODO 废弃
func (s *sAlarmRule) Cache(ctx context.Context) (rs map[string][]model.AlarmRuleOutput, err error) {
	var list []model.AlarmRuleOutput
	err = dao.AlarmRule.Ctx(ctx).Cache(gdb.CacheOption{
		Duration: 0,
		Name:     consts.CacheAlarmRule,
		Force:    false,
	}).WithAll().
		Where(dao.AlarmRule.Columns().Status, 1).
		Where(dao.AlarmRule.Columns().TriggerMode, 1).
		OrderDesc(dao.AlarmRule.Columns().Id).
		Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}

	rs = make(map[string][]model.AlarmRuleOutput, len(list))
	for _, v := range list {
		if v.TriggerCondition != "" {
			conditionErr := json.Unmarshal([]byte(v.TriggerCondition), &v.Condition)
			if conditionErr != nil {
				return nil, err
			}
		}
		if v.Action != "" {
			performActionErr := json.Unmarshal([]byte(v.Action), &v.PerformAction)
			if err != nil {
				return nil, performActionErr
			}
		}
		rs[v.ProductKey] = append(rs[v.ProductKey], v)
	}

	data := gconv.Map(rs)
	rs = make(map[string][]model.AlarmRuleOutput, len(data))
	for k, v := range data {
		var t []model.AlarmRuleOutput
		if err = gconv.Scan(v, &t); err == nil {
			rs[k] = t
		}
	}
	return
}

// CacheAllAlarmRule 缓存所有的告警规则
func (s *sAlarmRule) CacheAllAlarmRule(ctx context.Context) (err error) {
	var list []model.AlarmRuleOutput
	err = dao.AlarmRule.Ctx(ctx).WithAll().
		Where(dao.AlarmRule.Columns().Status, 1).      // 启用
		Where(dao.AlarmRule.Columns().TriggerMode, 1). // 设备触发
		OrderDesc(dao.AlarmRule.Columns().Id).
		Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}
	productList := make(map[string]int)
	for _, v := range list {
		productList[v.ProductKey] = 1
	}

	rs := make(map[string][]model.AlarmRuleOutput, len(productList))
	for k := range productList {
		for _, d := range list {
			if d.ProductKey == k {
				rs[k] = append(rs[k], d)
			}
		}
		//将告警规则缓存到redis
		err := dcache.SetDeviceAlarmRule(ctx, k, rs[k])
		if err != nil {
			g.Log().Debug(ctx, "CacheAllAlarmRule Error：", err)
		}
	}
	return
}

// 缓存单个产品的告警规则
func (s *sAlarmRule) cacheProductAlarmRuleChange(ctx context.Context, productKey string) (err error) {
	var list []model.AlarmRuleOutput
	err = dao.AlarmRule.Ctx(ctx).WithAll().
		Where(dao.AlarmRule.Columns().ProductKey, productKey).
		Where(dao.AlarmRule.Columns().TriggerMode, 1). // 设备触发
		OrderDesc(dao.AlarmRule.Columns().Id).
		Scan(&list)
	if err != nil || len(list) == 0 {
		return
	}
	productList := make(map[string]int)
	for _, v := range list {
		productList[v.ProductKey] = 1
	}

	rs := make(map[string][]model.AlarmRuleOutput, len(productList))
	for k := range productList {
		for _, d := range list {
			if d.ProductKey == k {
				rs[k] = append(rs[k], d)
			}
		}
		//将告警规则缓存到redis
		err := dcache.SetDeviceAlarmRule(ctx, k, rs[k])
		if err != nil {
			g.Log().Debug(ctx, "CacheAllAlarmRule Error：", err)
		}
	}
	return
}

// Detail 获取告警规则详情
func (s *sAlarmRule) Detail(ctx context.Context, id uint64) (out *model.AlarmRuleOutput, err error) {
	err = dao.AlarmRule.Ctx(ctx).WithAll().Where(dao.AlarmRule.Columns().Id, id).Scan(&out)
	if err != nil || out == nil {
		return
	}
	out.TriggerTypeName = model.AlarmTriggerType[out.TriggerType]

	// 触发类型为上下线
	if out.TriggerType == consts.AlarmTriggerTypeOnline || out.TriggerType == consts.AlarmTriggerTypeOffline {
		out.TriggerCondition = ""
	}

	if out.TriggerCondition != "" {
		switch out.TriggerMode {
		case consts.AlarmTriggerModeDevice:
			err = json.Unmarshal([]byte(out.TriggerCondition), &out.Condition)
		case consts.AlarmTriggerModeCron:
			err = json.Unmarshal([]byte(out.TriggerCondition), &out.CronCondition)
		}
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

	// 触发类型为上下线
	if in.TriggerType == consts.AlarmTriggerTypeOnline || in.TriggerType == consts.AlarmTriggerTypeOffline {
		in.AlarmTriggerCondition = model.AlarmTriggerCondition{
			TriggerCondition: []model.AlarmCondition{
				{
					Filters: []model.AlarmFilters{
						{
							Key:      "sysReportTime",
							Operator: "gt",
							Value:    []string{"0"},
						},
					},
				},
			},
		}
	}
	if param.TriggerCondition, err = json.Marshal(in.AlarmTriggerCondition); err != nil {
		return
	}
	if param.Action, err = json.Marshal(in.AlarmPerformAction); err != nil {
		return
	}

	_, err = dao.AlarmRule.Ctx(ctx).Data(do.AlarmRule{
		DeptId:           service.Context().GetUserDeptId(ctx),
		Name:             param.Name,
		Level:            param.Level,
		ProductKey:       param.ProductKey,
		DeviceKey:        param.DeviceKey,
		TriggerType:      param.TriggerType,
		EventKey:         param.EventKey,
		TriggerCondition: param.TriggerCondition,
		Action:           param.Action,
		Status:           0,
		CreatedBy:        uint(loginUserId),
	}).Insert()
	if err != nil {
		return
	}

	//更新缓存
	err = s.cacheProductAlarmRuleChange(ctx, in.ProductKey)
	if err != nil {
		return err
	}
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
	param.UpdatedBy = uint(loginUserId)
	param.Id = nil

	// 触发类型为上下线
	if in.TriggerType == consts.AlarmTriggerTypeOnline || in.TriggerType == consts.AlarmTriggerTypeOffline {
		in.AlarmTriggerCondition = model.AlarmTriggerCondition{
			TriggerCondition: []model.AlarmCondition{
				{
					Filters: []model.AlarmFilters{
						{
							Key:      "sysReportTime",
							Operator: "gt",
							Value:    []string{"0"},
						},
					},
				},
			},
		}
	}
	if param.TriggerCondition, err = json.Marshal(in.AlarmTriggerCondition); err != nil {
		return
	}
	if param.Action, err = json.Marshal(in.AlarmPerformAction); err != nil {
		return
	}

	_, err = dao.AlarmRule.Ctx(ctx).Data(param).Where(dao.AlarmRule.Columns().Id, in.Id).Update()
	if err != nil {
		return
	}

	//更新缓存
	err = s.cacheProductAlarmRuleChange(ctx, p.ProductKey)
	if err != nil {
		return err
	}

	return
}

// Deploy 启用告警规则
func (s *sAlarmRule) Deploy(ctx context.Context, id uint64) (err error) {
	var p *entity.AlarmRule
	err = dao.AlarmRule.Ctx(ctx).Where(dao.AlarmRule.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}

	if p == nil || p.Status == consts.AlarmRuleStatusOn {
		err = gerror.New("告警规则不存在，或已启用")
		return
	}

	// 定时触发
	if p.TriggerMode == consts.AlarmTriggerModeCron {
		if err = s.start(ctx, id); err != nil {
			return
		}
	}

	_, err = dao.AlarmRule.Ctx(ctx).
		Data(g.Map{dao.AlarmRule.Columns().Status: consts.AlarmRuleStatusOn}).
		Where(dao.AlarmRule.Columns().Id, id).
		Update()
	if err != nil {
		return
	}
	//更新缓存
	err = s.cacheProductAlarmRuleChange(ctx, p.ProductKey)
	if err != nil {
		return err
	}

	return
}

func (s *sAlarmRule) Undeploy(ctx context.Context, id uint64) (err error) {
	var p *entity.AlarmRule
	err = dao.AlarmRule.Ctx(ctx).Where(dao.AlarmRule.Columns().Id, id).Scan(&p)
	if err != nil {
		return
	}
	if p == nil || p.Status == consts.AlarmRuleStatusOff {
		err = gerror.New("告警规则不存在，或已禁用")
		return
	}

	// 定时触发
	if p.TriggerMode == consts.AlarmTriggerModeCron {
		if err = s.stop(ctx, id); err != nil {
			return
		}
	}

	_, err = dao.AlarmRule.Ctx(ctx).
		Data(g.Map{dao.AlarmRule.Columns().Status: consts.AlarmRuleStatusOff}).
		Where(dao.AlarmRule.Columns().Id, id).
		Update()
	if err != nil {
		return
	}

	//更新缓存
	err = s.cacheProductAlarmRuleChange(ctx, p.ProductKey)
	if err != nil {
		return err
	}
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

	if p.Status == consts.AlarmRuleStatusOn {
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
		Where(dao.AlarmRule.Columns().Status, consts.AlarmRuleStatusOff).
		Unscoped().
		Update()
	if err != nil {
		return
	}

	//更新缓存
	err = s.cacheProductAlarmRuleChange(ctx, p.ProductKey)
	if err != nil {
		return err
	}

	return
}

func (s *sAlarmRule) Operator(ctx context.Context) (out []model.OperatorOutput, err error) {
	out = []model.OperatorOutput{
		{Title: "等于", Type: consts.OperatorEq},
		{Title: "不等于", Type: consts.OperatorNe},
		{Title: "大于", Type: consts.OperatorGt},
		{Title: "大于等于", Type: consts.OperatorGte},
		{Title: "小于", Type: consts.OperatorLt},
		{Title: "小于等于", Type: consts.OperatorLte},
		{Title: "在...之间", Type: consts.OperatorBet},
		{Title: "不在...之间", Type: consts.OperatorNbet},
	}
	return
}

func (s *sAlarmRule) TriggerType(ctx context.Context, productKey string) (out []model.TriggerTypeOutput, err error) {
	out = []model.TriggerTypeOutput{
		{Title: model.AlarmTriggerType[consts.AlarmTriggerTypeOnline], Type: consts.AlarmTriggerTypeOnline},
		{Title: model.AlarmTriggerType[consts.AlarmTriggerTypeOffline], Type: consts.AlarmTriggerTypeOffline},
	}

	product, err := dcache.GetProductDetailInfo(productKey)
	if err != nil || product == nil {
		return
	}
	if product.TSL != nil {
		if len(product.TSL.Properties) > 0 {
			out = append(out, model.TriggerTypeOutput{
				Title: model.AlarmTriggerType[consts.AlarmTriggerTypeProperty], Type: consts.AlarmTriggerTypeProperty,
			})
		}
		if len(product.TSL.Events) > 0 {
			out = append(out, model.TriggerTypeOutput{
				Title: model.AlarmTriggerType[consts.AlarmTriggerTypeEvent], Type: consts.AlarmTriggerTypeEvent,
			})
		}
	}

	return
}

func (s *sAlarmRule) TriggerParam(ctx context.Context, productKey string, triggerType int, eventKey ...string) (out []model.TriggerParamOutput, err error) {
	if triggerType == consts.AlarmTriggerTypeOnline || triggerType == consts.AlarmTriggerTypeOffline {
		return
	}

	out = []model.TriggerParamOutput{
		// {Title: "系统时间", ParamKey: "sysTime"},
		{Title: "上报时间", ParamKey: "sysReportTime"},
	}

	product, err := service.DevProduct().Detail(ctx, productKey)
	if err != nil || product == nil {
		return
	}
	if product.TSL != nil {
		switch {
		case triggerType == consts.AlarmTriggerTypeProperty && len(product.TSL.Properties) > 0:
			for _, v := range product.TSL.Properties {
				out = append(out, model.TriggerParamOutput{
					Title: v.Name, ParamKey: v.Key,
				})
			}
		case triggerType == consts.AlarmTriggerTypeEvent && len(product.TSL.Events) > 0:
			for _, v := range product.TSL.Events {
				if len(eventKey) > 0 && v.Key == eventKey[0] {
					for _, ot := range v.Outputs {
						out = append(out, model.TriggerParamOutput{
							Title: ot.Name, ParamKey: ot.Key,
						})
					}
				}
			}
		}
	}

	return
}
