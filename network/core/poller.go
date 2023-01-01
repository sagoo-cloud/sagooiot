package core

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"github.com/sagoo-cloud/sagooiot/network/pkg/cron"
)

// Poller 采集器
type Poller struct {
	model.Poller
	Device *Device

	reading bool
	job     *cron.Job
}

// Start 启动
func (p *Poller) Start(ctx context.Context) (err error) {
	if p.job != nil {
		p.job.Cancel()
		//return errors.New("已经启动")
	}
	p.job, err = cron.Interval(p.Interval, func() {
		p.Execute(ctx)
	})
	return
}

// Execute 执行
func (p *Poller) Execute(ctx context.Context) {
	if p.reading {
		return
	}
	go p.read(ctx)
}

func (p *Poller) read(ctx context.Context) {
	p.reading = true
	err := p.Device.read(p.Length)
	p.reading = false

	if err != nil {
		//log error
		g.Log().Error(ctx, err)
	}
}

// Stop 结束
func (p *Poller) Stop() {
	if p.job != nil {
		p.job.Cancel()
	}
}
