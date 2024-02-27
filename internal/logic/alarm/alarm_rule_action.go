package alarm

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/dcache"
	"sagooiot/pkg/plugins"
	extModel "sagooiot/pkg/plugins/model"
	"sagooiot/pkg/utility/notifier"
	"sagooiot/pkg/utility/utils"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// doAction 告警执行动作
func (s *sAlarmRule) doAction(ctx context.Context, rule model.AlarmRuleOutput, expression string, deviceKey string, param any) {
	// 执行告警通知
	s.NoticeAction(ctx, rule, expression, deviceKey, param)
}

// NoticeAction 执行告警通知
func (s *sAlarmRule) NoticeAction(ctx context.Context, rule model.AlarmRuleOutput, expression string, deviceKey string, param any) {
	nt := notifier.NewNotifier(3 * time.Second)
	nt.SetCallbacks(
		func(state notifier.State) {
			s.notice(ctx, rule, expression, deviceKey, param)
		},
		func(state notifier.State) {
			s.notice(ctx, rule, expression, deviceKey, param)
		},
		func(state notifier.State) {
			s.notice(ctx, rule, expression, deviceKey, param)
		},
	)
	// 频率控制
	i := 0
	for {
		i++
		if i == 6 || i == 12 {
			key := consts.DeviceAlarmLogPrefix + rule.ProductKey + deviceKey + expression
			alarmLogData, err := cache.Instance().Get(ctx, key)
			if err != nil {
				g.Log().Errorf(ctx, "告警获取日志 - %s ：%s", key, err)
			}
			if alarmLogData.Val() == nil {
				break
			}
			var alarmLog model.AlarmLogOutput
			err = gconv.Scan(alarmLogData.Val(), &alarmLog)
			if err != nil {
				return
			}

			// 告警未处理，触发通知
			if alarmLog.Status == model.AlarmLogStatusUnhandle {
				nt.Trigger(true)
			}
		}
		if i == 9 || i == 15 {
			key := consts.DeviceAlarmLogPrefix + rule.ProductKey + deviceKey + expression
			alarmLogData, err := cache.Instance().Get(ctx, key)

			var alarmLog model.AlarmLogOutput
			err = gconv.Scan(alarmLogData.Val(), &alarmLog)
			if err != nil {
				return
			}

			if alarmLog.AlarmLog == nil {
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

// notice 发送告警通知
func (s *sAlarmRule) notice(ctx context.Context, rule model.AlarmRuleOutput, expression string, deviceKey string, param any) {
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

		device, err := dcache.GetDeviceDetailInfo(deviceKey)
		if err != nil {
			g.Log().Errorf(ctx, "告警获取设备信息 - %s ：%s", rule.DeviceKey, err)
		}
		if device == nil {
			return
		}

		// 模板数据准备，模板变量替换
		var contentData = make(map[string]any)
		contentData["Level"] = level.Name
		contentData["ProductName"] = device.ProductName
		contentData["ProductKey"] = device.Product.Key
		contentData["DeviceName"] = device.Name
		contentData["DeviceKey"] = deviceKey
		contentData["Rule"] = rule.Name + " " + expression
		if param != nil {
			for k, v := range gconv.Map(param) {
				var valueData model.ReportPropertyNode
				err := gconv.Scan(v, &valueData)
				if err != nil {
					continue
				}
				contentData[gconv.String(k)] = valueData.Value
				contentData[gconv.String(k)+"_time"] = gtime.New(valueData.CreateTime).Format("Y-m-d H:i:s")
			}
		}
		// 模板解析
		content, err := utils.ReplaceTemplate(tpl.Content, contentData)
		if err != nil {
			g.Log().Errorf(ctx, "告警模板解析 - %s ：%s", v.NoticeTemplate, err)
		}

		// 告警消息发送
		var msg = extModel.NoticeInfoData{}
		msg.TemplateCode = tpl.Code
		msg.MsgTitle = tpl.Title
		msg.MsgBody = content

		if plugins.GetNoticePlugin() != nil {
			var toTag []extModel.NoticeSendObject
			for _, u := range v.Addressee {
				toTag = append(toTag, extModel.NoticeSendObject{
					Name:  tpl.SendGateway,
					Value: u,
				})
			}
			msg.Totag = toTag
			noticeStatus := 0
			noticeResMsg := ""
			sendRes, err := plugins.GetNoticePlugin().NoticeSend(tpl.SendGateway, msg)
			if err != nil {
				noticeStatus = 0
				noticeResMsg = err.Error()
				g.Log().Errorf(ctx, "告警通知发送 - %s ：%s", tpl.SendGateway, err)
			} else {
				noticeResMsg = sendRes.Message
				if sendRes.Code != 0 { //如果发送返回的code不为0，说明发送失败
					noticeStatus = 0
				} else {
					noticeStatus = 1
				}
			}

			// 通知日志记录
			if err = service.NoticeLog().Add(ctx, &model.NoticeLogAddInput{
				TemplateId:  tpl.Id,
				SendGateway: tpl.SendGateway,
				Addressee:   strings.Join(v.Addressee, ""),
				Title:       tpl.Title,
				Content:     content,
				Status:      noticeStatus,
				FailMsg:     noticeResMsg,
				SendTime:    gtime.Now(),
			}); err != nil {
				g.Log().Errorf(ctx, "告警通知日志记录：%v", err)
			}
		}
	}
}
