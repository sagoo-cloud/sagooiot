package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type MenuButtonDoReq struct {
	g.Meta   `path:"/menu/button/tree"    tags:"菜单管理" method:"get" summary:"菜单与按钮树列表"`
	MenuId   int    `json:"menuId"          dc:"菜单ID" v:"required#菜单ID不能为空"`
	ParentId int    `json:"parentId"        dc:"父ID"`
	Status   int    `json:"status"           dc:"状态:-1为全部,0为正常,1为停用"`
	Name     string `json:"name"             dc:"名称"`
}
type MenuButtonDoRes struct {
	Data []*model.UserMenuButtonRes
}

type AddMenuButtonReq struct {
	g.Meta      `path:"/menu/button/add"     tags:"菜单管理" method:"post" summary:"添加菜单与按钮相关关联"`
	ParentId    int    `json:"parentId"    description:"父ID" v:"required#请选择上级"`
	MenuId      int    `json:"menuId"      description:"菜单ID" v:"required#请选择菜单ID"`
	Name        string `json:"name"        description:"名称" v:"required#请输入名称"`
	Types       string `json:"types"       description:"类型 自定义 add添加 edit编辑 del 删除" v:"required#请选择类型"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}
type AddMenuButtonRes struct {
}

type DetailMenuButtonReq struct {
	g.Meta `path:"/menu/button/detail" tags:"菜单管理" method:"get" summary:"根据ID获取菜单按钮详情"`
	Id     int64 `p:"id" description:"菜单按钮ID"  v:"required#ID不能为空"`
}
type DetailMenuButtonRes struct {
	Data *model.DetailMenuButtonRes
}

type EditMenuButtonReq struct {
	g.Meta      `path:"/menu/button/edit" method:"put" summary:"编辑菜单按钮" tags:"菜单管理"`
	Id          int    `json:"id"          description:"" v:"required#ID不能为空"`
	ParentId    int    `json:"parentId"    description:"父ID" v:"required#请选择上级"`
	MenuId      int    `json:"menuId"      description:"菜单ID" v:"required#请选择关联菜单"`
	Name        string `json:"name"        description:"名称" v:"required#请输入名称"`
	Types       string `json:"types"       description:"类型 自定义 add添加 edit编辑 del 删除" v:"required#请选择类型"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}
type EditMenuButtonRes struct {
}

type DelMenuButtonReq struct {
	g.Meta `path:"/menu/button/del" method:"delete" summary:"根据ID删除菜单按钮" tags:"菜单管理"`
	Id     int64 `json:"id" description:"菜单按钮ID"  v:"required#ID不能为空"`
}
type DelMenuButtonRes struct {
}

type EditMenuButtonStatusReq struct {
	g.Meta `path:"/menu/button/editStatus" method:"put" summary:"编辑菜单按钮状态" tags:"菜单管理"`
	Id     int `json:"id"          description:"" v:"required#ID不能为空"`
	MenuId int `json:"menuId"      description:"菜单ID" v:"required#请选择关联菜单"`
	Status int `json:"status"      description:"状态 0 停用 1启用"`
}

type EditMenuButtonStatusRes struct {
}
