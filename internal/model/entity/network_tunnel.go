// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NetworkTunnel is the golang structure for table network_tunnel.
type NetworkTunnel struct {
	Id        int         `json:"id"        description:""`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	ServerId  int         `json:"serverId"  description:"服务ID"`
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:""`
	Addr      string      `json:"addr"      description:""`
	Remote    string      `json:"remote"    description:""`
	Retry     string      `json:"retry"     description:"断线重连"`
	Heartbeat string      `json:"heartbeat" description:"心跳包"`
	Serial    string      `json:"serial"    description:"串口参数"`
	Protoccol string      `json:"protoccol" description:"适配协议"`
	DeviceKey string      `json:"deviceKey" description:"设备标识"`
	Status    int         `json:"status"    description:""`
	Last      *gtime.Time `json:"last"      description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Remark    string      `json:"remark"    description:"备注"`
}
