package alarm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/queues"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/iotModel"
	"strconv"
	"time"
)

// Check 告警检测
func (s *sAlarmRule) Check(ctx context.Context, productKey string, deviceKey string, triggerType int, param any, subKey ...string) (err error) {
	//g.Log().Debugf(ctx, "alarm_rule_check: deviceKey(%s), productKey(%s), triggerType(%d), param(%v)", deviceKey, productKey, triggerType, param)

	// 网关子设备
	if len(subKey) > 0 {
		skey := subKey[0]
		sub, _ := dcache.GetDeviceDetailInfo(skey)
		productKey = sub.Product.Key
		deviceKey = sub.Key
	}

	//重组param数据
	eventKey, data, err := s.getParamData(ctx, param)
	if data == nil || err != nil {
		return
	}

	//获取规则列表
	rules, err := s.getAlarmRuleList(ctx, productKey, deviceKey, triggerType, eventKey)
	if len(rules) == 0 || err != nil {
		return
	}

	logData, err := json.Marshal(param)
	if err != nil {
		g.Log().Errorf(ctx, "告警规则数据 - %s - %s：%s", productKey, deviceKey, err)
		return
	}

	// 填充告警触发时间
	if ts, ok := data["CreateTime"]; ok {
		switch tt := ts.(type) {
		case string:
			if t, err := time.Parse("2006-01-02 15:04:05", tt); err == nil {
				data["CreateTime"] = t.Unix()
			} else {
				data["CreateTime"] = time.Now().Unix()
			}
		case int, int32, int64, float32, float64:
			data["CreateTime"] = tt.(int64)
		}
	} else {
		data["CreateTime"] = time.Now().Unix()
	}

	for _, r := range rules {
		go func(rule model.AlarmRuleOutput) {
			exp := s.expression(ctx, rule)
			if exp != "" {
				gov, err := govaluate.NewEvaluableExpression(exp)
				if err != nil {
					g.Log().Errorf(ctx, "告警表达式 - %s - %s - %s：%s", productKey, deviceKey, exp, err)
					return
				}
				if gov == nil {
					return
				}
				rs, err := gov.Evaluate(data)
				if err != nil {
					g.Log().Errorf(ctx, "告警表达式参数 - %s - %s：%s - %v", productKey, deviceKey, err, data)
					return
				}
				if rs == nil {
					return
				}
				if y, ok := rs.(bool); ok && y {
					// 写告警日志
					log := model.AlarmLogAddInput{
						Type:       1,
						RuleId:     rule.Id,
						RuleName:   rule.Name,
						Level:      rule.Level,
						Data:       string(logData),
						Expression: exp,
						ProductKey: productKey,
						DeviceKey:  deviceKey,
					}

					if log.ProductKey == "" {
						return
					}

					// 写入队列
					logData, _ := json.Marshal(log)
					err = queues.ScheduledDeviceAlarmLog.Push(ctx, consts.QueueDeviceAlarmLogTopic, logData, 10)

					key := consts.DeviceAlarmLogPrefix + productKey + deviceKey + exp
					err = cache.Instance().Set(ctx, key, log, 60*time.Second)
					if err != nil {
						return
					}
					//告警执行动作
					s.doAction(ctx, rule, exp, deviceKey, param)
				}
			}
		}(r)
	}
	return
}

// getParamData 获取Param数据 eventKey 事件标识
func (s *sAlarmRule) getParamData(ctx context.Context, param any) (eventKey string, res map[string]any, err error) {
	var (
		data = make(map[string]any) // 上传数据
	)
	// 上传数据重组
	switch pd := param.(type) {
	case iotModel.ReportPropertyData:
		for k, v := range pd {
			vv := gconv.String(v.Value)
			if gstr.IsNumeric(vv) {
				data[k] = gconv.Float64(v.Value)
			} else if gt, err := gtime.StrToTime(vv); err == nil {
				data[k] = gt.Unix()
			} else {
				data[k] = v.Value
			}
			data[k+"_time"] = v.CreateTime
		}
	case iotModel.ReportEventData:
		for k, v := range pd.Param.Value {
			vv := gconv.String(v)
			if gstr.IsNumeric(vv) {
				data[k] = gconv.Float64(v)
			} else if gt, err := gtime.StrToTime(vv); err == nil {
				data[k] = gt.Unix()
			} else {
				data[k] = v
			}
		}
		data["CreateTime"] = pd.Param.CreateTime
		eventKey = pd.Key
	case iotModel.ReportStatusData:
		data["Status"] = pd.Status
		data["CreateTime"] = pd.CreateTime
	default:
		return "", nil, gerror.New("数据格式错误")
	}
	return eventKey, data, nil
}

