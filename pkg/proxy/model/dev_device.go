package model

import (
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gtime"
)

type DevDeviceTag struct {
	Id       uint   `json:"id" dc:"标签ID"`
	DeviceId uint   `json:"deviceId" dc:"设备ID"`
	Key      string `json:"key" dc:"标签标识"`
	Name     string `json:"name" dc:"标签名称"`
	Value    string `json:"value" dc:"标签值"`
}

// 物模型
type TSL struct {
	Key        string              `json:"key" dc:"产品标识"`
	Name       string              `json:"name" dc:"产品名称"`
	Properties []model.TSLProperty `json:"properties" dc:"属性"`
	Functions  []model.TSLFunction `json:"functions" dc:"功能"`
	Events     []model.TSLEvent    `json:"events" dc:"事件"`
	Tags       []model.TSLTag      `json:"tags" dc:"标签"`
}

// 参数值（类型、类型参数）
type TSLValueType struct {
	Type           string `json:"type" ` // 类型
	model.TSLParam        // 参数
}

type DeviceRes struct {
	*entity.DevDevice
	DeptName    string     `json:"deptName" dc:"部门名称"`
	ProductName string     `json:"productName" dc:"产品名称"`
	TSL         *model.TSL `json:"tsl" dc:"物模型"`

	Product *entity.DevProduct    `json:"product" orm:"with:key=product_key" dc:"产品信息"`
	Tags    []*model.DevDeviceTag `json:"tags" orm:"with:device_id=id" dc:"设备标签"`
}

type DeviceOutput struct {
	*entity.DevDevice
	DeptName    string     `json:"deptName" dc:"部门名称"`
	ProductName string     `json:"productName" dc:"产品名称"`
	TSL         *model.TSL `json:"tsl" dc:"物模型"`

	Product *entity.DevProduct    `json:"product" orm:"with:key=product_key" dc:"产品信息"`
	Tags    []*model.DevDeviceTag `json:"tags" orm:"with:device_id=id" dc:"设备标签"`
}

// 设备运行状态
type DeviceRunStatusOutput struct {
	Status         int                     `json:"status" dc:"状态：0=未启用,1=离线,2=在线"`
	LastOnlineTime *gtime.Time             `json:"lastOnlineTime" dc:"最后上线时间"`
	Properties     []model.DevicePropertiy `json:"properties" dc:"属性列表"`
}

// 属性
type TSLProperty struct {
	Key        string       `json:"key" dc:"属性标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入属性标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name       string       `json:"name" dc:"属性名称" v:"required#请输入属性名称"`
	AccessMode int          `json:"accessMode" dc:"属性访问类型:0=读写,1=只读" v:"required#请选择是否只读"`
	ValueType  TSLValueType `json:"valueType" dc:"属性值"`
	Desc       string       `json:"desc" dc:"描述"`
}

type ListDeviceForPageInput struct {
	Key         string   `json:"key" dc:"设备标识"`
	Name        string   `json:"name" dc:"设备名称"`
	ProductKey  string   `json:"productKey" dc:"所属产品"`
	TunnelId    int      `json:"tunnelId"       description:"tunnelId"`
	Status      string   `p:"status"` //设备状态
	Keys        []string `json:"keys" dc:"设备标识"`
	DeviceTypes []string `json:"deviceTypes" dc:"设备类型"`
	DeptIds     []int    `json:"deptIds" dc:"组织ID"`
	model.PaginationInput
}

type DeviceGetPropertyListInput struct {
	DeviceKey   string `json:"device_key" dc:"设备ID" v:"required#设备key不能为空"`
	PropertyKey string `json:"propertyKey" dc:"属性标识" v:"required#属性标识不能为空"`
	model.PaginationInput
}
type DeviceGetPropertyListOutput struct {
	List []*DevicePropertiyOut `json:"list" dc:"属性列表"`
	model.PaginationOutput
}
type DevicePropertiyOut struct {
	Ts    *gtime.Time `json:"ts" dc:"时间"`
	Value *gvar.Var   `json:"value" dc:"属性值"`
}
