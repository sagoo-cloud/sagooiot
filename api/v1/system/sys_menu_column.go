package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type MenuColumnDoReq struct {
	g.Meta   `path:"/menu/column/tree"    tags:"菜单管理" method:"get" summary:"菜单与列表树列表"`
	MenuId   string `json:"menuId"          dc:"菜单ID" v:"required#菜单ID不能为空"`
	ParentId string `json:"parentId"        dc:"父ID"`
	Status   int    `json:"status"           dc:"状态:-1为全部,0为正常,1为停用"`
	Name     string `json:"name"             dc:"名称"`
}
type MenuColumnDoRes struct {
	Data []*model.UserMenuColumnRes
}

type AddMenuColumnReq struct {
	g.Meta      `path:"/menu/column/add"     tags:"菜单管理" method:"post" summary:"添加菜单与列表相关关联"`
	ParentId    int    `json:"parentId"    description:"父ID" v:"required#请选择上级"`
	MenuId      int    `json:"menuId"      description:"菜单ID" v:"required#请选择菜单ID"`
	Name        string `json:"name"        description:"名称" v:"required#请输入名称"`
	Code        string `json:"code"        description:"代表列表"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}
type AddMenuColumnRes struct {
}

type DetailMenuColumnReq struct {
	g.Meta `path:"/menu/column/detail" tags:"菜单管理" method:"get" summary:"根据ID获取菜单列表详情"`
	Id     int64 `p:"id" description:"菜单列表ID"  v:"required#ID不能为空"`
}
type DetailMenuColumnRes struct {
	Data *model.DetailMenuColumnRes
}

type EditMenuColumnReq struct {
	g.Meta      `path:"/menu/column/edit" method:"put" summary:"编辑菜单列表" tags:"菜单管理"`
	Id          int    `json:"id"          description:"" v:"required#ID不能为空"`
	ParentId    int    `json:"parentId"    description:"父ID" v:"required#请选择上级"`
	MenuId      int    `json:"menuId"      description:"菜单ID" v:"required#请选择关联菜单"`
	Name        string `json:"name"        description:"名称" v:"required#请输入名称"`
	Code        string `json:"code"        description:"代表列表"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}
type EditMenuColumnRes struct {
}

type DelMenuColumnReq struct {
	g.Meta `path:"/menu/column/del" method:"delete" summary:"根据ID删除菜单列表" tags:"菜单管理"`
	Id     int64 `p:"id" description:"菜单列表ID"  v:"required#ID不能为空"`
}
type DelMenuColumnRes struct {
}

type EditMenuColumnStatusReq struct {
	g.Meta `path:"/menu/column/editStatus" method:"put" summary:"编辑菜单列表状态" tags:"菜单管理"`
	Id     int `json:"id"          description:"" v:"required#ID不能为空"`
	MenuId int `json:"menuId"      description:"菜单ID" v:"required#请选择关联菜单"`
	Status int `json:"status"      description:"状态 0 停用 1启用"`
}

type EditMenuColumnStatusRes struct {
}
