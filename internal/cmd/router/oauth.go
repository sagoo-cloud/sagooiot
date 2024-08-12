package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	oauthController "sagooiot/internal/controller/oauth"
)

// Analysis 分析统计相关的接口
func OAuth(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/oauth", func(group *ghttp.RouterGroup) {
		//group.Middleware(service.Middleware().Auth)
		group.Bind(
			oauthController.OProvider, // 第三方授权配置提供
			oauthController.OUser,     // 第三方授权用户登录
		)
	})
}
