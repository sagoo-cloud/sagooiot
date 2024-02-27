package network

import (
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取列表api
type GetNetworkTunnelListReq struct {
	g.Meta    `path:"/tunnel/list" method:"get" summary:"获取通道列表" tags:"网络组件管理"`
	ServiceId int    `json:"serviceId" dc:"服务ID"`
	DeviceKey string `json:"deviceKey" dc:"设备标识"`
	*common.PaginationReq
}
type GetNetworkTunnelListRes struct {
	Data []*model.NetworkTunnelRes
	common.PaginationRes
}

// 获取指定ID的数据api
type GetNetworkTunnelByIdReq struct {
	g.Meta `path:"/tunnel/get" method:"get" summary:"获取单个通道" tags:"网络组件管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetNetworkTunnelByIdRes struct {
	Data *model.NetworkTunnelRes
}

// 添加数据api
type AddNetworkTunnelReq struct {
	g.Meta    `path:"/tunnel/add" method:"post" summary:"添加通道" tags:"网络组件管理"`
	Name      string `json:"name"          description:"" v:"required#名称不能为空"`
	Types     string `json:"types"          description:"" v:"required#类型不能为空"`
	Addr      string `json:"addr"          description:""`
	Remote    string `json:"remote"          description:""`
	Status    string `json:"status"          description:""`
	Retry     string `json:"retry"          description:""`
	Heartbeat string `json:"heartbeat"          description:""`
	Serial    string `json:"serial"          description:""`
	Protoccol string `json:"protoccol"          description:""`
	Remark    string `json:"remark"    description:"备注"`
}
type AddNetworkTunnelRes struct{}

// 编辑数据api
type EditNetworkTunnelReq struct {
	g.Meta    `path:"/tunnel/edit" method:"put" summary:"编辑通道" tags:"网络组件管理"`
	Id        int    `json:"id"        description:"id" v:"required#id不能为空"`
	Name      string `json:"name"          description:"" v:"required#名称不能为空"`
	Types     string `json:"types"          description:"" v:"required#类型不能为空"`
	Addr      string `json:"addr"          description:""`
	Remote    string `json:"remote"          description:""`
	Status    string `json:"status"          description:""`
	Retry     string `json:"retry"          description:""`
	Heartbeat string `json:"heartbeat"          description:""`
	Serial    string `json:"serial"          description:""`
	Protoccol string `json:"protoccol"          description:""`
	Remark    string `json:"remark"    description:"备注"`
}
type EditNetworkTunnelRes struct{}

// 删除数据api
type DeleteNetworkTunnelReq struct {
	g.Meta `path:"/tunnel/delete" method:"delete" summary:"删除通道" tags:"网络组件管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteNetworkTunnelRes struct{}

// 通道状态api
type SetNetworkTunnelStatusReq struct {
	g.Meta `path:"/tunnel/status" method:"post" summary:"修改通道状态" tags:"网络组件管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
	Status int `json:"status"        description:"status" v:"required#status不能为空"`
}
type SetNetworkTunnelStatusRes struct{}
