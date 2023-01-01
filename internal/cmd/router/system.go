package router

import (
	"context"
	alarmController "github.com/sagoo-cloud/sagooiot/internal/controller/alarm"
	commonController "github.com/sagoo-cloud/sagooiot/internal/controller/common"
	networkController "github.com/sagoo-cloud/sagooiot/internal/controller/network"
	noticeController "github.com/sagoo-cloud/sagooiot/internal/controller/notice"
	productController "github.com/sagoo-cloud/sagooiot/internal/controller/product"
	sourceController "github.com/sagoo-cloud/sagooiot/internal/controller/source"
	systemController "github.com/sagoo-cloud/sagooiot/internal/controller/system"
	tdengineController "github.com/sagoo-cloud/sagooiot/internal/controller/tdengine"

	"github.com/sagoo-cloud/sagooiot/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

// System 系统默认功能的路由，不含业务属性的
func System(ctx context.Context, group *ghttp.RouterGroup) {
	//系统登录路由
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Bind(
			systemController.Login,   // 登录
			systemController.Captcha, // 验证码
			commonController.SysInfo, //系统信息

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
			commonController.BaseDbLink, //数据源管理
			commonController.CityData,   //城市管理
		)
	})

	// 产品设备相关路由
	group.Group("/product", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			productController.Category,    // 产品分类
			productController.Product,     // 产品
			productController.Device,      // 设备
			productController.DeviceTag,   // 设备标签
			productController.DeviceLog,   // 设备日志
			productController.TSLDataType, // 物模型：数据类型
			productController.TSLProperty, // 物模型：属性
			productController.TSLFunction, // 物模型：功能
			productController.TSLEvent,    // 物模型：事件
			productController.TSLTag,      // 物模型：标签
		)
	})

	// 告警相关路由
	group.Group("/alarm", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			alarmController.AlarmLevel, // 告警级别
			alarmController.AlarmRule,  // 告警规则
			alarmController.AlarmLog,   // 告警日志
		)
	})

	// 数据源相关路由
	group.Group("/source", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			sourceController.DataSource,       // 数据源
			sourceController.DataNode,         // 数据节点
			sourceController.DataTemplate,     // 数据模型
			sourceController.DataTemplateNode, // 数据模型节点
		)
	})

	// 网络通道相关路由
	group.Group("/network", func(group *ghttp.RouterGroup) {
		//group.Middleware(service.Middleware().Auth)
		group.Bind(
			networkController.Tunnel, // 通讯通道管理
			networkController.Server, // 通讯服务管理

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
			systemController.SysApi,          // 接口
			systemController.SysAuthorize,    //权限管理
			systemController.SysOrganization, //组织管理
			systemController.SysOperLog,      //操作日志管理
			systemController.SysLoginLog,     //访问日志管理

			systemController.SysJob, //定时任务管理

			systemController.SysMonitor,    //服务监控
			systemController.SysUserOnline, //在线用户

			systemController.SysNotifications, //消息中心
			systemController.SysPlugins,       //插件管理
			systemController.SysPluginsConfig, //插件配置管理

		)
	})

	//时序数据库相关路由
	group.Group("/tdengine", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			tdengineController.TdEngine, //websocket
		)
	})

	//通知服务相关路由
	group.Group("/notice", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			noticeController.NoticeInfo,
			noticeController.NoticeConfig,
			noticeController.NoticeTemplate,
			noticeController.NoticeLog,
		)
	})

}
