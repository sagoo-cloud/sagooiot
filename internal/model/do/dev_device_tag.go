// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDeviceTag is the golang structure of table dev_device_tag for DAO operations like Where/Data.
type DevDeviceTag struct {
	g.Meta    `orm:"table:dev_device_tag, do:true"`
	Id        interface{} //
	DeptId    interface{} // 部门ID
	DeviceId  interface{} // 设备ID
	DeviceKey interface{} // 设备标识
	Key       interface{} // 标签标识
	Name      interface{} // 标签名称
	Value     interface{} // 标签值
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	DeletedBy interface{} // 删除者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
