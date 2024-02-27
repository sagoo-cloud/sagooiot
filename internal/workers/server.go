package workers

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hibiken/asynq"
	"sagooiot/internal/workers/initialize"
	"sagooiot/internal/workers/middleware"
)

func Run() {
	addr := g.Cfg().MustGet(context.Background(), "redis.default.address").String()
	db := g.Cfg().MustGet(context.Background(), "redis.default.db").Int()
	user := g.Cfg().MustGet(context.Background(), "redis.default.user").String()
	if user == "" {
		user = "default"
	}
	pass := g.Cfg().MustGet(context.Background(), "redis.default.pass", 0).String()

	concurrency := g.Cfg().MustGet(context.Background(), "redis.default.pass", 10).Int()

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: addr, DB: db, Password: pass, Username: user},
		asynq.Config{Concurrency: concurrency, Logger: NewLogger(g.Log())},
	)
	mux := asynq.NewServeMux()
	mux.Use(middleware.LoggingMiddleware)
	initialize.Handles(mux)
	if err := srv.Run(mux); err != nil {
		glog.Error(context.Background(), err.Error())
	}
}
