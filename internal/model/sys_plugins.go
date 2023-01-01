package model

type GetSysPluginsListInput struct {
	PaginationInput
}

type SysPluginsOutput struct {
	Intro     string `json:"intro"          description:"介绍"`
	Status    int    `json:"status"          description:"状态"`
	Types     string `json:"types"          description:"插件类型"`
	StartTime string `json:"startTime"          description:""`
	Id        int    `json:"id"          description:"ID"`
	Name      string `json:"name"          description:"名称"`
	Title     string `json:"title"          description:"标题"`
	Version   string `json:"version"          description:"版本"`
	Author    string `json:"author"          description:""`
}
type SysPluginsAddInput struct {
	Version   string `json:"version"          description:"版本"`
	Author    string `json:"author"          description:""`
	Status    int    `json:"status"          description:"状态"`
	Types     string `json:"types"          description:"插件类型"`
	StartTime string `json:"startTime"          description:""`
	Name      string `json:"name"          description:"名称"`
	Title     string `json:"title"          description:"标题"`
	Intro     string `json:"intro"          description:"介绍"`
}
type SysPluginsEditInput struct {
	Id int `json:"id"          description:"ID"`
	SysPluginsAddInput
}
