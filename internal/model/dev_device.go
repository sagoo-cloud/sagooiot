package model

import (
	"sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

const (
	DeviceStatusNoEnable int = iota // 设备未启用
	DeviceStatusOff                 // 设备离线
	DeviceStatusOn                  // 设备在线
)

type DeviceInput struct {
	Key        string `json:"key" dc:"设备标识"`
	Name       string `json:"name" dc:"设备名称"`
	ProductKey string `json:"productKey" dc:"所属产品"`
	TunnelId   int    `json:"tunnelId"       description:"tunnelId"`
	Status     string `p:"status"` //设备状态
}

type DevDevice struct {
	Id   uint   `json:"id" dc:"设备ID"`
	Name string `json:"name" dc:"设备名称"`
	Key  string `json:"key" dc:"设备标识"`
}
type DevProductWithName struct {
	gmeta.Meta `orm:"table:dev_product"`
	Id         uint   `json:"id" dc:"产品ID"`
	Name       string `json:"name" dc:"产品名称"`
	Key        string `json:"key" dc:"产品标识"`
}
type DevProduct struct {
	Id         uint   `json:"id" dc:"产品ID"`
	Name       string `json:"name" dc:"产品名称"`
	Key        string `json:"key" dc:"产品标识"`
	Metadata   string `json:"metadata" dc:"物模型"`
	DeviceType string `json:"deviceType" dc:"设备类型"`
}
type DevDeviceTag struct {
	Id       uint   `json:"id" dc:"标签ID"`
	DeviceId uint   `json:"deviceId" dc:"设备ID"`
	Key      string `json:"key" dc:"标签标识"`
	Name     string `json:"name" dc:"标签名称"`
	Value    string `json:"value" dc:"标签值"`
}

type DeviceOutput struct {
	*entity.DevDevice
	ProductName string `json:"productName" dc:"产品名称"`
	TSL         *TSL   `json:"tsl" dc:"物模型"`

	Product *entity.DevProduct `json:"product" orm:"with:key=product_key" dc:"产品信息"`
	Tags    []*DevDeviceTag    `json:"tags" orm:"with:device_id=id" dc:"设备标签"`
}

type AddDeviceInput struct {
	Key           string `json:"key" dc:"设备标识" v:"required#请输入设备标识"`
	Name          string `json:"name" dc:"设备名称" v:"required#请输入设备名称"`
	ProductKey    string `json:"productKey" dc:"所属产品" v:"required#请选择所属产品"`
	Desc          string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Version       string `json:"version" dc:"固件版本号"`
	Lng           string `json:"lng" dc:"经度"`
	Lat           string `json:"lat" dc:"纬度"`
	OnlineTimeout int    `json:"online_timeout" dc:"设备超时时间，单位秒"`

	Tags []AddTagInput `json:"tags" dc:"设备标签"`

	// 认证信息
	AuthType      int    `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string `json:"authUser" dc:"认证用户"`
	AuthPasswd    string `json:"authPasswd" dc:"认证密码"`
	AccessToken   string `json:"accessToken" dc:"AccessToken"`
	CertificateId int    `json:"certificateId" dc:"证书ID"`

	ExtensionInfo string `json:"extensionInfo" dc:"设备扩展信息"`

	Address string `json:"address" dc:"详细地址"`
}

type EditDeviceInput struct {
	Key     string        `json:"key" dc:"设备标识" v:"required#请输入设备标识"`
	Name    string        `json:"name" dc:"设备名称" v:"required#请输入设备名称"`
	Desc    string        `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Version string        `json:"version" dc:"固件版本号"`
	Lng     string        `json:"lng" dc:"经度"`
	Lat     string        `json:"lat" dc:"纬度"`
	Tags    []AddTagInput `json:"tags" dc:"设备标签"`

	// 认证信息
	AuthType      int    `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string `json:"authUser" dc:"认证用户"`
	AuthPasswd    string `json:"authPasswd" dc:"认证密码"`
	AccessToken   string `json:"accessToken" dc:"AccessToken"`
	CertificateId int    `json:"certificateId" dc:"证书ID"`

	ExtensionInfo string `json:"extensionInfo" dc:"设备扩展信息"`

	Address string `json:"address" dc:"详细地址"`
}

type DeviceExtendInput struct {
	Key           string `json:"key" dc:"设备标识" v:"required#请输入设备标识"`
	OnlineTimeout int    `json:"onlineTimeout" dc:"设备在线超时设置，单位：秒"`
	// 认证信息
	AuthType      int    `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string `json:"authUser" dc:"认证用户"`
	AuthPasswd    string `json:"authPasswd" dc:"认证密码"`
	AccessToken   string `json:"accessToken" dc:"AccessToken"`
	CertificateId int    `json:"certificateId" dc:"证书ID"`
}

type ListDeviceInput struct {
	ProductKey string `json:"productKey" dc:"产品ID"`
}

type ListForSubInput struct {
	ProductKey string `json:"productKey" dc:"产品标识" v:"required#产品KEY不能为空"`
	GatewayKey string `json:"gatewayKey" dc:"网关标识"`
	PaginationInput
}

type ListDeviceForPageInput struct {
	*DeviceInput
	PaginationInput
}
type ListDeviceForPageOutput struct {
	Device []*DeviceOutput `json:"device" dc:"设备列表"`
	PaginationOutput
}

// 设备运行状态
type DeviceRunStatusOutput struct {
	Status         int               `json:"status" dc:"状态：0=未启用,1=离线,2=在线"`
	LastOnlineTime *gtime.Time       `json:"lastOnlineTime" dc:"最后上线时间"`
	Properties     []DevicePropertiy `json:"properties" dc:"属性列表"`
}

type DevicePropertiy struct {
	Key   string      `json:"key" dc:"属性标识"`
	Name  string      `json:"name" dc:"属性名称"`
	Type  string      `json:"type" dc:"属性值类型"`
	Unit  string      `json:"unit" dc:"属性值单位"`
	Value *gvar.Var   `json:"value" dc:"属性值"`
	List  []*gvar.Var `json:"list" dc:"当天属性值列表"`
}

type DeviceLatestProperty struct {
	Key   string    `json:"key" dc:"属性标识"`
	Name  string    `json:"name" dc:"属性名称"`
	Type  string    `json:"type" dc:"属性值类型"`
	Unit  string    `json:"unit" dc:"属性值单位"`
	Value *gvar.Var `json:"value" dc:"属性值"`
}

type DevicePropertiyRes struct {
	Ts    *gtime.Time `json:"ts" dc:"时间"`
	Value *gvar.Var   `json:"value" dc:"属性值"`
}

type DevicePropertiyOut struct {
	Ts    *gtime.Time `json:"ts" dc:"时间"`
	Value *gvar.Var   `json:"value" dc:"属性值"`
}

type DeviceGetPropertyInput struct {
	DeviceKey   string `json:"device_key" dc:"设备ID" v:"required#设备key不能为空"`
	PropertyKey string `json:"propertyKey" dc:"属性标识" v:"required#属性标识不能为空"`
}

type DeviceGetPropertyListInput struct {
	DeviceKey   string `json:"device_key" dc:"设备ID" v:"required#设备key不能为空"`
	PropertyKey string `json:"propertyKey" dc:"属性标识" v:"required#属性标识不能为空"`
	PaginationInput
}
type DeviceGetPropertyListOutput struct {
	List []*DevicePropertiyOut
	PaginationOutput
}

type DeviceGetDataInput struct {
	DeviceKey   string   `json:"device_key" dc:"设备ID" v:"required#设备key不能为空"`
	PropertyKey string   `json:"propertyKey" dc:"属性标识" v:"required#属性标识不能为空"`
	DateRange   []string `json:"dateRange" dc:"日期范围"`
	IsDesc      int      `json:"isDesc" dc:"排序：0=顺序，1=倒序"`
}

type DeviceTotalOutput struct {
	DeviceTotal   int `json:"deviceTotal" dc:"设备总量"`
	DeviceOffline int `json:"deviceOffline" dc:"离线设备数量"`

	ProductTotal        int `json:"productTotal" dc:"产品总量"`
	ProductAdded        int `json:"productAdded" dc:"今日产品增量"`
	ProductActivation   int `json:"productActivation" dc:"产品启用"`
	ProductDeactivation int `json:"productDeactivation" dc:"产品停用"`

	MsgTotal int64 `json:"msgTotal" dc:"设备消息总量"`
	MsgAdded int64 `json:"msgAdded" dc:"今日设备消息增量"`

	AlarmTotal int `json:"alarmTotal" dc:"设备报警总量"`
	AlarmAdded int `json:"alarmAdded" dc:"今日设备报警增量"`
}

type DeviceBindInput struct {
	GatewayKey string   `json:"gatewayKey" dc:"网关标识" v:"required#网关标识不能为空"`
	SubKeys    []string `json:"subKeys" dc:"子设备标识列表"`
}

type DeviceBindListInput struct {
	GatewayKey string `json:"gatewayKey" dc:"网关标识" v:"required#网关标识不能为空"`
	PaginationInput
}

type DeviceBindListOutput struct {
	List []*DeviceOutput `json:"list" dc:"网关绑定子设备列表"`
	PaginationOutput
}

type CheckBindInput struct {
	GatewayKey string `json:"gatewayKey" dc:"网关标识" v:"required#网关标识不能为空"`
	SubKey     string `json:"subKey" dc:"子设备标识" v:"required#子设备标识不能为空"`
}

type AuthInfoInput struct {
	DeviceKey  string `json:"deviceKey" dc:"设备标识"`
	ProductKey string `json:"ProductKey" dc:"产品标识"`
}
type AuthInfoOutput struct {
	AuthType      int                    `json:"authType" dc:"认证方式（1=Basic，2=AccessToken，3=证书）"`
	AuthUser      string                 `json:"authUser" dc:"认证用户"`
	AuthPasswd    string                 `json:"authPasswd" dc:"认证密码"`
	AccessToken   string                 `json:"accessToken" dc:"AccessToken"`
	CertificateId int                    `json:"certificateId" dc:"证书ID"`
	Certificate   *entity.SysCertificate `json:"certificate" dc:"证书信息"`
}
type DeviceExport struct {
	ProductName string `json:"productName" dc:"产品名称"`
	DeviceName  string `json:"deviceName" dc:"设备名称"`
	DeviceKey   string `json:"deviceKey" dc:"设备标识"`
	DeviceType  string `json:"deviceType" dc:"设备类型"`
	Status      string `json:"status" dc:"状态"`
	Desc        string `json:"desc" dc:"说明" v:"max-length:200#描述长度不能超过200个字符"`
	Version     string `json:"version" dc:"固件版本号"`
}

type DeviceDataListInput struct {
	DeviceKey string `json:"deviceKey" dc:"设备标识" v:"required#请输入设备标识"`
	Interval  uint   `json:"interval" dc:"聚合时间间隔" v:"required|min:1#请输入时间间隔"`
	TimeUnit  uint   `json:"timeUnit" dc:"时间单位: 1=秒,2=分,3=小时,4=天" v:"in:1,2,3,4#时间单位不正确"`

	PaginationInput
}
type DeviceDataListOutput struct {
	List gdb.Result `json:"list" dc:"设备数据列表"`

	PaginationOutput
}

type DeviceExtensionInfoInput struct {
	Id            uint   `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
	ExtensionInfo string `json:"extensionInfo" dc:"设备扩展信息"`
}
