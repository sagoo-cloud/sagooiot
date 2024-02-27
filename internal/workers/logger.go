package workers

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hibiken/asynq"
)

type Logger struct {
	base *glog.Logger
	ctx  context.Context
}

func NewLogger(l *glog.Logger) asynq.Logger {
	return &Logger{
		base: l,
		ctx:  context.Background(),
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.base.Debug(l.ctx, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.base.Info(l.ctx, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.base.Warning(l.ctx, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.base.Error(l.ctx, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.base.Fatal(l.ctx, args...)
}
