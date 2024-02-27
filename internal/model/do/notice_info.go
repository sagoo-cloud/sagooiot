// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeInfo is the golang structure of table notice_info for DAO operations like Where/Data.
type NoticeInfo struct {
	g.Meta     `orm:"table:notice_info, do:true"`
	Id         interface{} //
	ConfigId   interface{} //
	ComeFrom   interface{} //
	Method     interface{} //
	MsgTitle   interface{} //
	MsgBody    interface{} //
	MsgUrl     interface{} //
	UserIds    interface{} //
	OrgIds     interface{} //
	Totag      interface{} //
	Status     interface{} //
	MethodCron interface{} //
	MethodNum  interface{} //
	CreatedAt  *gtime.Time //
}
