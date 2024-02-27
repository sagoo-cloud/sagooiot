// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// NetworkTunnel is the golang structure of table network_tunnel for DAO operations like Where/Data.
type NetworkTunnel struct {
	g.Meta    `orm:"table:network_tunnel, do:true"`
	Id        interface{} //
	DeptId    interface{} // 部门ID
	ServerId  interface{} // 服务ID
	Name      interface{} //
	Types     interface{} //
	Addr      interface{} //
	Remote    interface{} //
	Retry     interface{} // 断线重连
	Heartbeat interface{} // 心跳包
	Serial    interface{} // 串口参数
	Protoccol interface{} // 适配协议
	DeviceKey interface{} // 设备标识
	Status    interface{} //
	Last      *gtime.Time //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	Remark    interface{} // 备注
}
