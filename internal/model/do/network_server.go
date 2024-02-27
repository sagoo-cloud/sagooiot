// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NetworkServer is the golang structure of table network_server for DAO operations like Where/Data.
type NetworkServer struct {
	g.Meta        `orm:"table:network_server, do:true"`
	Id            interface{} //
	DeptId        interface{} // 部门ID
	Name          interface{} //
	Types         interface{} // tcp/udp
	Addr          interface{} //
	Register      interface{} // 注册包
	Heartbeat     interface{} // 心跳包
	Protocol      interface{} // 协议
	Devices       interface{} // 默认设备
	Status        interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	CreateBy      interface{} //
	Remark        interface{} // 备注
	IsTls         interface{} // 开启TLS:1=是，0=否
	AuthType      interface{} // 认证方式（1=Basic，2=AccessToken，3=证书）
	AuthUser      interface{} // 认证用户
	AuthPasswd    interface{} // 认证密码
	AccessToken   interface{} // AccessToken
	CertificateId interface{} // 证书ID
	Stick         interface{} // 粘包处理方式
}
