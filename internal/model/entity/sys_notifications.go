// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNotifications is the golang structure for table sys_notifications.
type SysNotifications struct {
	Id        int         `json:"id"        description:""`
	Title     string      `json:"title"     description:"标题"`
	Doc       string      `json:"doc"       description:"描述"`
	Source    string      `json:"source"    description:"消息来源"`
	Types     string      `json:"types"     description:"类型"`
	CreatedAt *gtime.Time `json:"createdAt" description:"发送时间"`
	Status    int         `json:"status"    description:"0，未读，1，已读"`
}
