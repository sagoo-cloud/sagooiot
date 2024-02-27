package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

// GetCertificateListReq 获取数据列表
type GetCertificateListReq struct {
	g.Meta `path:"/certificate/list" method:"get" summary:"获取列表" tags:"证书管理"`
	Name   string `json:"name"        description:"名称"`
	Status int    `json:"status"      description:"状态 0 停用 1启用"`
	common.PaginationReq
}
type GetCertificateListRes struct {
	Info []model.SysCertificateListRes
	common.PaginationRes
}

// GetCertificateByIdReq 获取指定ID的数据
type GetCertificateByIdReq struct {
	g.Meta `path:"/certificate/getById" method:"get" summary:"根据ID获取数据" tags:"证书管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetCertificateByIdRes struct {
	Info *model.SysCertificateOut
}

// AddCertificateReq 添加数据
type AddCertificateReq struct {
	g.Meta            `path:"/certificate/add" method:"post" summary:"添加证书" tags:"证书管理"`
	Name              string `json:"name"              description:"名称" v:"required#名称不能为空"`
	Standard          string `json:"standard"          description:"证书标准" v:"required#证书标准不能为空"`
	FileContent       string `json:"fileContent"       description:"证书文件内容" v:"required#证书内容不能为空"`
	PublicKeyContent  string `json:"publicKeyContent"  description:"证书公钥内容" v:"required#证书内容不能为空"`
	PrivateKeyContent string `json:"privateKeyContent" description:"证书私钥内容" v:"required#证书私钥内容不能为空"`
	Description       string `json:"description"       description:"说明"`
}
type AddCertificateRes struct{}

// EditCertificateReq 编辑数据
type EditCertificateReq struct {
	g.Meta            `path:"/certificate/edit" method:"put" summary:"编辑证书" tags:"证书管理"`
	Id                int    `json:"id"                description:"" v:"required#id不能为空"`
	Name              string `json:"name"              description:"名称" v:"required#名称不能为空"`
	Standard          string `json:"standard"          description:"证书标准" v:"required#证书标准不能为空"`
	FileContent       string `json:"fileContent"       description:"证书文件内容" v:"required#证书内容不能为空"`
	PublicKeyContent  string `json:"publicKeyContent"  description:"证书公钥内容" v:"required#证书内容不能为空"`
	PrivateKeyContent string `json:"privateKeyContent" description:"证书私钥内容" v:"required#证书私钥内容不能为空"`
	Description       string `json:"description"       description:"说明"`
}
type EditCertificateRes struct{}

// DeleteCertificateReq 删除数据
type DeleteCertificateReq struct {
	g.Meta `path:"/certificate/delete" method:"delete" summary:"删除证书" tags:"证书管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type DeleteCertificateRes struct{}

// EditCertificateStatusReq 更新状态
type EditCertificateStatusReq struct {
	g.Meta `path:"/certificate/editStatus" method:"post" summary:"更新证书状态" tags:"证书管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
	Status int `json:"status"          description:"状态" v:"required#状态不能为空"`
}
type EditCertificateStatusRes struct{}

// GetCertificateAllReq 获取所有证书
type GetCertificateAllReq struct {
	g.Meta `path:"/certificate/getAll" method:"get" summary:"获取所有证书" tags:"证书管理"`
}
type GetCertificateAllRes struct {
	Info []*model.SysCertificateOut
}
