// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysApi is the golang structure of table sys_api for DAO operations like Where/Data.
type SysApi struct {
	g.Meta    `orm:"table:sys_api, do:true"`
	Id        interface{} //
	ParentId  interface{} //
	Name      interface{} // 名称
	Types     interface{} // 1 分类 2接口
	ApiTypes  interface{} // 数据字典维护
	Method    interface{} // 请求方式(数据字典维护)
	Address   interface{} // 接口地址
	Remark    interface{} // 备注
	Status    interface{} // 状态 0 停用 1启用
	Sort      interface{} // 排序
	IsDeleted interface{} // 是否删除 0未删除 1已删除
	CreatedBy interface{} // 创建者
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy interface{} // 更新者
	UpdatedAt *gtime.Time // 修改时间
	DeletedBy interface{} // 删除人
	DeletedAt *gtime.Time // 删除时间
}
