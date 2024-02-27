// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrganization is the golang structure for table sys_organization.
type SysOrganization struct {
	Id        int64       `json:"id"        description:"组织ID"`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	ParentId  int64       `json:"parentId"  description:"父组织id"`
	Ancestors string      `json:"ancestors" description:"祖级列表"`
	Name      string      `json:"name"      description:"组织名称"`
	Number    string      `json:"number"    description:"组织编号"`
	OrderNum  int         `json:"orderNum"  description:"显示顺序"`
	Leader    string      `json:"leader"    description:"负责人"`
	Phone     string      `json:"phone"     description:"联系电话"`
	Email     string      `json:"email"     description:"邮箱"`
	Status    uint        `json:"status"    description:"组织状态（0停用 1正常）"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	CreatedBy uint        `json:"createdBy" description:"创建人"`
	UpdatedBy int         `json:"updatedBy" description:"修改人"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
