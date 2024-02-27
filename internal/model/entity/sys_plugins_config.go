// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SysPluginsConfig is the golang structure for table sys_plugins_config.
type SysPluginsConfig struct {
	Id    int    `json:"id"    description:""`
	Type  string `json:"type"  description:"插件类型"`
	Name  string `json:"name"  description:"插件名称"`
	Value string `json:"value" description:"配置内容"`
	Doc   string `json:"doc"   description:"配置说明"`
}
