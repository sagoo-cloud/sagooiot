// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfig is the golang structure of table sys_config for DAO operations like Where/Data.
type SysConfig struct {
	g.Meta         `orm:"table:sys_config, do:true"`
	ConfigId       interface{} // 参数主键
	ModuleClassify interface{} // 所属字典类型数据code
	ConfigName     interface{} // 参数名称
	ConfigKey      interface{} // 参数键名
	ConfigValue    interface{} // 参数键值
	ConfigType     interface{} // 系统内置（1是 2否）
	Remark         interface{} // 备注
	Status         interface{} // 状态 0 停用 1启用
	IsDeleted      interface{} // 是否删除 0未删除 1已删除
	CreatedBy      interface{} // 创建者
	CreatedAt      *gtime.Time // 创建时间
	UpdatedBy      interface{} // 更新者
	UpdatedAt      *gtime.Time // 修改时间
	DeletedBy      interface{} // 删除人
	DeletedAt      *gtime.Time // 删除时间
}
