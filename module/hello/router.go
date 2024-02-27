package hello

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	helloController "sagooiot/module/hello/controller/hello"
	_ "sagooiot/module/hello/logic/hello"
)

// Router 加载模块的路由
func Router(ctx context.Context, group *ghttp.RouterGroup) {
	//访问的地址为：http://127.0.0.1:8199/api/v1/hello
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			helloController.ControllerV1{}, // 获取大屏项目数据
		)
	})

}
