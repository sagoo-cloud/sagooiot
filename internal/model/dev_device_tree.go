package model

import (
	"sagooiot/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

type AddDeviceTreeInfoInput struct {
	Name      string      `json:"name" dc:"名称" v:"required#请输入名称"`
	Address   string      `json:"address" dc:"地址" v:"required#请输入地址"`
	Lng       string      `json:"lng" dc:"经度"`
	Lat       string      `json:"lat" dc:"纬度"`
	Contact   string      `json:"contact" dc:"联系人"`
	Phone     string      `json:"phone" dc:"联系电话"`
	StartDate *gtime.Time `json:"startDate" dc:"服务周期：开始日期"`
	EndDate   *gtime.Time `json:"endDate" dc:"服务周期：截止日期"`
	Image     string      `json:"image" dc:"图片"`
	DeviceKey string      `json:"deviceKey" dc:"设备标识"`
	Area      string      `json:"area" dc:"区域"`
	Company   string      `json:"company" dc:"所属公司"`
	ParentId  int         `json:"parentId" dc:"父ID"`
	Duration  int         `json:"duration" dc:"时间窗口值"`
	TimeUnit  int         `json:"timeUnit" dc:"时间单位：1=秒，2=分钟，3=小时，4=天"`
	Template  string      `json:"template" dc:"页面模板，默认：default" d:"default"`
	Category  string      `json:"category" dc:"分类"`
	Types     string      `json:"types" dc:"类型"`
}

type EditDeviceTreeInfoInput struct {
	Id        int         `json:"id" dc:"信息ID" v:"required#信息ID不能为空"`
	Name      string      `json:"name" dc:"名称" v:"required#请输入名称"`
	Address   string      `json:"address" dc:"地址" v:"required#请输入地址"`
	Lng       string      `json:"lng" dc:"经度"`
	Lat       string      `json:"lat" dc:"纬度"`
	Contact   string      `json:"contact" dc:"联系人"`
	Phone     string      `json:"phone" dc:"联系电话"`
	StartDate *gtime.Time `json:"startDate" dc:"服务周期：开始日期"`
	EndDate   *gtime.Time `json:"endDate" dc:"服务周期：截止日期"`
	Image     string      `json:"image" dc:"图片"`
	DeviceKey string      `json:"deviceKey" dc:"设备标识"`
	Area      string      `json:"area" dc:"区域"`
	Company   string      `json:"company" dc:"所属公司"`
	ParentId  int         `json:"parentId" dc:"父ID"`
	Duration  int         `json:"duration" dc:"时间窗口值"`
	TimeUnit  int         `json:"timeUnit" dc:"时间单位：1=秒，2=分钟，3=小时，4=天"`
	Template  string      `json:"template" dc:"页面模板，默认：default" d:"default"`
	Category  string      `json:"category" dc:"分类"`
	Types     string      `json:"types" dc:"类型"`
}

type DetailDeviceTreeInfoOutput struct {
	entity.DevDeviceTreeInfo
	ParentId int `json:"parentId" dc:"父信息ID"`
}

type (
	DeviceTreeListOutput struct {
		*DeviceTree
		Children []*DeviceTreeListOutput `json:"children" dc:"子集"`
	}
	DeviceTree struct {
		Id           int    `json:"id" dc:"设备树ID"`
		InfoId       int    `json:"infoId" dc:"信息ID"`
		ParentInfoId int    `json:"parentInfoId" dc:"父信息ID"`
		Name         string `json:"name" dc:"名称"`
	}
)
