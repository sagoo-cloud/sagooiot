// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPlugins is the golang structure of table sys_plugins for DAO operations like Where/Data.
type SysPlugins struct {
	g.Meta                `orm:"table:sys_plugins, do:true"`
	Id                    interface{} // ID
	DeptId                interface{} // 部门ID
	Types                 interface{} // 插件与SagooIOT的通信方式
	HandleType            interface{} // 功能类型
	Name                  interface{} // 名称
	Title                 interface{} // 标题
	Description           interface{} // 介绍
	Version               interface{} // 版本
	Author                interface{} // 作者
	Icon                  interface{} // 插件图标
	Link                  interface{} // 插件的网址。指向插件的 github 链接。值应为一个可访问的网址
	Command               interface{} // 插件的运行指令
	Args                  interface{} // 插件的指令参数
	Status                interface{} // 状态  0未启用  1启用
	FrontendUi            interface{} // 是否有插件页面
	FrontendUrl           interface{} // 插件页面地址
	FrontendConfiguration interface{} // 是否显示配置页面
	StartTime             *gtime.Time // 启动时间
	IsDeleted             interface{} // 是否删除 0未删除 1已删除
	CreatedBy             interface{} // 创建者
	CreatedAt             *gtime.Time // 创建日期
	UpdatedBy             interface{} // 修改人
	UpdatedAt             *gtime.Time // 更新时间
	DeletedBy             interface{} // 删除人
	DeletedAt             *gtime.Time // 删除时间
}
