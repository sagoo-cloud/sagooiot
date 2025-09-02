package command

import (
	"context"
	"expvar"
	"os"
	"os/signal"
	"runtime/debug"
	router2 "sagooiot/command/router"
	"sagooiot/internal/service"
	"sagooiot/internal/sse"
	"sagooiot/module"
	"sagooiot/pkg/utility"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
)

type SagooIoTApp struct {
	Server         *ghttp.Server
	ApiRouterGroup *ghttp.RouterGroup
}

func GetServer(ctx context.Context) SagooIoTApp {
	var s = g.Server()
	// 自定义丰富文档
	enhanceOpenAPIDoc(s)
	// 错误状态码接管
	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.Writeln("404 - 没有找到…")
	})
	s.BindStatusHandler(403, func(r *ghttp.Request) {
		r.Response.Writeln("403 - 拒绝显示")
	})

	// exp var 监控
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/debug/vars", ghttp.WrapH(expvar.Handler()))
	})

	// 静态目录设置
	uploadPath := g.Cfg().MustGet(ctx, "system.upload.path").String()
	if uploadPath == "" {
		g.Log().Fatal(ctx, "文件上传配置路径不能为空")
	}

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
		group.GET("/sysenv", sse.SysenvMessageEvent)
		group.GET("/redisinfo", sse.RedisInfoMessageEvent)
		group.GET("/mysqlinfo", sse.MysqlInfoMessageEvent)
		group.GET("/sysMessage", sse.SysMessageEntvt)
		group.GET("/logInfo", sse.LogInfoEvent)
	})

	apiV1 := s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Ctx,
			service.Middleware().ResponseHandler,
			service.Middleware().MiddlewareCORS,
			service.Middleware().I18n,
		)
		service.SysToken().GfToken().Middleware(group)
		router2.System(ctx, group)   //系统默认功能的路由
		router2.Iot(ctx, group)      //Iot功能的路由
		router2.Analysis(ctx, group) //分析统计功能的路由
		module.Router(ctx, group)    //加载模块的路由

	})

	// pprof性能分析
	enablePProf := g.Cfg().MustGet(context.Background(), "system.enablePProf").Bool()
	if enablePProf {
		// exp var 监控
		s.Group("/", func(group *ghttp.RouterGroup) {
			group.GET("/debug/vars", ghttp.WrapH(expvar.Handler()))
		})
		s.EnablePProf() //打开pprof性能分析工具，不需要的时候可以注掉
	}

	return SagooIoTApp{
		s,
		apiV1,
	}
}

func RunServer(ctx context.Context, stopSignal chan os.Signal, s *ghttp.Server) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				g.Log().Debugf(ctx, "RunServer: panic 产生，错误:%s", err)
			}
		}()

		//捕获panic意处出错，进行出错信息记录
		loggerPath := g.Cfg().MustGet(context.Background(), "logger.path").String()
		err := debug.SetCrashOutput(utility.InitFatalLog(loggerPath), debug.CrashOptions{})
		if err != nil {
			g.Log().Error(ctx, err)
		}

		// https
		https := g.Cfg().MustGet(ctx, "server.https").Bool()
		if https {
			certFile := g.Cfg().MustGet(ctx, "server.httpsCertFile").String()
			keyFile := g.Cfg().MustGet(ctx, "server.httpsKeyFile").String()
			s.EnableHTTPS(certFile, keyFile)
		}

		s.Run()
		stopSignal <- syscall.SIGQUIT
	}()

	// 在程序退出信号处理中添加协程清理
	// 监听系统信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		g.Log().Infof(ctx, "接收到信号 %v，开始优雅关闭...", sig)

		// 优雅关闭各个组件，防止协程泄漏
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// 4. 关闭HTTP服务器
		if s != nil {
			g.Log().Info(shutdownCtx, "关闭HTTP服务器...")
			if err := s.Shutdown(); err != nil {
				g.Log().Errorf(shutdownCtx, "HTTP服务器关闭失败: %v", err)
			}
		}

		// 5. 给协程一些时间完成清理
		time.Sleep(2 * time.Second)

		g.Log().Info(shutdownCtx, "优雅关闭完成")
		os.Exit(0)
	}()
	return
}
