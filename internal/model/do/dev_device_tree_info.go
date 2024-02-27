// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DevDeviceTreeInfo is the golang structure of table dev_device_tree_info for DAO operations like Where/Data.
type DevDeviceTreeInfo struct {
	g.Meta    `orm:"table:dev_device_tree_info, do:true"`
	Id        interface{} //
	DeptId    interface{} // 部门ID
	Name      interface{} // 名称
	Code      interface{} // 编码
	DeviceKey interface{} // 设备标识
	Company   interface{} // 所属公司
	Area      interface{} // 区域
	Address   interface{} // 地址
	Lng       interface{} // 经度
	Lat       interface{} // 纬度
	Contact   interface{} // 联系人
	Phone     interface{} // 联系电话
	StartDate *gtime.Time // 服务周期：开始日期
	EndDate   *gtime.Time // 服务周期：截止日期
	Image     interface{} // 图片
	Duration  interface{} // 时间窗口值
	TimeUnit  interface{} // 时间单位：1=秒，2=分钟，3=小时，4=天
	Template  interface{} // 页面模板，默认：default
	Category  interface{} // 分类
	Types     interface{} // 类型
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	DeletedBy interface{} // 删除者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
