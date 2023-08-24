package model

type GetRemoteconfListInput struct {
	PaginationInput
	RemoteconfInput
}
type RemoteconfInput struct {
	ProductKey string `json:"productKey"          description:"产品key"`
}
type RemoteconfListOutput struct {
	Data []*RemoteconfOutput `json:"remoteconf" dc:"远程配置列表"`
	PaginationOutput
}
type RemoteconfOutput struct {
	ConfigContent string `json:"configContent"          description:"配置内容"`
	GmtCreate     string `json:"gmtCreate"          description:"创建时间"`
	ConfigFormat  string `json:"configFormat"          description:"配置格式，json等"`
	ConfigNumber  string `json:"configNumber"          description:"配置编号"`
	Id            string `json:"id"          description:"配置ID"`
	Scope         string `json:"scope"          description:"配置范围：产品=product 设备=device"`
}
type RemoteconfAddInput struct {
	Scope         string `json:"scope"          description:"配置范围：产品=product 设备=device" v:"required#配置范围不能为空"`
	ConfigFormat  string `json:"configFormat"          description:"配置格式，json等" v:"required#配置格式不能为空"`
	ConfigContent string `json:"configContent"          description:"配置内容"`
	ConfigSize    string `json:"configSize"          description:"配置文件大小（按kb计算）" v:"required#配置文件大小不能为空"`
	Status        string `json:"status"          description:"状态： 0=停用 1=启用" v:"required#配置状态不能为空"`
	ProductKey    string `json:"productKey"          description:"产品key"  v:"required#产品key不能为空"`
}
type RemoteconfEditInput struct {
	Id            string `json:"id"          description:"配置ID" v:"required#配置ID不能为空"`
}
