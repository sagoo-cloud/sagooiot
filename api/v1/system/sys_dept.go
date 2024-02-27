package system

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type DeptDoReq struct {
	g.Meta   `path:"/dept/tree" method:"get" summary:"获取部门列表" tags:"部门管理"`
	Status   int    `p:"status" description:"状态:-1为全部,0为正常,1为停用" `
	DeptName string `p:"dept_name" description:"部门名称"`
}
type DeptDoRes struct {
	Data []*model.DeptRes
}

type AddDeptReq struct {
	g.Meta         `path:"/dept/add" method:"post" summary:"添加部门" tags:"部门管理"`
	ParentId       int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	OrganizationId int    `json:"organizationId" description:"组织ID" v:"required#请输入选择组织"`
	DeptName       string `json:"deptName"  description:"部门名称" v:"required#请输入部门名称"`
	OrderNum       int    `json:"orderNum"  description:"排序"`
	Status         uint   `json:"status"    description:"部门状态（0停用 1正常）" v:"required#请选择状态"`
	Leader         string `json:"leader"    description:"负责人" v:"required#请输入部门负责人"`
	Phone          string `json:"phone"     description:"联系电话" v:"phone#请输入联系电话或格式错误"`
	Email          string `json:"email"     description:"邮箱"`
}
type AddDeptRes struct {
}

type EditDeptReq struct {
	g.Meta         `path:"/dept/edit" method:"put" summary:"编辑部门" tags:"部门管理"`
	DeptId         int64  `json:"deptId"    description:"部门id"`
	ParentId       int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	OrganizationId int    `json:"organizationId" description:"组织ID" v:"required#请输入选择组织"`
	DeptName       string `json:"deptName"  description:"部门名称" v:"required#请输入部门名称"`
	OrderNum       int    `json:"orderNum"  description:"排序"`
	Status         uint   `json:"status"    description:"部门状态（0停用 1正常）" v:"required#请选择状态"`
	Leader         string `json:"leader"    description:"负责人" v:"required#请输入部门负责人"`
	Phone          string `json:"phone"     description:"联系电话" v:"phone#请输入联系电话或格式错误"`
	Email          string `json:"email"     description:"邮箱"`
}
type EditDeptRes struct {
}

type DetailDeptReq struct {
	g.Meta `path:"/dept/detail" method:"get" summary:"根据ID获取部门详情" tags:"部门管理"`
	DeptId int64 `p:"dept_id" description:"部门ID"  v:"required#ID不能为空"`
}
type DetailDeptRes struct {
	Data *model.DetailDeptRes
}

type DelDeptReq struct {
	g.Meta `path:"/dept/del" method:"delete" summary:"根据ID删除部门" tags:"部门管理"`
	DeptId int64 `p:"dept_id" description:"部门ID"  v:"required#ID不能为空"`
}
type DelDeptRes struct {
}
