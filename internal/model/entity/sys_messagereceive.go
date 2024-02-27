// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMessagereceive is the golang structure for table sys_messagereceive.
type SysMessagereceive struct {
	Id        int         `json:"id"        description:""`
	UserId    int         `json:"userId"    description:"用户ID"`
	MessageId int         `json:"messageId" description:"消息ID"`
	IsRead    int         `json:"isRead"    description:"是否已读 0 未读 1已读"`
	IsPush    int         `json:"isPush"    description:"是否已经推送0 否 1是"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	ReadTime  *gtime.Time `json:"readTime"  description:"阅读时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
