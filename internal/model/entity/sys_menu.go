// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysMenu is the golang structure for table sys_menu.
type SysMenu struct {
	Id         uint        `json:"id"         description:""`
	ParentId   int         `json:"parentId"   description:"父ID"`
	Name       string      `json:"name"       description:"规则名称"`
	Title      string      `json:"title"      description:"菜单名称"`
	Icon       string      `json:"icon"       description:"图标"`
	Condition  string      `json:"condition"  description:"条件"`
	Remark     string      `json:"remark"     description:"备注"`
	MenuType   uint        `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int         `json:"weigh"      description:"权重"`
	IsHide     uint        `json:"isHide"     description:"显示状态"`
	Path       string      `json:"path"       description:"路由地址"`
	Component  string      `json:"component"  description:"组件路径"`
	IsLink     uint        `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string      `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    uint        `json:"modelId"    description:"模型ID"`
	IsIframe   uint        `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   uint        `json:"isCached"   description:"是否缓存"`
	Redirect   string      `json:"redirect"   description:"路由重定向地址"`
	IsAffix    uint        `json:"isAffix"    description:"是否固定"`
	LinkUrl    string      `json:"linkUrl"    description:"链接地址"`
	Status     int         `json:"status"     description:"状态 0 停用 1启用"`
	IsDeleted  int         `json:"isDeleted"  description:"是否删除 0未删除 1已删除"`
	CreatedBy  uint        `json:"createdBy"  description:"创建人"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedBy  int         `json:"updatedBy"  description:"修改人"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"更新时间"`
	DeletedBy  int         `json:"deletedBy"  description:"删除人"`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:"删除时间"`
}
