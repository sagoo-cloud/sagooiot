// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeLog is the golang structure of table notice_log for DAO operations like Where/Data.
type NoticeLog struct {
	g.Meta      `orm:"table:notice_log, do:true"`
	Id          interface{} //
	DeptId      interface{} // 部门ID
	SendGateway interface{} // 通知渠道
	TemplateId  interface{} // 通知模板ID
	Addressee   interface{} // 收信人列表
	Title       interface{} // 通知标题
	Content     interface{} // 通知内容
	Status      interface{} // 发送状态：0=失败，1=成功
	FailMsg     interface{} // 失败信息
	SendTime    *gtime.Time // 发送时间
}
