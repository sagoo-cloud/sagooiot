// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMessage is the golang structure for table sys_message.
type SysMessage struct {
	Id        int         `json:"id"        description:""`
	Title     string      `json:"title"     description:"标题"`
	Types     int         `json:"types"     description:"字典表"`
	Scope     int         `json:"scope"     description:"消息范围"`
	Content   string      `json:"content"   description:"内容"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedBy uint        `json:"createdBy" description:"创建者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建日期"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
