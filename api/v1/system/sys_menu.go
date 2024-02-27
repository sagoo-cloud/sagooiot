package system

import (
	"sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type MenuDoReq struct {
	g.Meta `path:"/menu/tree"    tags:"菜单管理" method:"get" summary:"菜单列表"`
	Title  string `p:"title"     dc:"菜单标题"`
	Status int    `p:"status"    dc:"状态:0为停用,1为正常"`
}
type MenuDoRes struct {
	Data []*model.SysMenuRes
}

type AddMenuReq struct {
	g.Meta     `path:"/menu/add"     tags:"菜单管理" method:"post" summary:"添加菜单"`
	MenuType   uint   `p:"menuType"  v:"min:0|max:2#菜单类型最小值为:min|菜单类型最大值为:max" dc:"菜单类型 0目录 1菜单 2按钮"`
	ParentId   int    `p:"parentId"  v:"required#请选择上级" dc:"父ID"`
	Name       string `p:"name"      v:"required#请填写规则名称" dc:"规则名称:具有唯一性"`
	Title      string `p:"title"     v:"required|length:1,100#请填写标题|标题长度在:min到:max位" dc:"规则名称"`
	Icon       string `p:"icon"      dc:"图标"`
	Weigh      int    `p:"weigh"     dc:"权重"`
	Condition  string `p:"condition" dc:"条件"`
	Remark     string `p:"remark"    dc:"备注"`
	IsHide     uint   `p:"isHide"    dc:"显示状态"`
	Path       string `p:"path"      dc:"路由地址"`
	Component  string `p:"component" v:"required-if:menuType,1#组件路径不能为空" dc:"组件路径"`
	IsLink     uint   `p:"isLink"    dc:"是否外链 1是 0否"`
	IsIframe   uint   `p:"isIframe"  dc:"是否内嵌iframe"`
	IsCached   uint   `p:"isKeepAlive" dc:"是否缓存"`
	IsAffix    uint   `p:"isAffix"   dc:"是否固定"`
	LinkUrl    string `p:"linkUrl"   dc:"链接地址"`
	Status     int    `p:"status"   dc:"状态 0 停用 1启用"`
	ModuleType string `p:"moduleType"   dc:"所属模块 system 运维 company企业"`
}
type AddMenuRes struct {
}

type DetailMenuReq struct {
	g.Meta `path:"/menu/detail" tags:"菜单管理" method:"get" summary:"根据ID获取菜单详情"`
	Id     int64 `p:"id" description:"菜单ID"  v:"required#ID不能为空"`
}
type DetailMenuRes struct {
	Data *model.DetailMenuRes
}

type EditMenuReq struct {
	g.Meta     `path:"/menu/edit" method:"put" summary:"编辑菜单" tags:"菜单管理"`
	Id         int64  `json:"id"    description:"菜单ID" v:"required#菜单ID不能为空"`
	MenuType   uint   `p:"menuType"  v:"min:0|max:2#菜单类型最小值为:min|菜单类型最大值为:max" dc:"菜单类型 0目录 1菜单 2按钮"`
	ParentId   int    `p:"parentId"  v:"required#请选择上级" dc:"父ID"`
	Name       string `p:"name"      v:"required#请填写规则名称" dc:"规则名称:具有唯一性"`
	Title      string `p:"title"     v:"required|length:1,100#请填写标题|标题长度在:min到:max位" dc:"规则名称"`
	Icon       string `p:"icon"      dc:"图标"`
	Weigh      int    `p:"weigh"     dc:"权重"`
	Condition  string `p:"condition" dc:"条件"`
	Remark     string `p:"remark"    dc:"备注"`
	IsHide     uint   `p:"isHide"    dc:"显示状态"`
	Path       string `p:"path"      dc:"路由地址"`
	Component  string `p:"component" v:"required-if:menuType,1#组件路径不能为空" dc:"组件路径"`
	IsLink     uint   `p:"isLink"    dc:"是否外链 1是 0否"`
	IsIframe   uint   `p:"isIframe"  dc:"是否内嵌iframe"`
	IsCached   uint   `p:"isKeepAlive" dc:"是否缓存"`
	IsAffix    uint   `p:"isAffix"   dc:"是否固定"`
	LinkUrl    string `p:"linkUrl"   dc:"链接地址"`
	Status     int    `p:"status"   dc:"状态 0 停用 1启用"`
	ModuleType string `p:"moduleType"   dc:"所属模块 system 运维 company企业"`
}
type EditMenuRes struct {
}

type DelMenuReq struct {
	g.Meta `path:"/menu/del" method:"delete" summary:"根据ID删除菜单" tags:"菜单管理"`
	Id     int64 `p:"id" description:"菜单ID"  v:"required#ID不能为空"`
}
type DelMenuRes struct {
}
