package command

import (
	"context"
	"sagooiot/internal/mqtt"
	"sagooiot/internal/queues"
	_ "sagooiot/internal/queues"
	"sagooiot/module"
	"sagooiot/pkg/worker"
)

func RunQueue(ctx context.Context) (error, func(context.Context) error) {
	worker.TasksInstance() //启用系统的任务队列
	queues.Run()           //启用系统的消息队列
	module.WorkerRun()     //启用模块应用的任务
	return nil, nil
}

func wrapperMqtt(ctx context.Context) (error, func(context.Context) error) {
	if err := mqtt.InitSystemMqtt(); err != nil {
		return err, nil
	}
	return nil, func(ctx context.Context) error {
		mqtt.Close()
		return nil
	}
}

type DeferFunc struct {
	F    func(ctx context.Context) (error, func(context.Context) error)
	Desc string
}
