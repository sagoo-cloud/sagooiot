// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCertificate is the golang structure of table sys_certificate for DAO operations like Where/Data.
type SysCertificate struct {
	g.Meta            `orm:"table:sys_certificate, do:true"`
	Id                interface{} //
	DeptId            interface{} // 部门ID
	Name              interface{} // 名称
	Standard          interface{} // 证书标准
	FileContent       interface{} // 证书文件内容
	PublicKeyContent  interface{} // 证书公钥内容
	PrivateKeyContent interface{} // 证书私钥内容
	Description       interface{} // 说明
	Status            interface{} // 状态  0未启用  1启用
	IsDeleted         interface{} // 是否删除 0未删除 1已删除
	CreatedBy         interface{} // 创建者
	CreatedAt         *gtime.Time // 创建日期
	UpdatedBy         interface{} // 修改人
	UpdatedAt         *gtime.Time // 更新时间
	DeletedBy         interface{} // 删除人
	DeletedAt         *gtime.Time // 删除时间
}
