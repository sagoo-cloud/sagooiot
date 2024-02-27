package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	commonController "sagooiot/internal/controller/common"
	systemController "sagooiot/internal/controller/system"
	"sagooiot/internal/service"
)

// System 系统默认功能的路由，不含业务属性的
func System(ctx context.Context, group *ghttp.RouterGroup) {
	//系统登录路由
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			systemController.Login,     // 登录
			systemController.Captcha,   // 验证码
			commonController.SysInfo,   //系统信息
			commonController.CheckAuth, //权限验证
		)
	})

	// 公共接口相关路由
	group.Group("/common", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			commonController.Upload,
			commonController.ConfigData,
			commonController.DictType,
			commonController.DictData,
		)
	})

	//系统权限控制路由
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			systemController.SysRole,         // 角色
			systemController.SysDept,         // 部门
			systemController.SysPost,         // 岗位
			systemController.SysUser,         // 用户
			systemController.SysMenu,         // 菜单
			systemController.SysMenuButton,   // 菜单按钮
			systemController.SysMenuColumn,   // 菜单列表
			systemController.SysMenuApi,      // 菜单API
			systemController.SysApi,          // 接口
			systemController.SysAuthorize,    //权限管理
			systemController.SysOrganization, //组织管理
			systemController.SysOperLog,      //操作日志管理
			systemController.SysLoginLog,     //访问日志管理

			systemController.SysJob, //定时任务管理

			systemController.SysUserOnline, //在线用户

			systemController.SysNotifications, //消息中心
			systemController.SysPlugins,       //插件管理
			systemController.SysPluginsConfig, //插件配置管理

			systemController.SysMessage,     // 通知中心
			systemController.SysCertificate, // 证书管理

		)
	})

}
