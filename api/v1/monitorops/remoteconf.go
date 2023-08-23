package monitorops

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/model"
)

// GetRemoteconfListReq 获取数据列表
type GetRemoteconfListReq struct {
	g.Meta `path:"/remoteconf/queryThingConfig" method:"get" summary:"获取远程配置列表" tags:"监控运维"`
	model.GetRemoteconfListInput
}
type GetRemoteconfListRes struct {
	model.RemoteconfListOutput
}

// AddRemoteconfReq 添加数据
type AddRemoteconfReq struct {
	g.Meta          `path:"/remoteconf/addThingConfig" method:"post" summary:"添加远程配置" tags:"监控运维"`
	*model.RemoteconfAddInput
}
type AddRemoteconfRes struct{}

// EditRemoteconfReq 编辑数据api
type EditRemoteconfReq struct {
	g.Meta          `path:"/remoteconf/edit" method:"put" summary:"编辑远程配置" tags:"监控运维"`
	*model.RemoteconfEditInput
}
type EditRemoteconfRes struct{}

// DeleteRemoteconfReq 删除数据
type DeleteRemoteconfReq struct {
	g.Meta `path:"/remoteconf/delete" method:"delete" summary:"删除远程配置" tags:"监控运维"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteRemoteconfRes struct{}

// GetRemoteconfByIdReq 获取指定ID的数据
type GetRemoteconfByIdReq struct {
	g.Meta `path:"/remoteconf/get" method:"get" summary:"获取远程配置" tags:"监控运维"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetRemoteconfByIdRes struct {
	Scope           string `json:"scope"          description:"配置范围：产品=product 设备=device"`
	OssPath         string `json:"ossPath"          description:"Oss文件位置"`
	UtcCreate       string `json:"utcCreate"          description:"UTC格式的创建时间"`
	ConfigFormat    string `json:"configFormat"          description:"配置格式，json等"`
	ConfigSize      string `json:"configSize"          description:"配置文件大小（按字节计算）"`
	Status          string `json:"status"          description:"状态： 0=停用 1=启用"`
	OssUrl          string `json:"ossUrl"          description:"Oss链接"`
	Sign            string `json:"sign"          description:"签名"`
	ConfigName      string `json:"configName"          description:"配置名称"`
	ProductKey      string `json:"productKey"          description:"产品key"`
	SignMethod      string `json:"signMethod"          description:"签名方式，sha256等"`
	GmtCreate       string `json:"gmtCreate"          description:"创建时间"`
	ConfigContent   string `json:"configContent"          description:"配置内容"`
	ContainedOssUrl string `json:"containedOssUrl"          description:"包含OssURL"`
	Id              string `json:"id"          description:"配置ID"`
}