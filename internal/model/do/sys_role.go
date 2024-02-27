// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table sys_role for DAO operations like Where/Data.
type SysRole struct {
	g.Meta    `orm:"table:sys_role, do:true"`
	Id        interface{} //
	DeptId    interface{} // 部门ID
	ParentId  interface{} // 父ID
	ListOrder interface{} // 排序
	Name      interface{} // 角色名称
	DataScope interface{} // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Remark    interface{} // 备注
	Status    interface{} // 状态;0:禁用;1:正常
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建日期
	UpdatedBy interface{} // 更新者
	UpdatedAt *gtime.Time // 修改日期
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}
