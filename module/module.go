package module

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"sagooiot/module/hello"
)

// Router 加载模块的普通路由
func Router(ctx context.Context, group *ghttp.RouterGroup) {
	hello.Router(ctx, group) // 加载hello模块路由

}

// OpenAPIRouter 加载模块的OpenAPI路由
func OpenAPIRouter(ctx context.Context, group *ghttp.RouterGroup) {
	hello.Router(ctx, group) // 加载hello模块路由

}

// NorthRouter 加载模块的北向路由
func NorthRouter(ctx context.Context, group *ghttp.RouterGroup) {
	hello.Router(ctx, group) // 加载hello模块路由

}

func WorkerRun() {

}
