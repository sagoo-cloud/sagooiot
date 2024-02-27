// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// SysPluginsConfig is the golang structure of table sys_plugins_config for DAO operations like Where/Data.
type SysPluginsConfig struct {
	g.Meta `orm:"table:sys_plugins_config, do:true"`
	Id     interface{} //
	Type   interface{} // 插件类型
	Name   interface{} // 插件名称
	Value  interface{} // 配置内容
	Doc    interface{} // 配置说明
}
