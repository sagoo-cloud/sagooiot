// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDeviceGateway is the golang structure of table dev_device_gateway for DAO operations like Where/Data.
type DevDeviceGateway struct {
	g.Meta     `orm:"table:dev_device_gateway, do:true"`
	Id         interface{} //
	GatewayKey interface{} // 网关标识
	SubKey     interface{} // 子设备标识
	CreatedBy  interface{} // 创建者
	UpdatedBy  interface{} // 更新者
	DeletedBy  interface{} // 删除者
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
