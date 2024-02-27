package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type GetNetworkTunnelListInput struct {
	ServiceId int    `json:"serviceId" dc:"服务ID"`
	DeviceKey string `json:"deviceKey" dc:"设备标识"`
	*PaginationInput
}
type NetworkTunnelOut struct {
	Id        int         `json:"id"          description:"ID"`
	ServerId  int         `json:"serverId"  description:"服务ID"`
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:""`
	Addr      string      `json:"addr"      description:""`
	Remote    string      `json:"remote"    description:""`
	Retry     string      `json:"retry"     description:""`
	Heartbeat string      `json:"heartbeat" description:""`
	Serial    string      `json:"serial"    description:""`
	Protoccol string      `json:"protoccol" description:""`
	DeviceKey string      `json:"deviceKey" description:"设备标识"`
	Status    int         `json:"status"    description:""`
	Last      *gtime.Time `json:"last"      description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Remark    string      `json:"remark"    description:"备注"`
}

type NetworkTunnelRes struct {
	Id        int         `json:"id"          description:"ID"`
	ServerId  int         `json:"serverId"  description:"服务ID"`
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:""`
	Addr      string      `json:"addr"      description:""`
	Remote    string      `json:"remote"    description:""`
	Retry     string      `json:"retry"     description:""`
	Heartbeat string      `json:"heartbeat" description:""`
	Serial    string      `json:"serial"    description:""`
	Protoccol string      `json:"protoccol" description:""`
	DeviceKey string      `json:"deviceKey" description:"设备标识"`
	Status    int         `json:"status"    description:""`
	Last      *gtime.Time `json:"last"      description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Remark    string      `json:"remark"    description:"备注"`
}
type NetworkTunnelAddInput struct {
	ServerId  int    `json:"serverId"  description:"服务ID"`
	Name      string `json:"name"      description:""`
	Types     string `json:"types"     description:""`
	Addr      string `json:"addr"      description:""`
	Remote    string `json:"remote"    description:""`
	Retry     string `json:"retry"     description:""`
	Heartbeat string `json:"heartbeat" description:""`
	Serial    string `json:"serial"    description:""`
	Protocol  string `json:"protocol" description:""`
	Status    int    `json:"status"    description:""`
	Remark    string `json:"remark"    description:"备注"`
}
type NetworkTunnelEditInput struct {
	Id int `json:"id"          description:"ID"`
	NetworkTunnelAddInput
}
