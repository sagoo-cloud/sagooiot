// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserPasswordHistory is the golang structure of table sys_user_password_history for DAO operations like Where/Data.
type SysUserPasswordHistory struct {
	g.Meta         `orm:"table:sys_user_password_history, do:true"`
	Id             interface{} //
	UserId         interface{} // 用户ID
	BeforePassword interface{} // 变更之前密码
	AfterPassword  interface{} // 变更之后密码
	ChangeTime     *gtime.Time // 变更时间
	CreatedAt      *gtime.Time //
	CreatedBy      interface{} //
}
