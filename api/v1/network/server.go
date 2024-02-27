package network

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 这里是需要处理的地方，需要在这里调试好下面的几个接口

// 获取列表api
type GetNetworkServerListReq struct {
	g.Meta `path:"/server/list" method:"get" summary:"获取通讯服务列表" tags:"网络组件管理"`
	common.PaginationReq
}
type GetNetworkServerListRes struct {
	Data []*model.NetworkServerRes
	common.PaginationRes
}

// 获取指定ID的数据api
type GetNetworkServerByIdReq struct {
	g.Meta `path:"/get" method:"get" summary:"获取通讯服务详情" tags:"网络组件管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetNetworkServerByIdRes struct {
	Data *model.NetworkServerRes
}

// 添加数据api
type AddNetworkServerReq struct {
	g.Meta    `path:"/server/add" method:"post" summary:"添加通讯服务" tags:"网络组件管理"`
	Name      string `json:"name"      description:""`
	Types     string `json:"types"     description:"tcp/udp"`
	Addr      string `json:"addr"      description:""`
	Register  string `json:"register"  description:"注册包"`
	Heartbeat string `json:"heartbeat" description:"心跳包"`
	Protocol  string `json:"protocol"  description:"协议"`
	Devices   string `json:"devices"   description:"默认设备"`
	Remark    string `json:"remark"    description:"备注"`
	Status    int    `json:"status"    description:""`
	// 认证信息
	IsTls         uint        `json:"isTls" dc:"开启TLS:1=是，0=否"`
	AuthType      int         `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string      `json:"authUser" dc:"认证用户"`
	AuthPasswd    string      `json:"authPasswd" dc:"认证密码"`
	AccessToken   string      `json:"accessToken" dc:"AccessToken"`
	CertificateId int         `json:"certificateId" dc:"证书ID"`
	Stick         model.Stick `json:"stick" dc:"粘包处理方式"`
}
type AddNetworkServerRes struct{}

// 编辑数据api
type EditNetworkServerReq struct {
	g.Meta    `path:"/server/edit" method:"put" summary:"编辑通讯服务" tags:"网络组件管理"`
	Id        int    `json:"id"        description:"id" v:"required#id不能为空"`
	Name      string `json:"name"      description:""`
	Types     string `json:"types"     description:"tcp/udp"`
	Addr      string `json:"addr"      description:""`
	Register  string `json:"register"  description:"注册包"`
	Heartbeat string `json:"heartbeat" description:"心跳包"`
	Protocol  string `json:"protocol"  description:"协议"`
	Devices   string `json:"devices"   description:"默认设备"`
	Status    int    `json:"status"    description:"状态"`
	Remark    string `json:"remark"    description:"备注"`
	// 认证信息
	IsTls         uint        `json:"isTls" dc:"开启TLS:1=是，0=否"`
	AuthType      int         `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string      `json:"authUser" dc:"认证用户"`
	AuthPasswd    string      `json:"authPasswd" dc:"认证密码"`
	AccessToken   string      `json:"accessToken" dc:"AccessToken"`
	CertificateId int         `json:"certificateId" dc:"证书ID"`
	Stick         model.Stick `json:"stick" dc:"粘包处理方式"`
}
type EditNetworkServerRes struct{}

// 删除数据api
type DeleteNetworkServerReq struct {
	g.Meta `path:"/server/delete" method:"delete" summary:"删除通讯服务" tags:"网络组件管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteNetworkServerRes struct{}

// 服务状态api
type SetNetworkServerStatusReq struct {
	g.Meta `path:"/server/status" method:"post" summary:"修改服务状态" tags:"网络组件管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
	Status int `json:"status"        description:"status" v:"required#status不能为空"`
}
type SetNetworkServerStatusRes struct{}
