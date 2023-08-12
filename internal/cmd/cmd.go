package cmd

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/cmd/router"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/mqtt"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/sagoo-cloud/sagooiot/internal/task"
	"github.com/sagoo-cloud/sagooiot/network"
	"github.com/sagoo-cloud/sagooiot/utility/notifier"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/util/gmode"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start sagoo-iot server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			var (
				s = g.Server()
			)

			// 静态目录设置
			uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
			if uploadPath == "" {
				g.Log().Fatal(ctx, "文件上传配置路径不能为空")
			}
			//s.AddStaticPath("/upload", uploadPath)

			// HOOK, 开发阶段禁止浏览器缓存,方便调试
			if gmode.IsDevelop() {
				s.BindHookHandler("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
					r.Response.Header().Set("Cache-Control", "no-store")
				})
			}

			//操作日志
			s.BindHookHandler("/*", ghttp.HookAfterOutput, func(r *ghttp.Request) {
				service.Middleware().OperationLog(r)
			})

			//sse 实时数据推送
			s.Group("/subscribe", func(group *ghttp.RouterGroup) {
				group.GET("/sysenv", notifier.SysenvMessageEvent)
			})

			s.Group("/api/v1", func(group *ghttp.RouterGroup) {

				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
					service.Middleware().MiddlewareCORS,
				)

				router.System(ctx, group)   //系统默认功能的路由
				router.Business(ctx, group) //业务专属功能的路由

			})

			//初始化系统权限缓存
			if err := service.SysAuthorize().InitAuthorize(ctx); err != nil {
				g.Log().Fatal(ctx, "系统权限缓存初始化失败：", err)
			}

			// 初始化插件配置数据
			if err := service.SystemPluginsConfig().UpdateAllPluginsConfigCache(); err != nil {
				g.Log().Error(ctx, "初始化插件配置数据失败：", err)
			}

			// TDengine 初始化
			if err := service.TSLTable().CreateDatabase(ctx); err != nil {
				g.Log().Fatal(ctx, "TDengine 数据库创建失败：", err)
			}
			if err := service.TdLogTable().CreateStable(ctx); err != nil {
				g.Log().Fatal(ctx, "TDengine 日志超级表创建失败：", err)
			}
			if err := mqtt.InitSystemMqtt(); err != nil {
				g.Log().Errorf(ctx, "MQTT 初始化mqtt客户端失败,失败原因:%+#v", err)
			}
			defer mqtt.Close()
			// 启动失败的话请注释掉
			if err := network.ReloadNetwork(context.Background()); err != nil {
				g.Log().Errorf(ctx, "载入网络错误,错误原因:%+#v", err)
			}

			//初始化任务
			task.StartInit()

			// 自定义丰富文档
			enhanceOpenAPIDoc(s)
			// 启动Http Server
			s.Run()
			return
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info.Title = `sagooAdmin Project`
	openapi.Info.Description = ``

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		{Name: consts.OpenAPITagNameLogin},
		{Name: consts.OpenAPITagNameOrganization},
		{Name: consts.OpenAPITagNameDept},
		{Name: consts.OpenAPITagNamePost},
		{Name: consts.OpenAPITagNameRole},
		{Name: consts.OpenAPITagNameUser},
		{Name: consts.OpenAPITagNameMenu},
		{Name: consts.OpenAPITagNameApi},
		{Name: consts.OpenAPITagNameAuthorize},
	}
}
