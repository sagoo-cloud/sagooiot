// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenuApi is the golang structure for table sys_menu_api.
type SysMenuApi struct {
	Id        uint        `json:"id"        description:"id"`
	MenuId    int         `json:"menuId"    description:"菜单ID"`
	ApiId     int         `json:"apiId"     description:"apiId"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedBy uint        `json:"createdBy" description:"创建人"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
