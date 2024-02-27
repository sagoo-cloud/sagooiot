// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"sagooiot/internal/model"
)

type (
	IAlarmLevel interface {
		Detail(ctx context.Context, level uint) (out model.AlarmLevelOutput, err error)
		All(ctx context.Context) (out *model.AlarmLevelListOutput, err error)
		Edit(ctx context.Context, in []*model.AlarmLevelEditInput) (err error)
	}
	IAlarmLog interface {
		Detail(ctx context.Context, id uint64) (out *model.AlarmLogOutput, err error)
		Add(ctx context.Context, in *model.AlarmLogAddInput) (id uint64, err error)
		List(ctx context.Context, in *model.AlarmLogListInput) (out *model.AlarmLogListOutput, err error)
		Handle(ctx context.Context, in *model.AlarmLogHandleInput) (err error)
		TotalForLevel(ctx context.Context) (total []model.AlarmLogLevelTotal, err error)
		// ClearLogByDays 按日期删除日志
		ClearLogByDays(ctx context.Context, days int) (err error)
	}
	IAlarmRule interface {
		List(ctx context.Context, in *model.AlarmRuleListInput) (out *model.AlarmRuleListOutput, err error)
		// TODO 废弃
		Cache(ctx context.Context) (rs map[string][]model.AlarmRuleOutput, err error)
		// CacheAllAlarmRule 缓存所有的告警规则
		CacheAllAlarmRule(ctx context.Context) (err error)
		// Detail 获取告警规则详情
		Detail(ctx context.Context, id uint64) (out *model.AlarmRuleOutput, err error)
		Add(ctx context.Context, in *model.AlarmRuleAddInput) (err error)
		Edit(ctx context.Context, in *model.AlarmRuleEditInput) (err error)
		// Deploy 启用告警规则
		Deploy(ctx context.Context, id uint64) (err error)
		Undeploy(ctx context.Context, id uint64) (err error)
		Del(ctx context.Context, id uint64) (err error)
		Operator(ctx context.Context) (out []model.OperatorOutput, err error)
		TriggerType(ctx context.Context, productKey string) (out []model.TriggerTypeOutput, err error)
		TriggerParam(ctx context.Context, productKey string, triggerType int, eventKey ...string) (out []model.TriggerParamOutput, err error)
		// NoticeAction 执行告警通知
		NoticeAction(ctx context.Context, rule model.AlarmRuleOutput, expression string, deviceKey string, param any)
		// Check 告警检测
		Check(ctx context.Context, productKey string, deviceKey string, triggerType int, param any, subKey ...string) (err error)
		// AddCronRule 添加定时触发规则
		AddCronRule(ctx context.Context, in *model.AlarmCronRuleAddInput) (err error)
		// EditCronRule 编辑定时触发规则
		EditCronRule(ctx context.Context, in *model.AlarmCronRuleEditInput) (err error)
	}
)

var (
	localAlarmLevel IAlarmLevel
	localAlarmLog   IAlarmLog
	localAlarmRule  IAlarmRule
)

func AlarmLevel() IAlarmLevel {
	if localAlarmLevel == nil {
		panic("implement not found for interface IAlarmLevel, forgot register?")
	}
	return localAlarmLevel
}

func RegisterAlarmLevel(i IAlarmLevel) {
	localAlarmLevel = i
}

func AlarmLog() IAlarmLog {
	if localAlarmLog == nil {
		panic("implement not found for interface IAlarmLog, forgot register?")
	}
	return localAlarmLog
}

func RegisterAlarmLog(i IAlarmLog) {
	localAlarmLog = i
}

func AlarmRule() IAlarmRule {
	if localAlarmRule == nil {
		panic("implement not found for interface IAlarmRule, forgot register?")
	}
	return localAlarmRule
}

func RegisterAlarmRule(i IAlarmRule) {
	localAlarmRule = i
}
