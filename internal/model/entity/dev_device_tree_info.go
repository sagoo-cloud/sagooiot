// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDeviceTreeInfo is the golang structure for table dev_device_tree_info.
type DevDeviceTreeInfo struct {
	Id        int         `json:"id"        description:""`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	Name      string      `json:"name"      description:"名称"`
	Code      string      `json:"code"      description:"编码"`
	DeviceKey string      `json:"deviceKey" description:"设备标识"`
	Company   string      `json:"company"   description:"所属公司"`
	Area      string      `json:"area"      description:"区域"`
	Address   string      `json:"address"   description:"地址"`
	Lng       string      `json:"lng"       description:"经度"`
	Lat       string      `json:"lat"       description:"纬度"`
	Contact   string      `json:"contact"   description:"联系人"`
	Phone     string      `json:"phone"     description:"联系电话"`
	StartDate *gtime.Time `json:"startDate" description:"服务周期：开始日期"`
	EndDate   *gtime.Time `json:"endDate"   description:"服务周期：截止日期"`
	Image     string      `json:"image"     description:"图片"`
	Duration  int         `json:"duration"  description:"时间窗口值"`
	TimeUnit  int         `json:"timeUnit"  description:"时间单位：1=秒，2=分钟，3=小时，4=天"`
	Template  string      `json:"template"  description:"页面模板，默认：default"`
	Category  string      `json:"category"  description:"分类"`
	Types     string      `json:"types"     description:"类型"`
	CreatedBy uint        `json:"createdBy" description:"创建者"`
	UpdatedBy uint        `json:"updatedBy" description:"更新者"`
	DeletedBy uint        `json:"deletedBy" description:"删除者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
