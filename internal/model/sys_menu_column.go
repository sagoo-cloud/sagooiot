package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type MenuColumnDoInput struct {
	MenuId   string `json:"menuId"`
	ParentId string `json:"parentId"`
	Status   int    `json:"status"`
	Name     string `json:"name"`
}

type UserMenuColumnRes struct {
	Id          int                  `json:"id"          description:""`
	ParentId    int                  `json:"parentId"    description:"父ID"`
	MenuId      int                  `json:"menuId"      description:"菜单ID"`
	Name        string               `json:"name"        description:"名称"`
	Code        string               `json:"code"        description:"代表列表"`
	Description string               `json:"description" description:"描述"`
	Status      int                  `json:"status"      description:"状态 0 停用 1启用"`
	CreatedBy   uint                 `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time          `json:"createdAt"   description:"创建时间"`
	Children    []*UserMenuColumnRes `json:"children" description:"子集"`
}

type UserMenuColumnOut struct {
	Id          int                  `json:"id"          description:""`
	ParentId    int                  `json:"parentId"    description:"父ID"`
	MenuId      int                  `json:"menuId"      description:"菜单ID"`
	Name        string               `json:"name"        description:"名称"`
	Title       string               `json:"title"        description:"标题"`
	Code        string               `json:"code"        description:"代表列表"`
	Description string               `json:"description" description:"描述"`
	Status      int                  `json:"status"      description:"状态 0 停用 1启用"`
	CreatedBy   uint                 `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time          `json:"createdAt"   description:"创建时间"`
	Children    []*UserMenuColumnOut `json:"children" description:"子集"`
}

type AddMenuColumnInput struct {
	ParentId    int    `json:"parentId"    description:"父ID"`
	MenuId      int    `json:"menuId"      description:"菜单ID"`
	Name        string `json:"name"        description:"名称"`
	Code        string `json:"code"        description:"代表列表"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}

type DetailMenuColumnRes struct {
	Id          int         `json:"id"          description:""`
	ParentId    int         `json:"parentId"    description:"父ID"`
	MenuId      int         `json:"menuId"      description:"菜单ID"`
	Name        string      `json:"name"        description:"名称"`
	Code        string      `json:"code"        description:"代表列表"`
	Description string      `json:"description" description:"描述"`
	Status      int         `json:"status"      description:"状态 0 停用 1启用"`
	CreatedBy   uint        `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
}

type EditMenuColumnInput struct {
	Id          int    `json:"id"          description:"" `
	ParentId    int    `json:"parentId"    description:"父ID" `
	MenuId      int    `json:"menuId"      description:"菜单ID"`
	Name        string `json:"name"        description:"名称"`
	Code        string `json:"code"        description:"代表列表"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}
