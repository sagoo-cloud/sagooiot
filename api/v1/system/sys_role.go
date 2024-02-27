package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type RoleTreeReq struct {
	g.Meta `path:"/role/tree" method:"get" summary:"角色树状列表" tags:"角色管理"`
	Name   string `json:"name"      description:"角色名称"`
	Status int    `json:"status"    description:"状态;0:禁用;1:正常"`
}
type RoleTreeRes struct {
	Data []*model.RoleTreeRes
}

type AddRoleReq struct {
	g.Meta    `path:"/role/add" method:"post" summary:"添加角色" tags:"角色管理"`
	ParentId  int    `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	Name      string `json:"name"      description:"角色名称" v:"required#请输入名称"`
	ListOrder uint   `json:"listOrder" description:"排序"`
	Status    uint   `json:"status"    description:"状态;0:禁用;1:正常" v:"required#请选择状态"`
	Remark    string `json:"remark"    description:"备注"`
}
type AddRoleRes struct {
}

type EditRoleReq struct {
	g.Meta    `path:"/role/edit" method:"put" summary:"编辑角色" tags:"角色管理"`
	Id        uint   `json:"id"        description:"ID" v:"required#ID不能为空"`
	ParentId  int    `json:"parentId"  description:"父ID" v:"required#请输入选择上级"`
	Name      string `json:"name"      description:"角色名称" v:"required#请输入名称"`
	ListOrder uint   `json:"listOrder" description:"排序"`
	Status    uint   `json:"status"    description:"状态;0:禁用;1:正常" v:"required#请选择状态"`
	Remark    string `json:"remark"    description:"备注"`
}
type EditRoleRes struct {
}

type GetRoleByIdReq struct {
	g.Meta `path:"/role/getInfoById" method:"get" summary:"根据ID获取角色" tags:"角色管理"`
	Id     uint `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type GetRoleByIdRes struct {
	Data *model.RoleInfoRes
}

type DeleteRoleByIdReq struct {
	g.Meta `path:"/role/delInfoById" method:"delete" summary:"根据ID删除角色" tags:"角色管理"`
	Id     uint `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type DeleteRoleByIdRes struct {
}

type DataScopeReq struct {
	g.Meta    `path:"/role/dataScope" method:"post" summary:"角色数据权限授权" tags:"角色管理"`
	Id        int     `json:"id"        description:"ID" v:"required#ID不能为空"`
	DataScope uint    `json:"dataScope"        description:"数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）" v:"required#数据权限范围不能为空不能为空"`
	DeptIds   []int64 `json:"deptIds"`
}

type DataScopeRes struct {
}

type GetAuthorizeByIdReq struct {
	g.Meta `path:"/role/getAuthorizeById" method:"get" summary:"根据ID获取权限信息" tags:"角色管理"`
	Id     int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type GetAuthorizeByIdRes struct {
	MenuIds   []string `json:"menuIds" description:"菜单ID"`
	ButtonIds []string `json:"buttonIds" description:"按钮ID"`
	ColumnIds []string `json:"columnIds" description:"列表字段ID"`
	ApiIds    []string `json:"apiIds" description:"接口Ids"`
}
