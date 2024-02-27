// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id        uint        `json:"id"        description:""`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	ParentId  int         `json:"parentId"  description:"父ID"`
	ListOrder uint        `json:"listOrder" description:"排序"`
	Name      string      `json:"name"      description:"角色名称"`
	DataScope uint        `json:"dataScope" description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    uint        `json:"status"    description:"状态;0:禁用;1:正常"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedBy uint        `json:"createdBy" description:"创建者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建日期"`
	UpdatedBy uint        `json:"updatedBy" description:"更新者"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改日期"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
