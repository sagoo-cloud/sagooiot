package notifier

import (
	"time"
)

type State struct {
	sentAt         *time.Time
	lastRemindedAt *time.Time
}

// LastRemindedAt 获取上一次提醒的时间
func (p *State) LastRemindedAt() time.Time {
	return *p.lastRemindedAt
}

// SentAt 获取首次触发的时间
func (p *State) SentAt() time.Time {
	return *p.sentAt
}

// FromSentAt 获取首次触发到现在的时间间隔
func (p *State) FromSentAt() time.Duration {
	if p.sentAt == nil {
		return 0
	}
	return time.Since(*p.sentAt)
}

// FromLastRemindedAt 获取上次提醒到现在的时间间隔
func (p *State) FromLastRemindedAt() time.Duration {
	if p.lastRemindedAt == nil {
		return 0
	}
	return time.Since(*p.lastRemindedAt)
}

type Callback func(state State)

func NewNotifier(remindDuration time.Duration) *Notifier {
	return &Notifier{
		RemindDuration: remindDuration,
	}
}

// Notifier 监控通知器
type Notifier struct {
	RemindDuration time.Duration // 提醒周期，在此时间段内，将不会触发提醒
	state          *State
	alertCallback  Callback
	remindCallback Callback
	repairCallback Callback
	initialed      bool
}

// SetCallbacks 设置通知器的触发回调函数，在里面实现告警、提醒、修复的消息触发
func (p *Notifier) SetCallbacks(alert, remind, repair Callback) {
	p.alertCallback = alert
	p.remindCallback = remind
	p.repairCallback = repair
}

// Trigger 是否触发报警，内部重置状态
func (p *Notifier) Trigger(trigger bool) {
	if trigger {
		if p.empty() {
			if p.alertCallback != nil {
				p.alertCallback(*p.getState())
			}
			p.sent()
		} else if p.remind() {
			if p.remindCallback != nil {
				p.remindCallback(*p.getState())
			}
		}
	} else if !p.empty() {
		if p.repairCallback != nil {
			p.repairCallback(*p.getState())
		}
		p.clear()
	}
}

func (p *Notifier) getState() *State {
	if p.state == nil {
		p.state = &State{}
	}
	return p.state
}

// 标记通知已发送状态
func (p *Notifier) sent() {
	now := time.Now()
	p.getState().sentAt = &now
}

// 判断通知器是否为空状态（未触发过的状态）
func (p *Notifier) empty() bool {
	return p.getState().sentAt == nil
}

// 重置通知器状态
func (p *Notifier) clear() {
	p.getState().sentAt = nil
	p.getState().lastRemindedAt = nil
}

//	获取是否到达发送提醒的时间点
//
// 获取再次触发后，静默周期将从头计算。
// 后续的触发动作有必要保证触发成功，否则将在下个周期获得重新触发的机会
func (p *Notifier) remind() bool {
	if p.getState().sentAt == nil {
		return false
	}
	if p.getState().lastRemindedAt == nil {
		p.getState().lastRemindedAt = p.getState().sentAt
	}
	if p.RemindDuration == 0 || p.getState().lastRemindedAt == nil {
		return false
	}
	if uint64(time.Since(*p.getState().lastRemindedAt)) > uint64(p.RemindDuration) {
		now := time.Now()
		p.getState().lastRemindedAt = &now
		return true
	}
	return false
}
