package router

import (
	"context"
	envirotronicsController "github.com/sagoo-cloud/sagooiot/internal/controller/envirotronics"
	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Business 业务相关功能的路由
func Business(ctx context.Context, group *ghttp.RouterGroup) {
	//环测相关路由
	group.Group("/envirotronics", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			envirotronicsController.Weather, // 天气监测
		)
	})
}
