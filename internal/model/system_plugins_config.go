package model

type GetPluginsConfigListInput struct {
	Id int `json:"id"          description:"ID"`
	PaginationInput
}
type PluginsConfigListOutput struct {
	Data []PluginsConfigOutput
	PaginationOutput
}
type PluginsConfigOutput struct {
	Value string `json:"value"          description:"配置内容"`
	Doc   string `json:"doc"          description:"配置说明"`
	Id    string `json:"id"          description:""`
	Type  string `json:"type"          description:"插件类型"`
	Name  string `json:"name"          description:"插件名称"`
}
type PluginsConfigAddInput struct {
	Type  string `json:"type"          description:"插件类型"`
	Name  string `json:"name"          description:"插件名称"`
	Value string `json:"value"          description:"配置内容"`
	Doc   string `json:"doc"          description:"配置说明"`
}
type PluginsConfigEditInput struct {
	Id int `json:"id"          description:"ID"`
	PluginsConfigAddInput
}

type PluginsConfigData struct {
	Msg  string
	Data interface{}
}
