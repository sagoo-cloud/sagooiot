// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPost is the golang structure of table sys_post for DAO operations like Where/Data.
type SysPost struct {
	g.Meta    `orm:"table:sys_post, do:true"`
	PostId    interface{} // 岗位ID
	DeptId    interface{} // 部门ID
	ParentId  interface{} // 父ID
	PostCode  interface{} // 岗位编码
	PostName  interface{} // 岗位名称
	PostSort  interface{} // 显示顺序
	Status    interface{} // 状态（0正常 1停用）
	Remark    interface{} // 备注
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedBy interface{} // 创建人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy interface{} // 修改人
	UpdatedAt *gtime.Time // 修改时间
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}
