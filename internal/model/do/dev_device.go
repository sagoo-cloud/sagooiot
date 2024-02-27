// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDevice is the golang structure of table dev_device for DAO operations like Where/Data.
type DevDevice struct {
	g.Meta         `orm:"table:dev_device, do:true"`
	Id             interface{} //
	DeptId         interface{} // 部门ID
	Key            interface{} // 设备标识
	Name           interface{} // 设备名称
	ProductKey     interface{} // 所属产品KEY
	Desc           interface{} // 描述
	MetadataTable  interface{} // 是否生成物模型子表：0=否，1=是
	Status         interface{} // 状态：0=未启用,1=离线,2=在线
	OnlineTimeout  interface{} // 设备在线超时设置，单位：秒
	RegistryTime   *gtime.Time // 激活时间
	LastOnlineTime *gtime.Time // 最后上线时间
	Version        interface{} // 固件版本号
	TunnelId       interface{} // tunnelId
	Lng            interface{} // 经度
	Lat            interface{} // 纬度
	AuthType       interface{} // 认证方式（1=Basic，2=AccessToken，3=证书）
	AuthUser       interface{} // 认证用户
	AuthPasswd     interface{} // 认证密码
	AccessToken    interface{} // AccessToken
	CertificateId  interface{} // 证书ID
	CreatedBy      interface{} // 创建者
	UpdatedBy      interface{} // 更新者
	DeletedBy      interface{} // 删除者
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
	DeletedAt      *gtime.Time // 删除时间
}
