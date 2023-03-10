// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BaseDbLink is the golang structure for table base_db_link.
type BaseDbLink struct {
	Id          int         `json:"id"          description:""`
	Name        string      `json:"name"        description:"名称"`
	Types       string      `json:"types"       description:"驱动类型 mysql或oracle"`
	Host        string      `json:"host"        description:"主机地址"`
	Port        int         `json:"port"        description:"端口号"`
	UserName    string      `json:"userName"    description:"用户名称"`
	Password    string      `json:"password"    description:"密码"`
	Description string      `json:"description" description:"描述"`
	Status      int         `json:"status"      description:"状态 0 停用 1启用"`
	IsDeleted   int         `json:"isDeleted"   description:"是否删除 0未删除 1已删除"`
	CreatedBy   uint        `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedBy   int         `json:"updatedBy"   description:"修改人"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"更新时间"`
	DeletedBy   int         `json:"deletedBy"   description:"删除人"`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:"删除时间"`
}
