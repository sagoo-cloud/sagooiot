// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DevProduct is the golang structure of table dev_product for DAO operations like Where/Data.
type DevProduct struct {
	g.Meta            `orm:"table:dev_product, do:true"`
	Id                interface{} //
	DeptId            interface{} // 部门ID
	Key               interface{} // 产品标识
	Name              interface{} // 产品名称
	CategoryId        interface{} // 所属品类
	MessageProtocol   interface{} // 消息协议
	TransportProtocol interface{} // 传输协议: MQTT,COAP,UDP
	ProtocolId        interface{} // 协议id
	DeviceType        interface{} // 设备类型: 网关，设备，子设备
	Desc              interface{} // 描述
	Icon              interface{} // 图片地址
	Metadata          interface{} // 物模型
	MetadataTable     interface{} // 是否生成物模型表：0=否，1=是
	Policy            interface{} // 采集策略
	Status            interface{} // 发布状态：0=未发布，1=已发布
	AuthType          interface{} // 认证方式（1=Basic，2=AccessToken，3=证书）
	AuthUser          interface{} // 认证用户
	AuthPasswd        interface{} // 认证密码
	AccessToken       interface{} // AccessToken
	CertificateId     interface{} // 证书ID
	ScriptInfo        interface{} // 脚本信息
	CreatedBy         interface{} // 创建者
	UpdatedBy         interface{} // 更新者
	DeletedBy         interface{} // 删除者
	CreatedAt         *gtime.Time // 创建时间
	UpdatedAt         *gtime.Time // 更新时间
	DeletedAt         *gtime.Time // 删除时间
}
