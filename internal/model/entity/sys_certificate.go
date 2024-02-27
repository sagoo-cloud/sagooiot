// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCertificate is the golang structure for table sys_certificate.
type SysCertificate struct {
	Id                int         `json:"id"                description:""`
	DeptId            int         `json:"deptId"            description:"部门ID"`
	Name              string      `json:"name"              description:"名称"`
	Standard          string      `json:"standard"          description:"证书标准"`
	FileContent       string      `json:"fileContent"       description:"证书文件内容"`
	PublicKeyContent  string      `json:"publicKeyContent"  description:"证书公钥内容"`
	PrivateKeyContent string      `json:"privateKeyContent" description:"证书私钥内容"`
	Description       string      `json:"description"       description:"说明"`
	Status            int         `json:"status"            description:"状态  0未启用  1启用"`
	IsDeleted         int         `json:"isDeleted"         description:"是否删除 0未删除 1已删除"`
	CreatedBy         uint        `json:"createdBy"         description:"创建者"`
	CreatedAt         *gtime.Time `json:"createdAt"         description:"创建日期"`
	UpdatedBy         int         `json:"updatedBy"         description:"修改人"`
	UpdatedAt         *gtime.Time `json:"updatedAt"         description:"更新时间"`
	DeletedBy         int         `json:"deletedBy"         description:"删除人"`
	DeletedAt         *gtime.Time `json:"deletedAt"         description:"删除时间"`
}
