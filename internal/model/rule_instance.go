package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type GetRuleInstanceListInput struct {
	Types int `json:"types"        description:"规则实例类型"`
	*PaginationInput
}

type RuleInstanceOut struct {
	Id        int         `json:"id"        description:"规则实例ID"`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	Name      string      `json:"name"      description:"规则实例名称"`
	Types     int         `json:"types"     description:"规则实例类型"`
	FlowId    string      `json:"flowId"    description:"流程ID"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Expound   string      `json:"expound"   description:"介绍"`
}

type RuleInstanceRes struct {
	Id        int         `json:"id"        description:"规则实例ID"`
	DeptId    int         `json:"deptId"    description:"部门ID"`
	Name      string      `json:"name"      description:"规则实例名称"`
	Types     int         `json:"types"     description:"规则实例类型"`
	FlowId    string      `json:"flowId"    description:"流程ID"`
	Status    int         `json:"status"    description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
	Expound   string      `json:"expound"   description:"介绍"`
}
type RuleInstanceAddInput struct {
	Name    string `json:"name"      description:"规则实例名称"`
	Types   int    `json:"types"     description:"规则实例类型"`
	Expound string `json:"expound"   description:"介绍"`
	FlowId  string `json:"flowId"    description:"流程ID"`
}
type RuleInstanceEditInput struct {
	Id int `json:"id"          description:"ID"`
	RuleInstanceAddInput
}
