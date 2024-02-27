package model

import "github.com/gogf/gf/v2/os/gtime"

type GetSysPluginsListRes struct {
	Id                    int         `json:"id"                    description:"ID"`
	Types                 string      `json:"types"                 description:"插件与SagooIOT的通信方式"`
	HandleType            string      `json:"handleType"            description:"功能类型"`
	Name                  string      `json:"name"                  description:"名称"`
	Title                 string      `json:"title"                 description:"标题"`
	Description           string      `json:"description"           description:"介绍"`
	Version               string      `json:"version"               description:"版本"`
	Author                string      `json:"author"                description:"作者"`
	Icon                  string      `json:"icon"                  description:"插件图标"`
	Link                  string      `json:"link"                  description:"插件的网址。指向插件的 github 链接。值应为一个可访问的网址"`
	Command               string      `json:"command"               description:"插件的运行指令"`
	Args                  string      `json:"args"                  description:"插件的指令参数"`
	Status                int         `json:"status"                description:"状态"`
	FrontendUi            int         `json:"frontendUi"            description:"是否有插件页面"`
	FrontendUrl           string      `json:"frontendUrl"           description:"插件页面地址"`
	FrontendConfiguration int         `json:"frontendConfiguration" description:"是否显示配置页面"`
	StartTime             *gtime.Time `json:"startTime"             description:"启动时间"`
	IsDeleted             int         `json:"isDeleted"             description:"是否删除 0未删除 1已删除"`
	CreatedBy             uint        `json:"createdBy"             description:"创建者"`
	CreatedAt             *gtime.Time `json:"createdAt"             description:"创建日期"`
	UpdatedBy             int         `json:"updatedBy"             description:"修改人"`
	UpdatedAt             *gtime.Time `json:"updatedAt"             description:"更新时间"`
	DeletedBy             int         `json:"deletedBy"             description:"删除人"`
	DeletedAt             *gtime.Time `json:"deletedAt"             description:"删除时间"`
}

type GetSysPluginsListOut struct {
	Id                    int         `json:"id"                    description:"ID"`
	Types                 string      `json:"types"                 description:"插件与SagooIOT的通信方式"`
	HandleType            string      `json:"handleType"            description:"功能类型"`
	Name                  string      `json:"name"                  description:"名称"`
	Title                 string      `json:"title"                 description:"标题"`
	Description           string      `json:"description"           description:"介绍"`
	Version               string      `json:"version"               description:"版本"`
	Author                string      `json:"author"                description:"作者"`
	Icon                  string      `json:"icon"                  description:"插件图标"`
	Link                  string      `json:"link"                  description:"插件的网址。指向插件的 github 链接。值应为一个可访问的网址"`
	Command               string      `json:"command"               description:"插件的运行指令"`
	Args                  string      `json:"args"                  description:"插件的指令参数"`
	Status                int         `json:"status"                description:"状态"`
	FrontendUi            int         `json:"frontendUi"            description:"是否有插件页面"`
	FrontendUrl           string      `json:"frontendUrl"           description:"插件页面地址"`
	FrontendConfiguration int         `json:"frontendConfiguration" description:"是否显示配置页面"`
	StartTime             *gtime.Time `json:"startTime"             description:"启动时间"`
	IsDeleted             int         `json:"isDeleted"             description:"是否删除 0未删除 1已删除"`
	CreatedBy             uint        `json:"createdBy"             description:"创建者"`
	CreatedAt             *gtime.Time `json:"createdAt"             description:"创建日期"`
	UpdatedBy             int         `json:"updatedBy"             description:"修改人"`
	UpdatedAt             *gtime.Time `json:"updatedAt"             description:"更新时间"`
	DeletedBy             int         `json:"deletedBy"             description:"删除人"`
	DeletedAt             *gtime.Time `json:"deletedAt"             description:"删除时间"`
}

type GetSysPluginsListInput struct {
	Status int `json:"status"                description:"状态"`
	PaginationInput
}

