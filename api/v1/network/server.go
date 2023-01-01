package network

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/api/v1/common"
	"github.com/sagoo-cloud/sagooiot/internal/model"
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
	g.Meta `path:"/get" method:"get" summary:"获取通讯服务列表" tags:"网络组件管理"`
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
