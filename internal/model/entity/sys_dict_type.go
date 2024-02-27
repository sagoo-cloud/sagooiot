// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictType is the golang structure for table sys_dict_type.
type SysDictType struct {
	DictId         uint64      `json:"dictId"         description:"字典主键"`
	ParentId       int         `json:"parentId"       description:"父主键ID"`
	DictName       string      `json:"dictName"       description:"字典名称"`
	DictType       string      `json:"dictType"       description:"字典类型"`
	ModuleClassify string      `json:"moduleClassify" description:"模块分类"`
	Remark         string      `json:"remark"         description:"备注"`
	Status         uint        `json:"status"         description:"状态（0正常 1停用）"`
	IsDeleted      int         `json:"isDeleted"      description:"是否删除 0未删除 1已删除"`
	CreatedBy      uint        `json:"createdBy"      description:"创建者"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建日期"`
	UpdatedBy      uint        `json:"updatedBy"      description:"更新者"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"修改日期"`
	DeletedBy      int         `json:"deletedBy"      description:"删除人"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
}
