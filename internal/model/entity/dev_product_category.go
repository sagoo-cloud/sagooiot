// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DevProductCategory is the golang structure for table dev_product_category.
type DevProductCategory struct {
	Id        uint        `json:"id"        description:""`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	ParentId  uint        `json:"parentId"  description:"父ID"`
	Key       string      `json:"key"       description:"分类标识"`
	Name      string      `json:"name"      description:"分类名称"`
	Sort      int         `json:"sort"      description:"排序"`
	Desc      string      `json:"desc"      description:"描述"`
	CreatedBy uint        `json:"createdBy" description:"创建者"`
	UpdatedBy uint        `json:"updatedBy" description:"更新者"`
	DeletedBy uint        `json:"deletedBy" description:"删除者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
