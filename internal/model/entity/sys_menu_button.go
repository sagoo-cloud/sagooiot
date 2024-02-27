// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenuButton is the golang structure for table sys_menu_button.
type SysMenuButton struct {
	Id          uint        `json:"id"          description:""`
	ParentId    int         `json:"parentId"    description:"父ID"`
	MenuId      int         `json:"menuId"      description:"菜单ID"`
	Name        string      `json:"name"        description:"名称"`
	Types       string      `json:"types"       description:"类型 自定义 add添加 edit编辑 del 删除"`
	Description string      `json:"description" description:"描述"`
	Status      int         `json:"status"      description:"状态 0 停用 1启用"`
	IsDeleted   int         `json:"isDeleted"   description:"是否删除 0未删除 1已删除"`
	CreatedBy   uint        `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedBy   int         `json:"updatedBy"   description:"修改人"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
	DeletedBy   int         `json:"deletedBy"   description:"删除人"`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:"删除时间"`
}
