// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NoticeConfig is the golang structure of table notice_config for DAO operations like Where/Data.
type NoticeConfig struct {
	g.Meta      `orm:"table:notice_config, do:true"`
	Id          interface{} //
	DeptId      interface{} // 部门ID
	Title       interface{} //
	SendGateway interface{} //
	Types       interface{} //
	CreatedAt   *gtime.Time //
}
