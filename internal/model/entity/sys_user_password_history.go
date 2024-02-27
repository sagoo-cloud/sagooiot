// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserPasswordHistory is the golang structure for table sys_user_password_history.
type SysUserPasswordHistory struct {
	Id             int         `json:"id"             description:""`
	UserId         int         `json:"userId"         description:"用户ID"`
	BeforePassword string      `json:"beforePassword" description:"变更之前密码"`
	AfterPassword  string      `json:"afterPassword"  description:"变更之后密码"`
	ChangeTime     *gtime.Time `json:"changeTime"     description:"变更时间"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:""`
	CreatedBy      int         `json:"createdBy"      description:""`
}
