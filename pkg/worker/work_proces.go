package worker

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

// Scheduled 任务调度器
type Scheduled struct {
	topic string
	w     *Worker
}

// Process 任务具体处理过程接口，实现该接口即可加入到任务队列中
type Process interface {
	GetTopic() string                                  // 获取消费主题
	Handle(ctx context.Context, p Payload) (err error) // 处理过程的方法
}

func RegisterProcess(p Process) (s *Scheduled) {
	s = &Scheduled{}
	topic := p.GetTopic()
	if topic != "" {
		s.w = New(
			WithGroup(topic),
			WithHandler(p.Handle),
		)
	} else {
		s.w = New(
			WithHandler(p.Handle),
		)
	}
	return
}

// Push 采用消息队列的方式执行任务
func (s *Scheduled) Push(ctx context.Context, topic string, data []byte, timeout int) (err error) {
	err = s.w.Once(
		WithRunUuid(guid.S()),
		WithRunPayload(data), //传递参数
		WithRunGroup(topic),
		//task.WithRunAt(time.Now().Add(time.Duration(10)*time.Second)), //延迟5秒执行
		WithRunTimeout(timeout), //超时时间10秒
	)
	if err != nil {
		g.Log().Debug(ctx, "Run Queue TaskWorker %s Error: %v", topic, err)
	}
	return
}

// Cron 采用定时任务的方式执行任务
func (s *Scheduled) Cron(ctx context.Context, topic, cronExpr string, data []byte) (err error) {
	s.topic = topic
	err = s.w.Cron(
		WithRunUuid(topic),
		WithRunGroup(topic),
		WithRunExpr(cronExpr),
		WithRunPayload(data), //传递参数
	)
	if err != nil {
		g.Log().Debug(ctx, "Run Cron TaskWorker %s Error: %v", topic, err)
	}
	return
}
