// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeTemplate is the golang structure of table notice_template for DAO operations like Where/Data.
type NoticeTemplate struct {
	g.Meta      `orm:"table:notice_template, do:true"`
	Id          interface{} //
	DeptId      interface{} // 部门ID
	ConfigId    interface{} //
	SendGateway interface{} //
	Code        interface{} //
	Title       interface{} //
	Content     interface{} //
	CreatedAt   *gtime.Time //
}
