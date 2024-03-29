// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure for table sys_post.
type SysPost struct {
	PostId    uint64      `json:"postId"    description:"岗位ID"`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	ParentId  int         `json:"parentId"  description:"父ID"`
	PostCode  string      `json:"postCode"  description:"岗位编码"`
	PostName  string      `json:"postName"  description:"岗位名称"`
	PostSort  int         `json:"postSort"  description:"显示顺序"`
	Status    uint        `json:"status"    description:"状态（0正常 1停用）"`
	Remark    string      `json:"remark"    description:"备注"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedBy uint        `json:"createdBy" description:"创建人"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedBy uint        `json:"updatedBy" description:"修改人"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
