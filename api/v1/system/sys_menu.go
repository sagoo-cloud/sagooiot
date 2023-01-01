package system

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

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

//**********  菜单按钮关联开始  **********

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

//**********  菜单按钮关联结束  **********

//**********  菜单列表关联开始  **********

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

//**********  菜单列表关联结束  **********
