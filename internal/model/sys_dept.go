package model

import "github.com/gogf/gf/v2/os/gtime"

type DeptRes struct {
	DeptId         int64       `json:"deptId"    description:"部门id"`
	OrganizationId int         `json:"organizationId" description:"组织ID"`
	ParentId       int64       `json:"parentId"  description:"父部门id"`
	Ancestors      string      `json:"ancestors" description:"祖级列表"`
	DeptName       string      `json:"deptName"  description:"部门名称"`
	OrderNum       int         `json:"orderNum"  description:"显示顺序"`
	Leader         string      `json:"leader"    description:"负责人"`
	Phone          string      `json:"phone"     description:"联系电话"`
	Email          string      `json:"email"     description:"邮箱"`
	Status         uint        `json:"status"    description:"部门状态（0停用 1正常）"`
	IsDeleted      int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt      *gtime.Time `json:"createdAt" description:"创建时间"`
	Children       []*DeptRes  `json:"children" description:"子集"`
}

type DeptOut struct {
	DeptId         int64       `json:"deptId"    description:"部门id"`
	OrganizationId int         `json:"organizationId" description:"组织ID"`
	ParentId       int64       `json:"parentId"  description:"父部门id"`
	Ancestors      string      `json:"ancestors" description:"祖级列表"`
	DeptName       string      `json:"deptName"  description:"部门名称"`
	OrderNum       int         `json:"orderNum"  description:"显示顺序"`
	Leader         string      `json:"leader"    description:"负责人"`
	Phone          string      `json:"phone"     description:"联系电话"`
	Email          string      `json:"email"     description:"邮箱"`
	Status         uint        `json:"status"    description:"部门状态（0停用 1正常）"`
	IsDeleted      int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt      *gtime.Time `json:"createdAt" description:"创建时间"`
	Children       []*DeptOut  `json:"children" description:"子集"`
}

type AddDeptInput struct {
	ParentId       int64  `json:"parentId"`
	OrganizationId int    `json:"organizationId"`
	DeptName       string `json:"deptName"`
	OrderNum       int    `json:"orderNum"`
	Status         uint   `json:"status"`
	Leader         string `json:"leader"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}

type EditDeptInput struct {
	DeptId         int64  `json:"deptId"`
	ParentId       int64  `json:"parentId"`
	OrganizationId int    `json:"organizationId"`
	DeptName       string `json:"deptName"`
	OrderNum       int    `json:"orderNum"`
	Status         uint   `json:"status"`
	Leader         string `json:"leader"`
	Phone          string `json:"phone"`
	Email          string `json:"email"`
}

type DetailDeptRes struct {
	DeptId         int64       `json:"deptId"    description:"部门id"`
	ParentId       int64       `json:"parentId"  description:"父部门id"`
	OrganizationId int         `json:"organizationId" description:"组织ID"`
	Ancestors      string      `json:"ancestors" description:"祖级列表"`
	DeptName       string      `json:"deptName"  description:"部门名称"`
	OrderNum       int         `json:"orderNum"  description:"显示顺序"`
	Leader         string      `json:"leader"    description:"负责人"`
	Phone          string      `json:"phone"     description:"联系电话"`
	Email          string      `json:"email"     description:"邮箱"`
	Status         uint        `json:"status"    description:"部门状态（0停用 1正常）"`
	IsDeleted      int         `json:"isDeleted" description:"是否删除 0未删除 1已删除"`
	CreatedAt      *gtime.Time `json:"createdAt" description:"创建时间"`
}
