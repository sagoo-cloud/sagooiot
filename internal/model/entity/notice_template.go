// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeTemplate is the golang structure for table notice_template.
type NoticeTemplate struct {
	Id          string      `json:"id"          description:""`
	DeptId      int         `json:"deptId"      description:"部门ID"`
	ConfigId    string      `json:"configId"    description:""`
	SendGateway string      `json:"sendGateway" description:""`
	Code        string      `json:"code"        description:""`
	Title       string      `json:"title"       description:""`
	Content     string      `json:"content"     description:""`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
}
