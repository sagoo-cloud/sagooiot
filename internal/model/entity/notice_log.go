// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeLog is the golang structure for table notice_log.
type NoticeLog struct {
	Id          uint64      `json:"id"          description:""`
	DeptId      int         `json:"deptId"      description:"部门ID"`
	SendGateway string      `json:"sendGateway" description:"通知渠道"`
	TemplateId  string      `json:"templateId"  description:"通知模板ID"`
	Addressee   string      `json:"addressee"   description:"收信人列表"`
	Title       string      `json:"title"       description:"通知标题"`
	Content     string      `json:"content"     description:"通知内容"`
	Status      int         `json:"status"      description:"发送状态：0=失败，1=成功"`
	FailMsg     string      `json:"failMsg"     description:"失败信息"`
	SendTime    *gtime.Time `json:"sendTime"    description:"发送时间"`
}
