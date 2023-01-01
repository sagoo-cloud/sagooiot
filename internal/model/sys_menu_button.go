package model

import "github.com/gogf/gf/v2/os/gtime"

type UserMenuButtonRes struct {
	Id          int                  `json:"id"          description:""`
	ParentId    int                  `json:"parentId"    description:"父ID"`
	MenuId      int                  `json:"menuId"      description:"菜单ID"`
	Name        string               `json:"name"        description:"名称"`
	Types       string               `json:"types"       description:"类型 自定义 add添加 edit编辑 del 删除"`
	Description string               `json:"description" description:"描述"`
	Status      int                  `json:"status"      description:"状态 0 停用 1启用"`
	Children    []*UserMenuButtonRes `json:"children" description:"子集"`
}

type UserMenuButtonOut struct {
	Id          int                  `json:"id"          description:""`
	ParentId    int                  `json:"parentId"    description:"父ID"`
	MenuId      int                  `json:"menuId"      description:"菜单ID"`
	Name        string               `json:"name"        description:"名称"`
	Title       string               `json:"title"        description:"标题"`
	Types       string               `json:"types"       description:"类型 自定义 add添加 edit编辑 del 删除"`
	Description string               `json:"description" description:"描述"`
	Status      int                  `json:"status"      description:"状态 0 停用 1启用"`
	Children    []*UserMenuButtonOut `json:"children" description:"子集"`
}

type AddMenuButtonInput struct {
	ParentId    int    `json:"parentId"`
	MenuId      int    `json:"menuId"`
	Name        string `json:"name"`
	Types       string `json:"types"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type DetailMenuButtonRes struct {
	Id          int         `json:"id"          description:""`
	ParentId    int         `json:"parentId"    description:"父ID"`
	MenuId      int         `json:"menuId"      description:"菜单ID"`
	Name        string      `json:"name"        description:"名称"`
	Types       string      `json:"types"       description:"类型 自定义 add添加 edit编辑 del 删除"`
	Description string      `json:"description" description:"描述"`
	Status      int         `json:"status"      description:"状态 0 停用 1启用"`
	IsDeleted   int         `json:"isDeleted"   description:"是否删除 0未删除 1已删除"`
	CreatedBy   uint        `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
}

type EditMenuButtonInput struct {
	Id          int    `json:"id"`
	ParentId    int    `json:"parentId"`
	MenuId      int    `json:"menuId"`
	Name        string `json:"name"`
	Types       string `json:"types"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}
