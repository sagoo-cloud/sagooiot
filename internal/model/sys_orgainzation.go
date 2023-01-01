package model

import "github.com/gogf/gf/v2/os/gtime"

type OrganizationRes struct {
	Id        int64              `json:"id"    description:"组织id"`
	ParentId  int64              `json:"parentId"  description:"父组织id"`
	Ancestors string             `json:"ancestors" description:"祖级列表"`
	Name      string             `json:"name"  description:"组织名称"`
	Number    string             `json:"number"    description:"组织编号"`
	OrderNum  int                `json:"orderNum"  description:"显示顺序"`
	Leader    string             `json:"leader"    description:"负责人"`
	Phone     string             `json:"phone"     description:"联系电话"`
	Email     string             `json:"email"     description:"邮箱"`
	Status    uint               `json:"status"    description:"部门状态（0停用 1正常）"`
	IsDeleted int                `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt *gtime.Time        `json:"createdAt" description:"创建时间"`
	Children  []*OrganizationRes `json:"children" description:"子集"`
}

type OrganizationOut struct {
	Id        int64              `json:"id"    description:"组织id"`
	ParentId  int64              `json:"parentId"  description:"父组织id"`
	Ancestors string             `json:"ancestors" description:"祖级列表"`
	Name      string             `json:"name"  description:"组织名称"`
	Number    string             `json:"number"    description:"组织编号"`
	OrderNum  int                `json:"orderNum"  description:"显示顺序"`
	Leader    string             `json:"leader"    description:"负责人"`
	Phone     string             `json:"phone"     description:"联系电话"`
	Email     string             `json:"email"     description:"邮箱"`
	Status    uint               `json:"status"    description:"部门状态（0停用 1正常）"`
	IsDeleted int                `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt *gtime.Time        `json:"createdAt" description:"创建时间"`
	Children  []*OrganizationOut `json:"children" description:"子集"`
}

type AddOrganizationInput struct {
	ParentId int64  `json:"parentId"  description:"父组织id"`
	Name     string `json:"name"  description:"组织名称"`
	OrderNum int    `json:"orderNum"  description:"排序"`
	Status   uint   `json:"status"    description:"部门状态（0停用 1正常）"`
	Leader   string `json:"leader"    description:"负责人"`
	Phone    string `json:"phone"     description:"联系电话"`
	Email    string `json:"email"     description:"邮箱"`
}

type EditOrganizationInput struct {
	Id       int64  `json:"id"    description:"组织id"`
	ParentId int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	Name     string `json:"name"  description:"组织名称"`
	OrderNum int    `json:"orderNum"  description:"排序"`
	Status   uint   `json:"status"    description:"部门状态（0停用 1正常）"`
	Leader   string `json:"leader"    description:"负责人"`
	Phone    string `json:"phone"     description:"联系电话"`
	Email    string `json:"email"     description:"邮箱"`
}

type DetailOrganizationRes struct {
	Id             int64       `json:"id"    description:"组织id"`
	ParentId       int64       `json:"parentId"  description:"父部门id"`
	OrganizationId int         `json:"organizationId" description:"组织ID"`
	Ancestors      string      `json:"ancestors" description:"祖级列表"`
	Name           string      `json:"name"  description:"组织名称"`
	Number         string      `json:"number"    description:"组织编号"`
	OrderNum       int         `json:"orderNum"  description:"显示顺序"`
	Leader         string      `json:"leader"    description:"负责人"`
	Phone          string      `json:"phone"     description:"联系电话"`
	Email          string      `json:"email"     description:"邮箱"`
	Status         uint        `json:"status"    description:"部门状态（0停用 1正常）"`
	IsDeleted      int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt      *gtime.Time `json:"createdAt" description:"创建时间"`
}
