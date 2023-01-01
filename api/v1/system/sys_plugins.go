package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/api/v1/common"
)

//GetSysPluginsListReq 获取数据列表
type GetSysPluginsListReq struct {
	g.Meta `path:"/plugins/list" method:"get" summary:"获取插件列表" tags:"插件管理"`
	common.PaginationReq
}
type GetSysPluginsListRes struct {
	Data []GetSysPluginsByIdRes
	common.PaginationRes
}

//GetSysPluginsByIdReq 获取指定ID的数据
type GetSysPluginsByIdReq struct {
	g.Meta `path:"/plugins/get" method:"get" summary:"获取插件" tags:"插件管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetSysPluginsByIdRes struct {
	Version   string `json:"version"          description:"版本"`
	Author    string `json:"author"          description:""`
	StartTime string `json:"startTime"          description:""`
	Id        int    `json:"id"          description:"ID"`
	Name      string `json:"name"          description:"名称"`
	Title     string `json:"title"          description:"标题"`
	Intro     string `json:"intro"          description:"介绍"`
	Status    int    `json:"status"          description:"状态"`
	Types     string `json:"types"          description:"插件类型"`
}

type EditSysPluginsStatusReq struct {
	g.Meta `path:"/plugins/set" method:"post" summary:"设置插件状态" tags:"插件管理"`
	Id     int `json:"id"          description:"ID"`
	Status int `json:"status"          description:"状态，0停用，1启用"`
}
type EditSysPluginsStatusRes struct{}
