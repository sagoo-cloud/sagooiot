package model

import "github.com/gogf/gf/v2/os/gtime"

type SysDeptOut struct {
	DeptId         int64       `json:"deptId"         description:"部门id"`
	OrganizationId int         `json:"organizationId" description:"组织ID"`
	ParentId       int64       `json:"parentId"       description:"父部门id"`
	Ancestors      string      `json:"ancestors"      description:"祖级列表"`
	DeptName       string      `json:"deptName"       description:"部门名称"`
	OrderNum       int         `json:"orderNum"       description:"显示顺序"`
	Leader         string      `json:"leader"         description:"负责人"`
	Phone          string      `json:"phone"          description:"联系电话"`
	Email          string      `json:"email"          description:"邮箱"`
	Status         uint        `json:"status"         description:"部门状态（0停用 1正常）"`
	IsDeleted      int         `json:"isDeleted"      description:"是否删除 0未删除 1已删除"`
	CreatedAt      *gtime.Time `json:"createdAt"      description:"创建时间"`
	CreatedBy      uint        `json:"createdBy"      description:"创建人"`
	UpdatedBy      int         `json:"updatedBy"      description:"修改人"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      description:"修改时间"`
	DeletedBy      int         `json:"deletedBy"      description:"删除人"`
	DeletedAt      *gtime.Time `json:"deletedAt"      description:"删除时间"`
}

type AddDeptInput struct {
	OrganizationId int    `json:"organizationId" description:"组织ID"`
	ParentId       int64  `json:"parentId"       description:"父部门id"`
	DeptName       string `json:"deptName"       description:"部门名称"`
	OrderNum       int    `json:"orderNum"       description:"显示顺序"`
	Leader         string `json:"leader"         description:"负责人"`
	Phone          string `json:"phone"          description:"联系电话"`
	Email          string `json:"email"          description:"邮箱"`
	Status         uint   `json:"status"         description:"部门状态（0停用 1正常）"`
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
