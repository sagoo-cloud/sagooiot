// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDeviceGateway is the golang structure for table dev_device_gateway.
type DevDeviceGateway struct {
	Id         uint        `json:"id"         description:""`
	GatewayKey string      `json:"gatewayKey" description:"网关标识"`
	SubKey     string      `json:"subKey"     description:"子设备标识"`
	CreatedBy  uint        `json:"createdBy"  description:"创建者"`
	UpdatedBy  uint        `json:"updatedBy"  description:"更新者"`
	DeletedBy  uint        `json:"deletedBy"  description:"删除者"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:"删除时间"`
}
