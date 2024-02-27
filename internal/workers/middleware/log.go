package middleware

import (
	"context"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hibiken/asynq"
)

func LoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		name := t.Type()
		glog.Debugf(ctx, "Start processing %s", name)
		err := h.ProcessTask(ctx, t)
		if err != nil {
			glog.Debugf(ctx, "Failure processing %s,Error: %s", name, err)
			return err
		}
		glog.Debugf(ctx, "Finished processing %s", name)
		return nil
	})
}
