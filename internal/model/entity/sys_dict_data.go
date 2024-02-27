// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
	DictCode  int64       `json:"dictCode"  description:"字典编码"`
	DictSort  int         `json:"dictSort"  description:"字典排序"`
	DictLabel string      `json:"dictLabel" description:"字典标签"`
	DictValue string      `json:"dictValue" description:"字典键值"`
	DictType  string      `json:"dictType"  description:"字典类型"`
	CssClass  string      `json:"cssClass"  description:"样式属性（其他样式扩展）"`
	ListClass string      `json:"listClass" description:"表格回显样式"`
	IsDefault int         `json:"isDefault" description:"是否默认（1是 0否）"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    int         `json:"status"    description:"状态（0正常 1停用）"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedBy uint        `json:"createdBy" description:"创建者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedBy uint        `json:"updatedBy" description:"更新者"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
