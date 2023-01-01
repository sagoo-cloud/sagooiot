package model

import "github.com/gogf/gf/v2/os/gtime"

type NetworkServerOut struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:"tcp/udp/mqtt"`
	Addr      string      `json:"addr"      description:""`
	Register  string      `json:"register"  description:"注册包"`
	Heartbeat string      `json:"heartbeat" description:"心跳包"`
	Protocol  string      `json:"protocol"  description:"协议"`
	Devices   string      `json:"devices"   description:"默认设备"`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	CreateBy  int         `json:"createBy"  description:""`
	Remark    string      `json:"remark"    description:"备注"`
}

type NetworkServerRes struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:"tcp/udp/mqtt"`
	Addr      string      `json:"addr"      description:""`
	Register  string      `json:"register"  description:"注册包"`
	Heartbeat string      `json:"heartbeat" description:"心跳包"`
	Protocol  string      `json:"protocol"  description:"协议"`
	Devices   string      `json:"devices"   description:"默认设备"`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	CreateBy  int         `json:"createBy"  description:""`
	Remark    string      `json:"remark"    description:"备注"`
}
type NetworkServerAddInput struct {
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:"tcp/udp"`
	Addr      string      `json:"addr"      description:""`
	Register  string      `json:"register"  description:"注册包"`
	Heartbeat string      `json:"heartbeat" description:"心跳包"`
	Protocol  string      `json:"protocol"  description:"协议"`
	Devices   string      `json:"devices"   description:"默认设备"`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	CreateBy  int         `json:"createBy"  description:""`
	Remark    string      `json:"remark"    description:"备注"`
}
type NetworkServerEditInput struct {
	Id int `json:"id"          description:"ID"`
	NetworkServerAddInput
}

type GetNetworkServerListInput struct {
	PaginationInput
}
