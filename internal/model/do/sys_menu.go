// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure of table sys_menu for DAO operations like Where/Data.
type SysMenu struct {
	g.Meta     `orm:"table:sys_menu, do:true"`
	Id         interface{} //
	ParentId   interface{} // 父ID
	Name       interface{} // 规则名称
	Title      interface{} // 菜单名称
	Icon       interface{} // 图标
	Condition  interface{} // 条件
	Remark     interface{} // 备注
	MenuType   interface{} // 类型 0目录 1菜单 2按钮
	Weigh      interface{} // 权重
	IsHide     interface{} // 显示状态
	Path       interface{} // 路由地址
	Component  interface{} // 组件路径
	IsLink     interface{} // 是否外链 1是 0否
	ModuleType interface{} // 所属模块 system 运维 company企业
	ModelId    interface{} // 模型ID
	IsIframe   interface{} // 是否内嵌iframe
	IsCached   interface{} // 是否缓存
	Redirect   interface{} // 路由重定向地址
	IsAffix    interface{} // 是否固定
	LinkUrl    interface{} // 链接地址
	Status     interface{} // 状态 0 停用 1启用
	IsDeleted  interface{} // 是否删除 0未删除 1已删除
	CreatedBy  interface{} // 创建人
	CreatedAt  *gtime.Time // 创建时间
	UpdatedBy  interface{} // 修改人
	UpdatedAt  *gtime.Time // 更新时间
	DeletedBy  interface{} // 删除人
	DeletedAt  *gtime.Time // 删除时间
}
