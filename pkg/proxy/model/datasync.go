package model

import "github.com/gogf/gf/v2/os/gtime"

type SyncProduct struct {
	TenantCode        string      `json:"tenantCode"        description:"租户code"`
	Id                uint        `json:"id"                description:""`
	DeptId            int         `json:"deptId"            description:"部门ID"`
	Key               string      `json:"key"               description:"产品标识"`
	Name              string      `json:"name"              description:"产品名称"`
	CategoryId        uint        `json:"categoryId"        description:"所属品类"`
	MessageProtocol   string      `json:"messageProtocol"   description:"消息协议"`
	TransportProtocol string      `json:"transportProtocol" description:"传输协议: MQTT,COAP,UDP"`
	ProtocolId        uint        `json:"protocolId"        description:"协议id"`
	DeviceType        string      `json:"deviceType"        description:"设备类型: 网关，设备，子设备"`
	Desc              string      `json:"desc"              description:"描述"`
	Icon              string      `json:"icon"              description:"图片地址"`
	Metadata          string      `json:"metadata"          description:"物模型"`
	MetadataTable     int         `json:"metadataTable"     description:"是否生成物模型表：0=否，1=是"`
	Policy            string      `json:"policy"            description:"采集策略"`
	Status            int         `json:"status"            description:"发布状态：0=未发布，1=已发布"`
	AuthType          int         `json:"authType"          description:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser          string      `json:"authUser"          description:"认证用户"`
	AuthPasswd        string      `json:"authPasswd"        description:"认证密码"`
	AccessToken       string      `json:"accessToken"       description:"AccessToken"`
	CertificateId     int         `json:"certificateId"     description:"证书ID"`
	ScriptInfo        string      `json:"scriptInfo"        description:"脚本信息"`
	CreatedBy         uint        `json:"createdBy"         description:"创建者"`
	UpdatedBy         uint        `json:"updatedBy"         description:"更新者"`
	DeletedBy         uint        `json:"deletedBy"         description:"删除者"`
	CreatedAt         *gtime.Time `json:"createdAt"         description:"创建时间"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         description:"更新时间"`
	DeletedAt         *gtime.Time `json:"deletedAt"         description:"删除时间"`
}

type SyncDevice struct {
	TenantCode     string      `json:"tenantCode"        description:"租户code"`
	Id             uint        `json:"id"             description:""`
	DeptId         int         `json:"deptId"         description:"部门ID"`
	Key            string      `json:"key"            description:"设备标识"`
	Name           string      `json:"name"           description:"设备名称"`
	ProductId      uint        `json:"productId"      description:"所属产品"`
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
	Address        string      `json:"address"        description:"详细地址"`
	AuthType       int         `json:"authType"       description:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser       string      `json:"authUser"       description:"认证用户"`
	AuthPasswd     string      `json:"authPasswd"     description:"认证密码"`
	AccessToken    string      `json:"accessToken"    description:"AccessToken"`
	CertificateId  int         `json:"certificateId"  description:"证书ID"`
	ExtensionInfo  string      `json:"extensionInfo"  description:"设备扩展信息"`
	CreatedBy      uint        `json:"createdBy"      description:"创建者"`
	UpdatedBy      uint        `json:"updatedBy"      description:"更新者"`
	DeletedBy      uint        `json:"deletedBy"      description:"删除者"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"更新时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
	ProductKey     string      `json:"productKey"     description:"产品key"`

	Tags []*DevDeviceTag `json:"tags" orm:"with:device_id=id" dc:"设备标签"`
}
