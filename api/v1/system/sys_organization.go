package system

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type OrganizationDoReq struct {
	g.Meta `path:"/organization/tree" method:"get" summary:"获取组织列表" tags:"组织管理"`
	Status int    `p:"status" description:"状态:-1为全部,0为正常,1为停用" `
	Name   string `p:"name" description:"组织名称"`
}
type OrganizationDoRes struct {
	Data []*model.OrganizationRes
}

type AddOrganizationReq struct {
	g.Meta   `path:"/organization/add" method:"post" summary:"添加组织" tags:"组织管理"`
	ParentId int64  `json:"parentId"  description:"父组织id"`
	Name     string `json:"name"  description:"组织名称"`
	OrderNum int    `json:"orderNum"  description:"排序"`
	Status   uint   `json:"status"    description:"部门状态（0停用 1正常）" v:"required#请选择状态"`
	Leader   string `json:"leader"    description:"负责人" v:"required#请输入部门负责人"`
	Phone    string `json:"phone"     description:"联系电话" v:"phone#请输入联系电话或格式错误"`
	Email    string `json:"email"     description:"邮箱"`
}
type AddOrganizationRes struct {
}

type EditOrganizationReq struct {
	g.Meta   `path:"/organization/edit" method:"put" summary:"编辑组织" tags:"组织管理"`
	Id       int64  `json:"id"    description:"组织id"`
	ParentId int64  `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	Name     string `json:"name"  description:"组织名称"`
	OrderNum int    `json:"orderNum"  description:"排序"`
	Status   uint   `json:"status"    description:"部门状态（0停用 1正常）" v:"required#请选择状态"`
	Leader   string `json:"leader"    description:"负责人" v:"required#请输入部门负责人"`
	Phone    string `json:"phone"     description:"联系电话" v:"phone#请输入联系电话或格式错误"`
	Email    string `json:"email"     description:"邮箱"`
}
type EditOrganizationRes struct {
}

type DetailOrganizationReq struct {
	g.Meta `path:"/organization/detail" method:"get" summary:"根据ID获取组织详情" tags:"组织管理"`
	Id     int64 `p:"id" description:"组织ID"  v:"required#ID不能为空"`
}
type DetailOrganizationRes struct {
	Data *model.DetailOrganizationRes
}

type DelOrganizationReq struct {
	g.Meta `path:"/organization/del" method:"delete" summary:"根据ID删除组织" tags:"组织管理"`
	Id     int64 `p:"id" description:"组织ID"  v:"required#ID不能为空"`
}
type DelOrganizationRes struct {
}

type GetOrganizationCountReq struct {
	g.Meta `path:"/organization/count" method:"get" summary:"获取组织数量" tags:"组织管理"`
}
type GetOrganizationCountRes struct {
	Count int
}
