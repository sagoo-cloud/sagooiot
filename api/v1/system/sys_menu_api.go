package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type MenuApiDoReq struct {
	g.Meta `path:"/menu/api/tree"    tags:"菜单管理" method:"get" summary:"菜单与API列表"`
	MenuId int `json:"menuId"          dc:"菜单ID" v:"required#菜单ID不能为空"`
}
type MenuApiDoRes struct {
	Data []*model.SysApiAllRes
}

type AddMenuApiReq struct {
	g.Meta `path:"/menu/api/add"    tags:"菜单管理" method:"post" summary:"绑定菜单和API关联关系"`
	MenuId int   `json:"menuId"          dc:"菜单ID" v:"required#菜单ID不能为空"`
	ApiIds []int `json:"apiIds"          dc:"API ID"`
}
type AddMenuApiRes struct {
}
