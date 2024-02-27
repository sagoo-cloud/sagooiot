package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	alarmController "sagooiot/internal/controller/alarm"
	networkController "sagooiot/internal/controller/network"
	noticeController "sagooiot/internal/controller/notice"
	productController "sagooiot/internal/controller/product"
	tdengineController "sagooiot/internal/controller/tdengine"

	"sagooiot/internal/service"
)

// Iot iot功能的路由
func Iot(ctx context.Context, group *ghttp.RouterGroup) {

	// 产品设备相关路由
	group.Group("/product", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			productController.Category,       // 产品分类
			productController.Product,        // 产品
			productController.Device,         // 设备
			productController.DeviceTag,      // 设备标签
			productController.DeviceLog,      // 设备日志
			productController.DeviceFunction, // 设备功能执行
			productController.DeviceProperty, // 设备属性设置
			productController.TSLDataType,    // 物模型：数据类型
			productController.TSLProperty,    // 物模型：属性
			productController.TSLFunction,    // 物模型：功能
			productController.TSLEvent,       // 物模型：事件
			productController.TSLTag,         // 物模型：标签
			productController.DeviceTree,     // 设备树
			productController.TSLImport,      // 物模型：导入/导出
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

	// 网络通道相关路由
	group.Group("/network", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Auth)
		group.Bind(
			networkController.Tunnel, // 通讯通道管理
			networkController.Server, // 通讯服务管理

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
