// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeInfo is the golang structure for table notice_info.
type NoticeInfo struct {
	Id         int64       `json:"id"         description:""`
	ConfigId   string      `json:"configId"   description:""`
	ComeFrom   string      `json:"comeFrom"   description:""`
	Method     string      `json:"method"     description:""`
	MsgTitle   string      `json:"msgTitle"   description:""`
	MsgBody    string      `json:"msgBody"    description:""`
	MsgUrl     string      `json:"msgUrl"     description:""`
	UserIds    string      `json:"userIds"    description:""`
	OrgIds     string      `json:"orgIds"     description:""`
	Totag      string      `json:"totag"      description:""`
	Status     int         `json:"status"     description:""`
	MethodCron string      `json:"methodCron" description:""`
	MethodNum  int         `json:"methodNum"  description:""`
	CreatedAt  *gtime.Time `json:"createdAt"  description:""`
}
