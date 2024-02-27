// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDeviceTag is the golang structure for table dev_device_tag.
type DevDeviceTag struct {
	Id        uint        `json:"id"        description:""`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	DeviceId  uint        `json:"deviceId"  description:"设备ID"`
	DeviceKey string      `json:"deviceKey" description:"设备标识"`
	Key       string      `json:"key"       description:"标签标识"`
	Name      string      `json:"name"      description:"标签名称"`
	Value     string      `json:"value"     description:"标签值"`
	CreatedBy uint        `json:"createdBy" description:"创建者"`
	UpdatedBy uint        `json:"updatedBy" description:"更新者"`
	DeletedBy uint        `json:"deletedBy" description:"删除者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
