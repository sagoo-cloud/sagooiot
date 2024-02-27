package helper

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
)

// AppName 应用名称
func AppName(ctx context.Context) string {
	return g.Cfg().MustGet(ctx, "system.name", "sagooiot").String()
}

// SafeGo 安全的调用协程，发生错误时输出错误日志，不抛出panic
func SafeGo(ctx context.Context, f func(ctx context.Context)) {
	err := grpool.AddWithRecover(ctx, func(ctx context.Context) {
		f(ctx)
	}, func(ctx context.Context, err error) {
		g.Log().Errorf(ctx, "SafeGo exec failed:%+v", err)
	})

	if err != nil {
		g.Log().Errorf(ctx, "SafeGo AddWithRecover err:%+v", err)
		return
	}
}
