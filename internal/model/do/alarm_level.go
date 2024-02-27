// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// AlarmLevel is the golang structure of table alarm_level for DAO operations like Where/Data.
type AlarmLevel struct {
	g.Meta `orm:"table:alarm_level, do:true"`
	Level  interface{} // 告警级别
	Name   interface{} // 名称
}
