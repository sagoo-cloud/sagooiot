package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

// GetSysPluginsListReq 获取数据列表
type GetSysPluginsListReq struct {
	g.Meta `path:"/plugins/list" method:"get" summary:"获取插件列表" tags:"插件管理"`
	common.PaginationReq
}
type GetSysPluginsListRes struct {
	Data []model.GetSysPluginsListRes
	common.PaginationRes
}

// GetSysPluginsByIdReq 获取指定ID的数据
type GetSysPluginsByIdReq struct {
	g.Meta `path:"/plugins/get" method:"get" summary:"获取插件" tags:"插件管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetSysPluginsByIdRes struct {
	*model.SysPluginsRes
}

// AddSysPluginsReq 添加插件
type AddSysPluginsReq struct {
	g.Meta `path:"/plugins/add" method:"post" summary:"添加插件" tags:"插件管理"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"上传文件" v:"required#请上传文件"`
}
type AddSysPluginsRes struct {
}

type EditSysPluginsStatusReq struct {
	g.Meta `path:"/plugins/set" method:"post" summary:"设置插件状态" tags:"插件管理"`
	Id     int `json:"id"          description:"ID"`
	Status int `json:"status"          description:"状态，0停用，1启用"`
}
type EditSysPluginsStatusRes struct{}

type EditSysPluginsReq struct {
	g.Meta                `path:"/plugins/edit" method:"put" summary:"修改插件信息" tags:"插件管理"`
	Id                    int      `json:"id"                    description:"ID" v:"required#id不能为空"`
	Types                 string   `json:"types"                 description:"插件与SagooIOT的通信方式" v:"required#通信方式不能为空"`
	HandleType            string   `json:"handleType"            description:"功能类型" v:"required#功能类型不能为空"`
	Name                  string   `json:"name"                  description:"名称" v:"required#名称不能为空"`
	Title                 string   `json:"title"                 description:"标题" v:"required#标题不能为空"`
	Description           string   `json:"description"           description:"介绍"`
	Version               string   `json:"version"               description:"版本"`
	Author                []string `json:"author"                description:"作者"`
	Icon                  string   `json:"icon"                  description:"插件图标"`
	Link                  string   `json:"link"                  description:"插件的网址。指向插件的 github 链接。值应为一个可访问的网址"`
	Command               string   `json:"command"               description:"插件的运行指令"`
	Args                  []string `json:"args"                  description:"插件的指令参数"`
	FrontendUi            int      `json:"frontendUi"            description:"是否有插件页面"`
	FrontendUrl           string   `json:"frontendUrl"           description:"插件页面地址"`
	FrontendConfiguration int      `json:"frontendConfiguration" description:"是否显示配置页面"`
}
type EditSysPluginsRes struct {
}

type DelSysPluginsStatusReq struct {
	g.Meta `path:"/plugins/del" method:"delete" summary:"删除插件" tags:"插件管理"`
	Ids    []int `json:"ids"          description:"ID"`
}
type DelSysPluginsStatusRes struct{}

type GetSysPluginsTypesAllReq struct {
	g.Meta `path:"/plugins/getTypesAll" method:"get" summary:"获取插件类型" tags:"插件管理"`
	Types  string `json:"types"            description:"功能类型" v:"required#插件类型不能为空 协议(protocol)或者通知(notice)"`
}
type GetSysPluginsTypesAllRes struct {
	Data []*model.SysPluginsInfoRes
}
