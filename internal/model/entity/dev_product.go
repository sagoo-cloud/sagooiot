// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DevProduct is the golang structure for table dev_product.
type DevProduct struct {
	Id                uint        `json:"id"                description:""`
	Key               string      `json:"key"               description:"产品标识"`
	Name              string      `json:"name"              description:"产品名称"`
	CategoryId        uint        `json:"categoryId"        description:"所属品类"`
	MessageProtocol   string      `json:"messageProtocol"   description:"消息协议"`
	TransportProtocol string      `json:"transportProtocol" description:"传输协议: MQTT,COAP,UDP"`
	ProtocolId        uint        `json:"protocolId"        description:"协议id"`
	DeviceType        string      `json:"deviceType"        description:"设备类型: 网关，设备"`
	Desc              string      `json:"desc"              description:"描述"`
	Icon              string      `json:"icon"              description:"图片地址"`
	Metadata          string      `json:"metadata"          description:"物模型"`
	MetadataTable     int         `json:"metadataTable"     description:"是否生成物模型表：0=否，1=是"`
	Policy            string      `json:"policy"            description:"采集策略"`
	Status            int         `json:"status"            description:"发布状态：0=未发布，1=已发布"`
	CreateBy          uint        `json:"createBy"          description:"创建者"`
	UpdateBy          uint        `json:"updateBy"          description:"更新者"`
	DeletedBy         uint        `json:"deletedBy"         description:"删除者"`
	CreatedAt         *gtime.Time `json:"createdAt"         description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         description:"更新时间"`
	DeletedAt         *gtime.Time `json:"deletedAt"         description:"删除时间"`
}
