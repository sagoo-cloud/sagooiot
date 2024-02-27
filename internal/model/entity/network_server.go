// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// NetworkServer is the golang structure for table network_server.
type NetworkServer struct {
	Id            int         `json:"id"            description:""`
	DeptId        int         `json:"deptId"        description:"部门ID"`
	Name          string      `json:"name"          description:""`
	Types         string      `json:"types"         description:"tcp/udp"`
	Addr          string      `json:"addr"          description:""`
	Register      string      `json:"register"      description:"注册包"`
	Heartbeat     string      `json:"heartbeat"     description:"心跳包"`
	Protocol      string      `json:"protocol"      description:"协议"`
	Devices       string      `json:"devices"       description:"默认设备"`
	Status        int         `json:"status"        description:""`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:""`
	CreateBy      int         `json:"createBy"      description:""`
	Remark        string      `json:"remark"        description:"备注"`
	IsTls         uint        `json:"isTls"         description:"开启TLS:1=是，0=否"`
	AuthType      int         `json:"authType"      description:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string      `json:"authUser"      description:"认证用户"`
	AuthPasswd    string      `json:"authPasswd"    description:"认证密码"`
	AccessToken   string      `json:"accessToken"   description:"AccessToken"`
	CertificateId int         `json:"certificateId" description:"证书ID"`
	Stick         string      `json:"stick"         description:"粘包处理方式"`
}