// 获取告警规则
func (s *sAlarmRule) getAlarmRuleList(ctx context.Context, productKey, deviceKey string, triggerType int, eventKey string) (res []model.AlarmRuleOutput, err error) {
	// 获取规则列表
	AlarmRuleList := dcache.GetDeviceAlarm(ctx, productKey)
	if AlarmRuleList == nil {
		return
	}
	// 根据触发类型，过滤告警规则
	for _, r := range AlarmRuleList {
		if r.Status == 0 {
			return
		}
		if r.TriggerCondition != "" {
			conditionErr := json.Unmarshal([]byte(r.TriggerCondition), &r.Condition)
			if conditionErr != nil {
				return nil, err
			}
		}
		if r.Action != "" {
			performActionErr := json.Unmarshal([]byte(r.Action), &r.PerformAction)
			if err != nil {
				return nil, performActionErr
			}
		}
		if triggerType == consts.AlarmTriggerTypeEvent && eventKey != r.EventKey {
			continue
		}
		if triggerType != r.TriggerType {
			continue
		}
		if r.DeviceKey == deviceKey || r.DeviceKey == "all" || r.DeviceKey == "" {
			res = append(res, r)
		}
	}

	return
}

// expression 告警条件表达式
func (s *sAlarmRule) expression(ctx context.Context, rule model.AlarmRuleOutput) string {
	glen := len(rule.Condition.TriggerCondition)
	var exp string
	for _, group := range rule.Condition.TriggerCondition {
		flen := len(group.Filters)
		var gexp string
		for _, v := range group.Filters {
			if (v.Operator == consts.OperatorBet ||
				v.Operator == consts.OperatorNbet) && len(v.Value) < 2 {
				continue
			}

			// 如果条件参数是上报时间，则将参数值转换成时间戳
			if v.Key == "sysReportTime" {
				v.Key = "CreateTime"
				if v.Value[0] != "" {
					if t, err := time.Parse("2006-01-02 15:04:05", v.Value[0]); err == nil {
						v.Value[0] = strconv.FormatInt(t.Unix(), 10)
					}
				}
				if len(v.Value) > 1 && v.Value[1] != "" {
					if t, err := time.Parse("2006-01-02 15:04:05", v.Value[1]); err == nil {
						v.Value[1] = strconv.FormatInt(t.Unix(), 10)
					}
				}
			}

			// 表达式生成
			var fexp string
			switch v.Operator {
			case consts.OperatorEq:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "==", v.Value[0])
			case consts.OperatorNe:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "!=", v.Value[0])
			case consts.OperatorGt:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, ">", v.Value[0])
			case consts.OperatorGte:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, ">=", v.Value[0])
			case consts.OperatorLt:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "<", v.Value[0])
			case consts.OperatorLte:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "<=", v.Value[0])
			case consts.OperatorBet:
				fexp = fmt.Sprintf("(%s >= %s && %s <= %s)", v.Key, v.Value[0], v.Key, v.Value[1])
			case consts.OperatorNbet:
				fexp = fmt.Sprintf("(%s < %s && %s > %s)", v.Key, v.Value[0], v.Key, v.Value[1])
			}
			if flen > 1 {
				switch v.AndOr {
				case 1:
					fexp = " && " + fexp
				case 2:
					fexp = " || " + fexp
				}
			}
			gexp += fexp
		}
		if flen > 1 {
			gexp = "(" + gexp + ")"
		}
		if glen > 1 {
			switch group.AndOr {
			case 1:
				gexp = " && " + gexp
			case 2:
				gexp = " || " + gexp
			}
		}
		exp += gexp
	}
	return exp
}
