package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type AuthorizeQueryReq struct {
	g.Meta    `path:"/authorize/query" method:"get" summary:"根据类型获取当前用户的权限" tags:"权限管理"`
	ItemsType string `json:"itemsType" description:"类型menu菜单 button按钮 column列表字段 api 接口" v:"required#项目类型不能为空" `
	MenuIds   []int  `json:"menuIds" description:"菜单ID"`
}
type AuthorizeQueryRes struct {
	Data []*model.AuthorizeQueryTreeRes
}

type AddAuthorizeReq struct {
	g.Meta    `path:"/authorize/Add" method:"post" summary:"添加权限" tags:"权限管理"`
	MenuIds   []string `json:"menuIds" description:"菜单ID" v:"required#菜单ID不能为空"`
	ButtonIds []string `json:"buttonIds" description:"按钮ID"`
	ColumnIds []string `json:"columnIds" description:"列表字段ID"`
	ApiIds    []string `json:"apiIds" description:"接口Ids"`
	RoleId    int      `json:"roleId" description:"角色ID" v:"required#角色ID不能为空"`
}
type AddAuthorizeRes struct {
}

type IsAllowAuthorizeReq struct {
	g.Meta `path:"/authorize/isAllow" method:"get" summary:"判断是否允许给角色授权" tags:"权限管理"`
	RoleId int `json:"roleId" description:"角色ID"  v:"required#角色ID不能为空"`
}
type IsAllowAuthorizeRes struct {
	IsAllow bool `json:"isAllow" description:"是否允许  true允许 false不允许"`
}
