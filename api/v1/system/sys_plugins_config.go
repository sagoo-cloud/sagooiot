package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
)

// GetPluginsConfigListReq 获取数据列表
type GetPluginsConfigListReq struct {
	g.Meta `path:"/plugins_config/list" method:"get" summary:"获取插件配置列表" tags:"插件管理"`
	common.PaginationReq
}
type GetPluginsConfigListRes struct {
	Data []GetPluginsConfigByIdRes
	common.PaginationRes
}

// GetPluginsConfigByIdReq 获取指定ID的数据
type GetPluginsConfigByIdReq struct {
	g.Meta `path:"/plugins_config/get" method:"get" summary:"获取插件配置" tags:"插件管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetPluginsConfigByIdRes struct {
	Value string `json:"value"          description:"配置内容"`
	Doc   string `json:"doc"          description:"配置说明"`
	Id    string `json:"id"          description:""`
	Type  string `json:"type"          description:"插件类型"`
	Name  string `json:"name"          description:"插件名称"`
}

// GetPluginsConfigByNameReq 获取指定类型与名称的的插件数据
type GetPluginsConfigByNameReq struct {
	g.Meta `path:"/plugins_config/getbyname" method:"get" summary:"获取指定类型与名称的的插件数据" tags:"插件管理"`
	Type   string `json:"type"         description:"type" v:"required#类型不能为空"`
	Name   string `json:"name"         description:"name" v:"required#名称不能为空"`
}
type GetPluginsConfigByNameRes struct {
	Value string `json:"value"          description:"配置内容"`
	Doc   string `json:"doc"          description:"配置说明"`
	Id    string `json:"id"          description:""`
	Type  string `json:"type"          description:"插件类型"`
	Name  string `json:"name"          description:"插件名称"`
}

// AddPluginsConfigReq 添加数据
type AddPluginsConfigReq struct {
	g.Meta `path:"/plugins_config/add" method:"post" summary:"添加插件配置" tags:"插件管理"`
	Type   string `json:"type"          description:"插件类型"`
	Name   string `json:"name"          description:"插件名称"`
	Value  string `json:"value"          description:"配置内容"`
	Doc    string `json:"doc"          description:"配置说明"`
}
type AddPluginsConfigRes struct{}

// SavePluginsConfigReq 直接更新配置数据
type SavePluginsConfigReq struct {
	g.Meta `path:"/plugins_config/save" method:"post" summary:"直接更新配置数据" tags:"插件管理"`
	Type   string `json:"type"           description:"type" v:"required#插件类型不能为空"`
	Name   string `json:"name"           description:"name" v:"required#插件名称不能为空"`
	Value  string `json:"value"          description:"配置内容，一行一条"`
	Doc    string `json:"doc"          description:"配置说明"`
}
type SavePluginsConfigRes struct{}

// EditPluginsConfigReq 编辑数据api
type EditPluginsConfigReq struct {
	g.Meta `path:"/plugins_config/edit" method:"put" summary:"编辑插件配置" tags:"插件管理"`
	Value  string `json:"value"          description:"配置内容"`
	Doc    string `json:"doc"          description:"配置说明"`
	Id     string `json:"id"          description:""`
	Type   string `json:"type"          description:"插件类型"`
	Name   string `json:"name"          description:"插件名称"`
}
type EditPluginsConfigRes struct{}

// DeletePluginsConfigReq 删除数据
type DeletePluginsConfigReq struct {
	g.Meta `path:"/plugins_config/delete" method:"delete" summary:"删除插件配置" tags:"插件管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeletePluginsConfigRes struct{}
