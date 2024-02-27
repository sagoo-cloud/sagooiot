package model

import "github.com/gogf/gf/v2/os/gtime"

type NetworkServerOut struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:"tcp/udp/mqtt_server/http/websocket"`
	Addr      string      `json:"addr"      description:""`
	Register  string      `json:"register"  description:"注册包"`
	Heartbeat string      `json:"heartbeat" description:"心跳包"`
	Protocol  string      `json:"protocol"  description:"协议"`
	Devices   string      `json:"devices"   description:"默认设备"`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	CreateBy  int         `json:"createBy"  description:""`
	Remark    string      `json:"remark"    description:"备注"`

	// 认证信息
	IsTls         uint   `json:"isTls" dc:"开启TLS:1=是，0=否"`
	AuthType      int    `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string `json:"authUser" dc:"认证用户"`
	AuthPasswd    string `json:"authPasswd" dc:"认证密码"`
	AccessToken   string `json:"accessToken" dc:"AccessToken"`
	CertificateId int    `json:"certificateId" dc:"证书ID"`
	Stick         string `json:"stick" dc:"粘包处理方式"`
}

type NetworkServerRes struct {
	Id        int         `json:"id"        description:""`
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:"tcp/udp/mqtt_server/http/websocket"`
	Addr      string      `json:"addr"      description:""`
	Register  string      `json:"register"  description:"注册包"`
	Heartbeat string      `json:"heartbeat" description:"心跳包"`
	Protocol  string      `json:"protocol"  description:"协议"`
	Devices   string      `json:"devices"   description:"默认设备"`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	CreateBy  int         `json:"createBy"  description:""`
	Remark    string      `json:"remark"    description:"备注"`

	// 认证信息
	IsTls         uint   `json:"isTls" dc:"开启TLS:1=是，0=否"`
	AuthType      int    `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string `json:"authUser" dc:"认证用户"`
	AuthPasswd    string `json:"authPasswd" dc:"认证密码"`
	AccessToken   string `json:"accessToken" dc:"AccessToken"`
	CertificateId int    `json:"certificateId" dc:"证书ID"`
	Stick         string `json:"stick" dc:"粘包处理方式"`
}
type NetworkServerAddInput struct {
	Name      string      `json:"name"      description:""`
	Types     string      `json:"types"     description:"tcp/udp/mqtt_server/http/websocket"`
	Addr      string      `json:"addr"      description:""`
	Register  string      `json:"register"  description:"注册包"`
	Heartbeat string      `json:"heartbeat" description:"心跳包"`
	Protocol  string      `json:"protocol"  description:"协议"`
	Devices   string      `json:"devices"   description:"默认设备"`
	Status    int         `json:"status"    description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	CreateBy  int         `json:"createBy"  description:""`
	Remark    string      `json:"remark"    description:"备注"`

	// 认证信息
	IsTls         uint   `json:"isTls" dc:"开启TLS:1=是，0=否"`
	AuthType      int    `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string `json:"authUser" dc:"认证用户"`
	AuthPasswd    string `json:"authPasswd" dc:"认证密码"`
	AccessToken   string `json:"accessToken" dc:"AccessToken"`
	CertificateId int    `json:"certificateId" dc:"证书ID"`
	Stick         Stick  `json:"stick" dc:"粘包处理方式"`
}
type NetworkServerEditInput struct {
	Id int `json:"id"          description:"ID"`
	NetworkServerAddInput
}

type GetNetworkServerListInput struct {
	PaginationInput
}

// 粘包处理方式
type Stick struct {
	Delimit  string `json:"delimit,omitempty" dc:"分隔符"`
	Custom   string `json:"custom,omitempty" dc:"自定义脚本"`
	FixedLen int    `json:"fixedLen,omitempty" dc:"固定长度"`
	Len      struct {
		Len    int    `json:"len" dc:"长度"`
		Offset int    `json:"offset" dc:"偏移量"`
		Endian string `json:"endian" dc:"大小端(big|little)"`
	} `json:"len,omitempty" dc:"长度字段"`
}
