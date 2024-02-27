// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DevProductCategory is the golang structure of table dev_product_category for DAO operations like Where/Data.
type DevProductCategory struct {
	g.Meta    `orm:"table:dev_product_category, do:true"`
	Id        interface{} //
	DeptId    interface{} // 部门ID
	ParentId  interface{} // 父ID
	Key       interface{} // 分类标识
	Name      interface{} // 分类名称
	Sort      interface{} // 排序
	Desc      interface{} // 描述
	CreatedBy interface{} // 创建者
	UpdatedBy interface{} // 更新者
	DeletedBy interface{} // 删除者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
