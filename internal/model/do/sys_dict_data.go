// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure of table sys_dict_data for DAO operations like Where/Data.
type SysDictData struct {
	g.Meta    `orm:"table:sys_dict_data, do:true"`
	DictCode  interface{} // 字典编码
	DictSort  interface{} // 字典排序
	DictLabel interface{} // 字典标签
	DictValue interface{} // 字典键值
	DictType  interface{} // 字典类型
	CssClass  interface{} // 样式属性（其他样式扩展）
	ListClass interface{} // 表格回显样式
	IsDefault interface{} // 是否默认（1是 0否）
	Remark    interface{} // 备注
	Status    interface{} // 状态（0正常 1停用）
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy interface{} // 更新者
	UpdatedAt *gtime.Time // 修改时间
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}
