// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMessage is the golang structure of table sys_message for DAO operations like Where/Data.
type SysMessage struct {
	g.Meta    `orm:"table:sys_message, do:true"`
	Id        interface{} //
	Title     interface{} // 标题
	Types     interface{} // 字典表
	Scope     interface{} // 消息范围
	Content   interface{} // 内容
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建日期
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}
