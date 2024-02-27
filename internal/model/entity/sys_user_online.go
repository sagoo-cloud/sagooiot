// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserOnline is the golang structure for table sys_user_online.
type SysUserOnline struct {
	Id        int         `json:"id"        description:""`
	Uuid      string      `json:"uuid"      description:"用户标识"`
	Key       string      `json:"key"       description:""`
	Token     string      `json:"token"     description:"用户token"`
	CreatedAt *gtime.Time `json:"createdAt" description:"登录时间"`
	UserName  string      `json:"userName"  description:"用户名"`
	Ip        string      `json:"ip"        description:"登录ip"`
	Explorer  string      `json:"explorer"  description:"浏览器"`
	Os        string      `json:"os"        description:"操作系统"`
}
