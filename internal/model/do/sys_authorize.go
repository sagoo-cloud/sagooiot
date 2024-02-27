// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuthorize is the golang structure of table sys_authorize for DAO operations like Where/Data.
type SysAuthorize struct {
	g.Meta     `orm:"table:sys_authorize, do:true"`
	Id         interface{} //
	RoleId     interface{} // 角色ID
	ItemsType  interface{} // 项目类型 menu菜单 button按钮 column列表字段 api接口
	ItemsId    interface{} // 项目ID
	IsCheckAll interface{} // 是否全选 1是 0否
	IsDeleted  interface{} // 是否删除 0未删除 1已删除
	CreatedBy  interface{} // 创建人
	CreatedAt  *gtime.Time // 创建时间
	DeletedBy  interface{} // 删除人
	DeletedAt  *gtime.Time // 删除时间
}
