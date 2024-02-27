// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrganization is the golang structure of table sys_organization for DAO operations like Where/Data.
type SysOrganization struct {
	g.Meta    `orm:"table:sys_organization, do:true"`
	Id        interface{} // 组织ID
	DeptId    interface{} // 部门ID
	ParentId  interface{} // 父组织id
	Ancestors interface{} // 祖级列表
	Name      interface{} // 组织名称
	Number    interface{} // 组织编号
	OrderNum  interface{} // 显示顺序
	Leader    interface{} // 负责人
	Phone     interface{} // 联系电话
	Email     interface{} // 邮箱
	Status    interface{} // 组织状态（0停用 1正常）
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedAt *gtime.Time // 创建时间
	CreatedBy interface{} // 创建人
	UpdatedBy interface{} // 修改人
	UpdatedAt *gtime.Time // 修改时间
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}
