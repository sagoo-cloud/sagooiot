package cron

import (
	"errors"
	"github.com/robfig/cron/v3"
)

var secondParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)

// ParseSpec 尝试转换定时任务表达式
func ParseSpec(spec string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("发生未知panic")
			}
		}
	}()
	_, err = secondParser.Parse(spec)
	return
}

// IsValidSpec 校验spec表达式是否合法
func IsValidSpec(spec string) (valid bool, err error) {
	err = ParseSpec(spec)
	if err == nil {
		//log.Infof("定时任务表达式(%+v)合法", spec)
		valid = true
	} else {
		//log.Errorf("定时任务表达式(%+v)非法", spec)
		valid = false
	}
	return
}
