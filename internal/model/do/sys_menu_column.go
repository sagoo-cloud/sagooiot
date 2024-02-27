// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenuColumn is the golang structure of table sys_menu_column for DAO operations like Where/Data.
type SysMenuColumn struct {
	g.Meta      `orm:"table:sys_menu_column, do:true"`
	Id          interface{} //
	ParentId    interface{} // 父ID
	MenuId      interface{} // 菜单ID
	Name        interface{} // 名称
	Code        interface{} // 代表字段
	Description interface{} // 描述
	Status      interface{} // 状态 0 停用 1启用
	IsDeleted   interface{} // 是否删除 0未删除 1已删除
	CreatedBy   interface{} // 创建人
	CreatedAt   *gtime.Time // 创建时间
	UpdatedBy   interface{} // 修改人
	UpdatedAt   *gtime.Time // 更新时间
	DeletedBy   interface{} // 删除人
	DeletedAt   *gtime.Time // 删除时间
}
