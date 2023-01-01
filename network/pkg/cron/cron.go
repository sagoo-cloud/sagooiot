package cron

import (
	"github.com/go-co-op/gocron"
	"time"
)

// Scheduler 调度器
var Scheduler *gocron.Scheduler

func init() {
	Scheduler = gocron.NewScheduler(time.Local) // time.UTC
	Scheduler.StartAsync()
}

//是否是使用单一协程？？？ 是则要改成协程池？？？

// Schedule 创建任务
func Schedule(crontab string, fn func()) (*Job, error) {
	job, err := Scheduler.Cron(crontab).Do(fn)
	if err != nil {
		return nil, err
	}
	return &Job{job: job}, nil
}

// Interval 创建周期任务
func Interval(interval int, fn func()) (*Job, error) {
	job, err := Scheduler.Every(interval).Milliseconds().Do(fn)
	if err != nil {
		return nil, err
	}
	return &Job{job: job}, nil
}

// Clock 创建每日任务
func Clock(hours int, minutes int, fn func()) (*Job, error) {
	job, err := Scheduler.Every(1).Day().At(hours).Hours().At(minutes).Minutes().Do(fn)
	if err != nil {
		return nil, err
	}
	return &Job{job: job}, nil
}

// ClockWithWeekdays 创建每日任务
func ClockWithWeekdays(hours int, minutes int, weekdays []time.Weekday, fn func()) (*Job, error) {
	s := Scheduler.Every(1).Day().At(hours).Hours().At(minutes).Minutes()
	if len(weekdays) > 0 {
		s.Weeks()
		for _, w := range weekdays {
			s.Weekday(w)
		}
	}

	job, err := s.Do(fn)
	if err != nil {
		return nil, err
	}
	return &Job{job: job}, nil
}

// Job 任务
type Job struct {
	job *gocron.Job
}

// Cancel 取消任务
func (j *Job) Cancel() {
	Scheduler.Remove(j.job)
}
