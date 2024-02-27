// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenuApi is the golang structure of table sys_menu_api for DAO operations like Where/Data.
type SysMenuApi struct {
	g.Meta    `orm:"table:sys_menu_api, do:true"`
	Id        interface{} // id
	MenuId    interface{} // 菜单ID
	ApiId     interface{} // apiId
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedBy interface{} // 创建人
	CreatedAt *gtime.Time // 创建时间
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}
