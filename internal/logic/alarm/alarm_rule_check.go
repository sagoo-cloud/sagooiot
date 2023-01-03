package alarm

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sagoo-cloud/sagooiot/extend"
	extModel "github.com/sagoo-cloud/sagooiot/extend/model"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/utility/notifier"
	"github.com/sagoo-cloud/sagooiot/utility/utils"
	"strconv"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 告警匹配检测
func (s *sAlarmRule) Check(ctx context.Context, productKey string, deviceKey string, data map[string]any) (err error) {
	list, err := s.Cache(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "告警规则缓存获取 - %s - %s：%s", productKey, deviceKey, err)
		return
	}
	if len(list) == 0 {
		return
	}
	pList, ok := list[productKey]
	if !ok {
		return
	}

	logData, err := json.Marshal(data)
	if err != nil {
		g.Log().Errorf(ctx, "告警规则数据 - %s - %s：%s", productKey, deviceKey, err)
		return
	}

	// 填充告警触发时间
	if ts, ok := data["ts"]; ok {
		if t, err := time.Parse("2006-01-02 15:04:05", ts.(string)); err == nil {
			data["ts"] = t.Unix()
		}
	} else {
		data["ts"] = time.Now().Unix()
	}

	for _, rule := range pList {
		if rule.DeviceKey == deviceKey {
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
						log := &model.AlarmLogAddInput{
							Type:       1,
							RuleId:     rule.Id,
							RuleName:   rule.Name,
							Level:      rule.Level,
							Data:       string(logData),
							ProductKey: productKey,
							DeviceKey:  deviceKey,
						}
						logId, err := service.AlarmLog().Add(ctx, log)
						if err != nil {
							g.Log().Errorf(ctx, "告警日志写入 - %s - %s：%s", productKey, deviceKey, err)
							return
						}

						// 触发告警通知
						nt := notifier.NewNotifier(3 * time.Second)
						nt.SetCallbacks(
							func(state notifier.State) {
								s.doAction(ctx, rule, exp)
							},
							func(state notifier.State) {
								s.doAction(ctx, rule, exp)
							},
							func(state notifier.State) {
								s.doAction(ctx, rule, exp)
							})
						// 频率控制
						i := 0
						for {
							i++
							if i == 6 || i == 12 {
								alarmLog, _ := service.AlarmLog().Detail(ctx, logId)
								if alarmLog == nil {
									break
								}
								// 告警未处理，触发通知
								if alarmLog.Status == model.AlarmLogStatusUnhandle {
									nt.Trigger(true)
								}
							}
							if i == 9 || i == 15 {
								alarmLog, _ := service.AlarmLog().Detail(ctx, logId)
								if alarmLog == nil {
									break
								}
								// 告警已处理或忽略，停止触发通知
								if alarmLog.Status == model.AlarmLogStatusHandle ||
									alarmLog.Status == model.AlarmLogStatusIgnore {
									nt.Trigger(false)
								}
							}

							time.Sleep(1 * time.Second)
							if i > 20 {
								break
							}
						}
					}
				}
			}(rule)
		}
	}
	return
}

// 告警执行动作
func (s *sAlarmRule) doAction(ctx context.Context, rule model.AlarmRuleOutput, expression string) {
	if len(rule.PerformAction.Action) == 0 {
		return
	}

	for _, v := range rule.PerformAction.Action {
		if v.NoticeTemplate == "" {
			continue
		}

		// 获取告警模板
		tpl, err := service.NoticeTemplate().GetNoticeTemplateById(ctx, v.NoticeTemplate)
		if err != nil {
			g.Log().Errorf(ctx, "告警获取通知模板 - %s ：%s", v.NoticeTemplate, err)
			continue
		}
		if tpl == nil {
			continue
		}

		// 获取告警级别名称
		level, err := service.AlarmLevel().Detail(ctx, rule.Level)
		if err != nil {
			g.Log().Errorf(ctx, "告警获取级别名称 - %s ：%s", v.NoticeTemplate, err)
		}

		// 获取设备、产品名称
		var (
			pname string
			dname string
		)
		d, err := service.DevDevice().Get(ctx, rule.DeviceKey)
		if err != nil {
			g.Log().Errorf(ctx, "告警获取设备信息 - %s ：%s", rule.DeviceKey, err)
		}
		if d != nil {
			pname = d.Product.Name
			dname = d.Name
		}

		// 模板解析
		content, err := utils.ReplaceTemplate(tpl.Content, map[string]any{
			"Level":   level.Name,
			"Product": pname,
			"Device":  dname,
			"Rule":    rule.Name + " " + expression,
		})
		if err != nil {
			g.Log().Errorf(ctx, "告警模板解析 - %s ：%s", v.NoticeTemplate, err)
		}

		// 告警消息发送
		var msg = extModel.NoticeInfoData{}
		msg.MsgTitle = tpl.Title
		msg.MsgBody = content

		for _, u := range v.Addressee {
			if extend.GetNoticePlugin() != nil {
				msg.Totag = fmt.Sprintf(`[{"name":"%s","value":"%s"}]`, tpl.SendGateway, u)
				noticeStatus := 1
				noticeFail := ""
				_, err = extend.GetNoticePlugin().NoticeSend(tpl.SendGateway, msg)
				if err != nil {
					noticeStatus = 0
					noticeFail = err.Error()
					g.Log().Errorf(ctx, "告警通知发送 - %s ：%s", tpl.SendGateway, err)
				}

				// 通知日志记录
				if err = service.NoticeLog().Add(ctx, &model.NoticeLogAddInput{
					TemplateId:  tpl.Id,
					SendGateway: tpl.SendGateway,
					Addressee:   u,
					Title:       tpl.Title,
					Content:     content,
					Status:      noticeStatus,
					FailMsg:     noticeFail,
					SendTime:    gtime.Now().String(),
				}); err != nil {
					g.Log().Errorf(ctx, "告警通知日志记录：%v", err)
				}
			}
		}
	}
}

// 告警条件表达式生成
func (s *sAlarmRule) expression(ctx context.Context, rule model.AlarmRuleOutput) string {
	glen := len(rule.Condition.TriggerCondition)
	var exp string
	for _, group := range rule.Condition.TriggerCondition {
		flen := len(group.Filters)
		var gexp string
		for _, v := range group.Filters {
			if (v.Operator == model.OperatorBet ||
				v.Operator == model.OperatorNbet) && len(v.Value) < 2 {
				continue
			}

			// 如果条件参数是上报时间，则将参数值转换成时间戳
			if v.Key == "sysReportTime" {
				v.Key = "ts"
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
			case model.OperatorEq:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "==", v.Value[0])
			case model.OperatorNe:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "!=", v.Value[0])
			case model.OperatorGt:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, ">", v.Value[0])
			case model.OperatorGte:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, ">=", v.Value[0])
			case model.OperatorLt:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "<", v.Value[0])
			case model.OperatorLte:
				fexp = fmt.Sprintf("(%s %s %s)", v.Key, "<=", v.Value[0])
			case model.OperatorBet:
				fexp = fmt.Sprintf("(%s >= %s && %s <= %s)", v.Key, v.Value[0], v.Key, v.Value[1])
			case model.OperatorNbet:
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
