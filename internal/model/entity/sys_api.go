// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysApi is the golang structure for table sys_api.
type SysApi struct {
	Id        uint        `json:"id"        description:""`
	ParentId  int         `json:"parentId"  description:""`
	Name      string      `json:"name"      description:"名称"`
	Types     int         `json:"types"     description:"1 分类 2接口"`
	ApiTypes  string      `json:"apiTypes"  description:"数据字典维护"`
	Method    string      `json:"method"    description:"请求方式(数据字典维护)"`
	Address   string      `json:"address"   description:"接口地址"`
	Remark    string      `json:"remark"    description:"备注"`
	Status    int         `json:"status"    description:"状态 0 停用 1启用"`
	Sort      int         `json:"sort"      description:"排序"`
	IsDeleted int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedBy uint        `json:"createdBy" description:"创建者"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedBy uint        `json:"updatedBy" description:"更新者"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"修改时间"`
	DeletedBy int         `json:"deletedBy" description:"删除人"`
	DeletedAt *gtime.Time `json:"deletedAt" description:"删除时间"`
}
