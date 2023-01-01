package model

import (
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gmeta"
)

const (
	DeviceStatusNoEnable int = iota // 设备未启用
	DeviceStatusOff                 // 设备离线
	DeviceStatusOn                  // 设备在线
)

type DeviceInput struct {
	Key       string `json:"key" dc:"设备标识" v:"regex:^[A-Za-z_]+[\\w]*$#请输入设备标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name      string `json:"name" dc:"设备名称"`
	ProductId uint   `json:"productId" dc:"所属产品"`
	TunnelId  int    `json:"tunnelId"       description:"tunnelId"`
	Status    string `p:"status"` //设备状态
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
	Id       uint   `json:"id" dc:"产品ID"`
	Name     string `json:"name" dc:"产品名称"`
	Key      string `json:"key" dc:"产品标识"`
	Metadata string `json:"metadata" dc:"物模型"`
}
type DevDeviceTag struct {
	Id    uint   `json:"id" dc:"标签ID"`
	Key   string `json:"key" dc:"标签标识"`
	Name  string `json:"name" dc:"标签名称"`
	Value string `json:"value" dc:"标签值"`
}

type DeviceOutput struct {
	*entity.DevDevice
	ProductName string `json:"productName" dc:"产品名称"`
	TSL         *TSL   `json:"tsl" dc:"物模型"`

	Product *DevProduct     `json:"product" orm:"with:id=product_id" dc:"产品信息"`
	Tags    []*DevDeviceTag `json:"tags" orm:"with:device_id=id" dc:"设备标签"`
}

type AddDeviceInput struct {
	Key         string `json:"key" dc:"设备标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入设备标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name        string `json:"name" dc:"设备名称" v:"required#请输入设备名称"`
	ProductId   uint   `json:"productId" dc:"所属产品" v:"required#请选择所属产品"`
	Desc        string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Certificate string `json:"certificate" dc:"设备证书"`
	SecureKey   string `json:"secureKey" dc:"设备密钥"`
	Version     string `json:"version" dc:"固件版本号"`
}

type EditDeviceInput struct {
	Id          uint   `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
	Name        string `json:"name" dc:"设备名称" v:"required#请输入设备名称"`
	Desc        string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Certificate string `json:"certificate" dc:"设备证书"`
	SecureKey   string `json:"secureKey" dc:"设备密钥"`
	Version     string `json:"version" dc:"固件版本号"`
}

type ListDeviceInput struct {
	ProductId uint `json:"productId" dc:"产品ID"`
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

type DevicePropertiyOut struct {
	Ts    *gtime.Time `json:"ts" dc:"时间"`
	Value *gvar.Var   `json:"value" dc:"属性值"`
}

type DeviceGetPropertyInput struct {
	Id          uint   `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
	PropertyKey string `json:"propertyKey" dc:"属性标识" v:"required#属性标识不能为空"`
}

type DeviceGetPropertyListInput struct {
	Id          uint   `json:"id" dc:"设备ID" v:"required#设备ID不能为空"`
	PropertyKey string `json:"propertyKey" dc:"属性标识" v:"required#属性标识不能为空"`
	PaginationInput
}
type DeviceGetPropertyListOutput struct {
	List []*DevicePropertiyOut
	PaginationOutput
}

type DeviceTotalOutput struct {
	DeviceTotal   int `json:"deviceTotal" dc:"设备总量"`
	DeviceOffline int `json:"deviceOffline" dc:"离线设备数量"`

	ProductTotal int `json:"productTotal" dc:"产品总量"`
	ProductAdded int `json:"productAdded" dc:"今日产品增量"`

	MsgTotal int `json:"msgTotal" dc:"设备消息总量"`
	MsgAdded int `json:"msgAdded" dc:"今日设备消息增量"`

	AlarmTotal int `json:"alarmTotal" dc:"设备报警总量"`
	AlarmAdded int `json:"alarmAdded" dc:"今日设备报警增量"`
}