type SysPluginsRes struct {
	Id                    int         `json:"id"                    description:"ID"`
	Types                 string      `json:"types"                 description:"插件与SagooIOT的通信方式"`
	HandleType            string      `json:"handleType"            description:"功能类型"`
	Name                  string      `json:"name"                  description:"名称"`
	Title                 string      `json:"title"                 description:"标题"`
	Description           string      `json:"description"           description:"介绍"`
	Version               string      `json:"version"               description:"版本"`
	Author                string      `json:"author"                description:"作者"`
	Icon                  string      `json:"icon"                  description:"插件图标"`
	Link                  string      `json:"link"                  description:"插件的网址。指向插件的 github 链接。值应为一个可访问的网址"`
	Command               string      `json:"command"               description:"插件的运行指令"`
	Args                  string      `json:"args"                  description:"插件的指令参数"`
	Status                int         `json:"status"                description:"状态"`
	FrontendUi            int         `json:"frontendUi"            description:"是否有插件页面"`
	FrontendUrl           string      `json:"frontendUrl"           description:"插件页面地址"`
	FrontendConfiguration int         `json:"frontendConfiguration" description:"是否显示配置页面"`
	StartTime             *gtime.Time `json:"startTime"             description:"启动时间"`
	IsDeleted             int         `json:"isDeleted"             description:"是否删除 0未删除 1已删除"`
	CreatedBy             uint        `json:"createdBy"             description:"创建者"`
	CreatedAt             *gtime.Time `json:"createdAt"             description:"创建日期"`
	UpdatedBy             int         `json:"updatedBy"             description:"修改人"`
	UpdatedAt             *gtime.Time `json:"updatedAt"             description:"更新时间"`
	DeletedBy             int         `json:"deletedBy"             description:"删除人"`
	DeletedAt             *gtime.Time `json:"deletedAt"             description:"删除时间"`
}

type SysPluginsAddInput struct {
	Id                    int         `json:"id"                    description:"ID"`
	Types                 string      `json:"types"                 description:"插件与SagooIOT的通信方式"`
	HandleType            string      `json:"handleType"            description:"功能类型"`
	Name                  string      `json:"name"                  description:"名称"`
	Title                 string      `json:"title"                 description:"标题"`
	Description           string      `json:"description"           description:"介绍"`
	Version               string      `json:"version"               description:"版本"`
	Author                string      `json:"author"                description:"作者"`
	Icon                  string      `json:"icon"                  description:"插件图标"`
	Link                  string      `json:"link"                  description:"插件的网址。指向插件的 github 链接。值应为一个可访问的网址"`
	Command               string      `json:"command"               description:"插件的运行指令"`
	Args                  string      `json:"args"                  description:"插件的指令参数"`
	Status                int         `json:"status"                description:"状态"`
	FrontendUi            int         `json:"frontendUi"            description:"是否有插件页面"`
	FrontendUrl           string      `json:"frontendUrl"           description:"插件页面地址"`
	FrontendConfiguration int         `json:"frontendConfiguration" description:"是否显示配置页面"`
	StartTime             *gtime.Time `json:"startTime"             description:"启动时间"`
	IsDeleted             int         `json:"isDeleted"             description:"是否删除 0未删除 1已删除"`
}
type SysPluginsEditInput struct {
	Id                    int      `json:"id"                    description:"ID"`
	Types                 string   `json:"types"                 description:"插件与SagooIOT的通信方式"`
	HandleType            string   `json:"handleType"            description:"功能类型"`
	Name                  string   `json:"name"                  description:"名称"`
	Title                 string   `json:"title"                 description:"标题"`
	Description           string   `json:"description"           description:"介绍"`
	Version               string   `json:"version"               description:"版本"`
	Author                string   `json:"author"                description:"作者"`
	Icon                  string   `json:"icon"                  description:"插件图标"`
	Link                  string   `json:"link"                  description:"插件的网址。指向插件的 github 链接。值应为一个可访问的网址"`
	Command               string   `json:"command"               description:"插件的运行指令"`
	Args                  []string `json:"args"                  description:"插件的指令参数"`
	FrontendUi            int      `json:"frontendUi"            description:"是否有插件页面"`
	FrontendUrl           string   `json:"frontendUrl"           description:"插件页面地址"`
	FrontendConfiguration int      `json:"frontendConfiguration" description:"是否显示配置页面"`
}

type SysPluginsInfoRes struct {
	Types      string `json:"types"                 description:"插件与SagooIOT的通信方式"`
	HandleType string `json:"handleType"            description:"功能类型"`
	Name       string `json:"name"                  description:"名称"`
	Title      string `json:"title"                 description:"标题"`
}

type SysPluginsInfoOut struct {
	Types      string `json:"types"                 description:"插件与SagooIOT的通信方式"`
	HandleType string `json:"handleType"            description:"功能类型"`
	Name       string `json:"name"                  description:"名称"`
	Title      string `json:"title"                 description:"标题"`
}
