// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDevice is the golang structure for table dev_device.
type DevDevice struct {
	Id             uint        `json:"id"             description:""`
	DeptId         int         `json:"deptId"         description:"部门ID"`
	Key            string      `json:"key"            description:"设备标识"`
	Name           string      `json:"name"           description:"设备名称"`
	ProductKey     string      `json:"productKey"     description:"所属产品KEY"`
	Desc           string      `json:"desc"           description:"描述"`
	MetadataTable  int         `json:"metadataTable"  description:"是否生成物模型子表：0=否，1=是"`
	Status         int         `json:"status"         description:"状态：0=未启用,1=离线,2=在线"`
	OnlineTimeout  int         `json:"onlineTimeout"  description:"设备在线超时设置，单位：秒"`
	RegistryTime   *gtime.Time `json:"registryTime"   description:"激活时间"`
	LastOnlineTime *gtime.Time `json:"lastOnlineTime" description:"最后上线时间"`
	Version        string      `json:"version"        description:"固件版本号"`
	TunnelId       int         `json:"tunnelId"       description:"tunnelId"`
	Lng            string      `json:"lng"            description:"经度"`
	Lat            string      `json:"lat"            description:"纬度"`
	AuthType       int         `json:"authType"       description:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser       string      `json:"authUser"       description:"认证用户"`
	AuthPasswd     string      `json:"authPasswd"     description:"认证密码"`
	AccessToken    string      `json:"accessToken"    description:"AccessToken"`
	CertificateId  int         `json:"certificateId"  description:"证书ID"`
	CreatedBy      uint        `json:"createdBy"      description:"创建者"`
	UpdatedBy      uint        `json:"updatedBy"      description:"更新者"`
	DeletedBy      uint        `json:"deletedBy"      description:"删除者"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
}
