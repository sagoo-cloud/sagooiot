package cmd

import (
	"context"
	"expvar"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gmode"
	"os"
	"sagooiot/internal/cmd/router"
	"sagooiot/internal/service"
	"sagooiot/internal/sse"
	"sagooiot/module"
	"syscall"
)

func RunServer(ctx context.Context, stopSignal chan os.Signal) {
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

	s.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(
			service.Middleware().Ctx,
			service.Middleware().ResponseHandler,
			service.Middleware().MiddlewareCORS,
			service.Middleware().I18n,
		)
		service.SysToken().GfToken().Middleware(group)
		router.System(ctx, group)   //系统默认功能的路由
		router.Iot(ctx, group)      //Iot功能的路由
		router.Analysis(ctx, group) //分析统计功能的路由
		module.Router(ctx, group)   //加载模块的路由
		router.OAuth(ctx, group)    //第三方授权登录

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

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("panic 产生，错误:", err)
			}
		}()
		// https
		https := g.Cfg().MustGet(ctx, "server.https").Bool()
		if https {
			certFile := g.Cfg().MustGet(ctx, "server.httpsCertFile").String()
			keyFile := g.Cfg().MustGet(ctx, "server.httpsKeyFile").String()
			s.EnableHTTPS(certFile, keyFile)
		}

		go s.Run()
		select {
		case <-ctx.Done():
		}
		stopSignal <- syscall.SIGQUIT
	}()
	return
}
