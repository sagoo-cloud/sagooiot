package alarm

import (
	"context"
	"encoding/json"
	"fmt"
	"sagooiot/internal/consts"
	"sagooiot/internal/dao"
	"sagooiot/internal/model"
	"sagooiot/internal/model/do"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/util/gconv"
)

// AddCronRule 添加定时触发规则
func (s *sAlarmRule) AddCronRule(ctx context.Context, in *model.AlarmCronRuleAddInput) (err error) {
	//获取当前登录用户ID
	loginUserId := service.Context().GetUserId(ctx)

	triggerCondition, _ := json.Marshal(in.AlarmCronCondition)
	action, _ := json.Marshal(in.AlarmPerformAction)

	_, err = dao.AlarmRule.Ctx(ctx).Data(do.AlarmRule{
		Name:             in.Name,
		Level:            in.Level,
		ProductKey:       in.ProductKey,
		DeviceKey:        in.DeviceKey,
		TriggerMode:      2,
		TriggerCondition: triggerCondition,
		Action:           action,
		Status:           0,
		CreatedBy:        uint(loginUserId),
	}).Insert()
	if err != nil {
		return
	}

	return
}

// EditCronRule 编辑定时触发规则
func (s *sAlarmRule) EditCronRule(ctx context.Context, in *model.AlarmCronRuleEditInput) (err error) {
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
	param.TriggerCondition, _ = json.Marshal(in.AlarmCronCondition)
	param.Action, _ = json.Marshal(in.AlarmPerformAction)

	_, err = dao.AlarmRule.Ctx(ctx).Data(param).Where(dao.AlarmRule.Columns().Id, in.Id).Update()
	if err != nil {
		return
	}

	return
}

// start 启动定时
func (s *sAlarmRule) start(ctx context.Context, id uint64) error {
	rule, err := s.Detail(ctx, id)
	if err != nil {
		return err
	}
	if rule == nil || rule.Status == consts.AlarmRuleStatusOn {
		return gerror.New("告警规则不存在，或已启用")
	}
	if rule.TriggerMode == consts.AlarmTriggerModeDevice {
		return nil
	}

	expList := rule.CronCondition.CronCondition
	if len(expList) == 0 {
		return gerror.New("请设置定时触发条件")
	}

	for i, ex := range expList {
		name := fmt.Sprintf("alarm-cron-%d-%d", id, i+1)

		_, err := gcron.AddSingleton(ctx, ex, func(ctx context.Context) {
			// 执行动作：发送告警通知
			s.notice(ctx, *rule, ex, "", nil)
		}, name)
		if err != nil {
			return err
		}

		gcron.Start(name)
	}

	return nil
}

// stop 停止定时
func (s *sAlarmRule) stop(ctx context.Context, id uint64) error {
	rule, err := s.Detail(ctx, id)
	if err != nil {
		return err
	}
	if rule == nil || rule.Status == consts.AlarmRuleStatusOff {
		return gerror.New("告警规则不存在，或已停用")
	}
	if rule.TriggerMode == consts.AlarmTriggerModeDevice {
		return nil
	}

	expList := rule.CronCondition.CronCondition
	if len(expList) == 0 {
		return nil
	}

	for i := 0; i < len(expList); i++ {
		name := fmt.Sprintf("alarm-cron-%d-%d", id, i+1)
		gcron.Remove(name)
	}

	return nil
}
