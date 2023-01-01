package model

import (
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
)

const (
	ProductStatusOff int = iota // 产品未发布
	ProductStatusOn             // 产品已发布
)

type ListForPageInput struct {
	ProductInput
	PaginationInput
}
type ListForPageOutput struct {
	Product []*ProductOutput `json:"product" dc:"产品列表"`
	PaginationOutput
}

type ProductInput struct {
	Name             string   `json:"name" dc:"产品名称" `
	CategoryId       uint     `json:"categoryId" dc:"所属品类"`
	MessageProtocols []string `json:"messageProtocol" dc:"消息协议"`
	DeviceTypes      []string `json:"deviceType" dc:"设备类型：网关、设备"`
	Status           string   `p:"status"` //产品状态
}

type ProductOutput struct {
	*entity.DevProduct
	DeviceTotal  int    `json:"deviceTotal" dc:"设备数量"`
	CategoryName string `json:"categoryName" dc:"分类名称"`

	Category *DevProductCategory `json:"category" orm:"with:id=category_id" dc:"分类信息"`
}

type DevProductCategory struct {
	Id   uint   `json:"id" dc:"分类ID"`
	Name string `json:"name" dc:"分类名称"`
}
type SysDept struct {
	DeptId   uint   `json:"deptId" dc:"部门ID"`
	DeptName string `json:"deptName" dc:"部门名称"`
}

type DetailProductOutput struct {
	*entity.DevProduct
	DeviceTotal  int    `json:"deviceTotal" dc:"设备数量"`
	CategoryName string `json:"categoryName" dc:"分类名称"`

	Category *DevProductCategory `json:"category" orm:"with:id=category_id" dc:"部门信息"`

	TSL *TSL `json:"tsl" dc:"物模型"`
}

type AddProductInput struct {
	Key               string `json:"key" dc:"产品标识" v:"required|regex:^[A-Za-z_]+[\\w]*$#请输入产品标识|标识由字母、数字和下划线组成,且不能以数字开头"`
	Name              string `json:"name" dc:"产品名称" v:"required#请输入产品名称"`
	CategoryId        uint   `json:"categoryId" dc:"所属品类" v:"required#请选择所属品类"`
	MessageProtocol   string `json:"messageProtocol" dc:"消息协议" v:"required#请选择消息协议"`
	TransportProtocol string `json:"transportProtocol" dc:"传输协议: MQTT,COAP,UDP" v:"required#请选择传输协议"`
	DeviceType        string `json:"deviceType" dc:"设备类型：网关、设备" v:"required#请选择设备类型"`
	Desc              string `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Icon              string `json:"icon" dc:"图片地址"`
}

type EditProductInput struct {
	Id                uint    `json:"id" dc:"产品ID" v:"required#产品ID不能为空"`
	Name              string  `json:"name" dc:"产品名称" v:"required#请输入产品名称"`
	CategoryId        uint    `json:"categoryId" dc:"所属品类" v:"required#请选择所属品类"`
	MessageProtocol   string  `json:"messageProtocol" dc:"消息协议" v:"required#请选择消息协议"`
	TransportProtocol string  `json:"transportProtocol" dc:"传输协议: MQTT,COAP,UDP" v:"required#请选择传输协议"`
	DeviceType        string  `json:"deviceType" dc:"设备类型：网关、设备" v:"required#请选择设备类型"`
	Desc              string  `json:"desc" dc:"描述" v:"max-length:200#描述长度不能超过200个字符"`
	Icon              *string `json:"icon" dc:"图片地址"`
}
