// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysNotifications is the golang structure of table sys_notifications for DAO operations like Where/Data.
type SysNotifications struct {
	g.Meta    `orm:"table:sys_notifications, do:true"`
	Id        interface{} //
	Title     interface{} // 标题
	Doc       interface{} // 描述
	Source    interface{} // 消息来源
	Types     interface{} // 类型
	CreatedAt *gtime.Time // 发送时间
	Status    interface{} // 0，未读，1，已读
}
