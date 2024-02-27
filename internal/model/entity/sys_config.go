// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure for table sys_config.
type SysConfig struct {
	ConfigId       uint        `json:"configId"       description:"参数主键"`
	ModuleClassify string      `json:"moduleClassify" description:"所属字典类型数据code"`
	ConfigName     string      `json:"configName"     description:"参数名称"`
	ConfigKey      string      `json:"configKey"      description:"参数键名"`
	ConfigValue    string      `json:"configValue"    description:"参数键值"`
	ConfigType     int         `json:"configType"     description:"系统内置（1是 2否）"`
	Remark         string      `json:"remark"         description:"备注"`
	Status         int         `json:"status"         description:"状态 0 停用 1启用"`
	IsDeleted      int         `json:"isDeleted"      description:"是否删除 0未删除 1已删除"`
	CreatedBy      uint        `json:"createdBy"      description:"创建者"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	UpdatedBy      uint        `json:"updatedBy"      description:"更新者"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"修改时间"`
	DeletedBy      int         `json:"deletedBy"      description:"删除人"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
}
