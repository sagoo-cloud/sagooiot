// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMessagereceive is the golang structure of table sys_messagereceive for DAO operations like Where/Data.
type SysMessagereceive struct {
	g.Meta    `orm:"table:sys_messagereceive, do:true"`
	Id        interface{} //
	UserId    interface{} // 用户ID
	MessageId interface{} // 消息ID
	IsRead    interface{} // 是否已读 0 未读 1已读
	IsPush    interface{} // 是否已经推送0 否 1是
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	ReadTime  *gtime.Time // 阅读时间
	DeletedAt *gtime.Time // 删除时间
}
