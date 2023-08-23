package model

type GetRemoteconfListInput struct {
	PaginationInput
	RemoteconfInput
}
type RemoteconfInput struct {
	ProductKey      string `json:"productKey"          description:"产品key"`
}
type RemoteconfListOutput struct {
	Data []*RemoteconfOutput `json:"remoteconf" dc:"远程配置列表"`
	PaginationOutput
}
type RemoteconfOutput struct {
	ConfigContent   string `json:"configContent"          description:"配置内容"`
	GmtCreate       string `json:"gmtCreate"          description:"创建时间"`
	ConfigFormat    string `json:"configFormat"          description:"配置格式，json等"`
	ConfigNumber    string `json:"configNumber"          description:"配置编号"`

}
type RemoteconfAddInput struct {
	Scope           string `json:"scope"          description:"配置范围：产品=product 设备=device"`
	ConfigFormat    string `json:"configFormat"          description:"配置格式，json等"`
	ConfigContent   string `json:"configContent"          description:"配置内容"`
	ConfigSize      string `json:"configSize"          description:"配置文件大小（按字节计算）"`
	Status          string `json:"status"          description:"状态： 0=停用 1=启用"`
	ConfigName      string `json:"configName"          description:"配置名称"`
	ProductKey      string `json:"productKey"          description:"产品key"`
}
type RemoteconfEditInput struct {
	Id int `json:"id"          description:"ID"`
	RemoteconfAddInput
}
