// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// ResponseHandler 返回处理中间件
		ResponseHandler(r *ghttp.Request)
		// Ctx 自定义上下文对象
		Ctx(r *ghttp.Request)
		// Auth 前台系统权限控制，用户必须登录才能访问
		Auth(r *ghttp.Request)
		// MiddlewareCORS 跨域处理
		MiddlewareCORS(r *ghttp.Request)
		// OperationLog 操作日志
		OperationLog(r *ghttp.Request)
		Tracing(r *ghttp.Request)
		I18n(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
