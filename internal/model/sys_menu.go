package model

import "github.com/gogf/gf/v2/frame/g"

// SysMenuRes 菜单列表返回字段
type SysMenuRes struct {
	Id         int64         `json:"id"         description:""`
	ParentId   int64         `json:"parentId"   description:"父ID"`
	Name       string        `json:"name"       description:"规则名称"`
	Title      string        `json:"title"      description:"规则名称"`
	Icon       string        `json:"icon"       description:"图标"`
	Condition  string        `json:"condition"  description:"条件"`
	Remark     string        `json:"remark"     description:"备注"`
	MenuType   int64         `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int           `json:"weigh"      description:"权重"`
	IsHide     int64         `json:"isHide"     description:"显示状态"`
	Path       string        `json:"path"       description:"路由地址"`
	Component  string        `json:"component"  description:"组件路径"`
	IsLink     int64         `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string        `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    int64         `json:"modelId"    description:"模型ID"`
	IsIframe   int64         `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   int64         `json:"isCached"   description:"是否缓存"`
	Redirect   string        `json:"redirect"   description:"路由重定向地址"`
	IsAffix    int64         `json:"isAffix"    description:"是否固定"`
	LinkUrl    string        `json:"linkUrl"    description:"链接地址"`
	Status     int           `json:"status"     description:"状态 0 停用 1启用"`
	Children   []*SysMenuRes `json:"children" description:"子集"`
}

type SysMenuOut struct {
	Id         int64         `json:"id"         description:""`
	ParentId   int64         `json:"parentId"   description:"父ID"`
	Name       string        `json:"name"       description:"规则名称"`
	Title      string        `json:"title"      description:"规则名称"`
	Icon       string        `json:"icon"       description:"图标"`
	Condition  string        `json:"condition"  description:"条件"`
	Remark     string        `json:"remark"     description:"备注"`
	MenuType   int64         `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int           `json:"weigh"      description:"权重"`
	IsHide     int64         `json:"isHide"     description:"显示状态"`
	Path       string        `json:"path"       description:"路由地址"`
	Component  string        `json:"component"  description:"组件路径"`
	IsLink     int64         `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string        `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    int64         `json:"modelId"    description:"模型ID"`
	IsIframe   int64         `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   int64         `json:"isCached"   description:"是否缓存"`
	Redirect   string        `json:"redirect"   description:"路由重定向地址"`
	IsAffix    int64         `json:"isAffix"    description:"是否固定"`
	LinkUrl    string        `json:"linkUrl"    description:"链接地址"`
	Status     int           `json:"status"     description:"状态 0 停用 1启用"`
	Children   []*SysMenuOut `json:"children" description:"子集"`
}

// SysMenuTreeRes 菜单树形结构
type SysMenuTreeRes struct {
	*SysMenuRes
	Children []*SysMenuTreeRes `json:"children"`
}

type DetailMenuRes struct {
	Id         int64  `json:"id"         description:""`
	ParentId   int    `json:"parentId"   description:"父ID"`
	Name       string `json:"name"       description:"规则名称"`
	Title      string `json:"title"      description:"规则名称"`
	Icon       string `json:"icon"       description:"图标"`
	Condition  string `json:"condition"  description:"条件"`
	Remark     string `json:"remark"     description:"备注"`
	MenuType   int64  `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int    `json:"weigh"      description:"权重"`
	IsHide     int64  `json:"isHide"     description:"显示状态"`
	Path       string `json:"path"       description:"路由地址"`
	Component  string `json:"component"  description:"组件路径"`
	IsLink     int64  `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    int64  `json:"modelId"    description:"模型ID"`
	IsIframe   int64  `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   int64  `json:"isCached"   description:"是否缓存"`
	Redirect   string `json:"redirect"   description:"路由重定向地址"`
	IsAffix    int64  `json:"isAffix"    description:"是否固定"`
	LinkUrl    string `json:"linkUrl"    description:"链接地址"`
	Status     int    `json:"status"     description:"状态 0 停用 1启用"`
}

type AddMenuInput struct {
	MenuType   uint   `p:"menuType"`
	ParentId   int    `p:"parentId"`
	Name       string `p:"name"`
	Title      string `p:"title"`
	Icon       string `p:"icon"`
	Weigh      int    `p:"weigh"`
	Condition  string `p:"condition"`
	Remark     string `p:"remark"`
	IsHide     uint   `p:"isHide"`
	Path       string `p:"path"`
	Component  string `p:"component"`
	IsLink     uint   `p:"isLink"`
	IsIframe   uint   `p:"isIframe"`
	IsCached   uint   `p:"isKeepAlive"`
	IsAffix    uint   `p:"isAffix"`
	LinkUrl    string `p:"linkUrl"`
	Status     int    `p:"status"`
	ModuleType string `p:"moduleType"`
}

type EditMenuInput struct {
	Id         int64  `json:"id"`
	MenuType   uint   `p:"menuType"`
	ParentId   int    `p:"parentId"`
	Name       string `p:"name"`
	Title      string `p:"title"`
	Icon       string `p:"icon"`
	Weigh      int    `p:"weigh"`
	Condition  string `p:"condition"`
	Remark     string `p:"remark"`
	IsHide     uint   `p:"isHide"`
	Path       string `p:"path"`
	Component  string `p:"component"`
	IsLink     uint   `p:"isLink"`
	IsIframe   uint   `p:"isIframe"`
	IsCached   uint   `p:"isKeepAlive"`
	IsAffix    uint   `p:"isAffix"`
	LinkUrl    string `p:"linkUrl"`
	Status     int    `p:"status"`
	ModuleType string `p:"moduleType"`
}

type UserMenu struct {
	Id        uint   `json:"id"        description:""`
	Pid       uint   `json:"pid"       description:"父ID"`
	Name      string `json:"name"      description:"规则名称"`
	Component string `json:"component" description:"组件路径"`
	Path      string `json:"path"      description:"路由地址"`
	*MenuMeta `json:"meta"`
}

type UserMenus struct {
	*UserMenu `json:""`
	Children  []*UserMenus `json:"children" description:"子集"`
}

type MenuMeta struct {
	Icon     string `json:"icon"        description:"图标"`
	Title    string `json:"title"       description:"规则名称"`
	IsLink   string `json:"isLink"      description:"是否外链 1是 0否"`
	IsHide   bool   `json:"isHide"      description:"显示状态"`
	IsAffix  bool   `json:"isAffix"     description:"是否固定"`
	IsIframe bool   `json:"isIframe"    description:"是否内嵌iframe"`
}

type AuthorizeQueryTreeRes struct {
	Id         uint    `json:"id"         description:""`
	ParentId   int     `json:"parentId"   description:"父ID"`
	Name       string  `json:"name"       description:"规则名称"`
	Title      string  `json:"title"      description:"规则名称"`
	Icon       string  `json:"icon"       description:"图标"`
	Condition  string  `json:"condition"  description:"条件"`
	Remark     string  `json:"remark"     description:"备注"`
	MenuType   uint    `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int     `json:"weigh"      description:"权重"`
	IsHide     uint    `json:"isHide"     description:"显示状态"`
	Path       string  `json:"path"       description:"路由地址"`
	Component  string  `json:"component"  description:"组件路径"`
	IsLink     uint    `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string  `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    uint    `json:"modelId"    description:"模型ID"`
	IsIframe   uint    `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   uint    `json:"isCached"   description:"是否缓存"`
	Redirect   string  `json:"redirect"   description:"路由重定向地址"`
	IsAffix    uint    `json:"isAffix"    description:"是否固定"`
	LinkUrl    string  `json:"linkUrl"    description:"链接地址"`
	Status     int     `json:"status"     description:"状态 0 停用 1启用"`
	Children   []g.Map `json:"children" description:"子集 菜单，按钮，列表，接口API"`
}

type AuthorizeQueryTreeOut struct {
	Id         uint    `json:"id"         description:""`
	ParentId   int     `json:"parentId"   description:"父ID"`
	Name       string  `json:"name"       description:"规则名称"`
	Title      string  `json:"title"      description:"菜单名称"`
	Icon       string  `json:"icon"       description:"图标"`
	Condition  string  `json:"condition"  description:"条件"`
	Remark     string  `json:"remark"     description:"备注"`
	MenuType   uint    `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int     `json:"weigh"      description:"权重"`
	IsHide     uint    `json:"isHide"     description:"显示状态"`
	Path       string  `json:"path"       description:"路由地址"`
	Component  string  `json:"component"  description:"组件路径"`
	IsLink     uint    `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string  `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    uint    `json:"modelId"    description:"模型ID"`
	IsIframe   uint    `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   uint    `json:"isCached"   description:"是否缓存"`
	Redirect   string  `json:"redirect"   description:"路由重定向地址"`
	IsAffix    uint    `json:"isAffix"    description:"是否固定"`
	LinkUrl    string  `json:"linkUrl"    description:"链接地址"`
	Status     int     `json:"status"     description:"状态 0 停用 1启用"`
	Children   []g.Map `json:"children" description:"子集 菜单，按钮，列表，接口API"`
}

type UserMenuTreeRes struct {
	Id         uint                 `json:"id"         description:""`
	ParentId   int                  `json:"parentId"   description:"父ID"`
	Name       string               `json:"name"       description:"规则名称"`
	Title      string               `json:"title"      description:"规则名称"`
	Icon       string               `json:"icon"       description:"图标"`
	Condition  string               `json:"condition"  description:"条件"`
	Remark     string               `json:"remark"     description:"备注"`
	MenuType   uint                 `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int                  `json:"weigh"      description:"权重"`
	IsHide     uint                 `json:"isHide"     description:"显示状态"`
	Path       string               `json:"path"       description:"路由地址"`
	Component  string               `json:"component"  description:"组件路径"`
	IsLink     uint                 `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string               `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    uint                 `json:"modelId"    description:"模型ID"`
	IsIframe   uint                 `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   uint                 `json:"isCached"   description:"是否缓存"`
	Redirect   string               `json:"redirect"   description:"路由重定向地址"`
	IsAffix    uint                 `json:"isAffix"    description:"是否固定"`
	LinkUrl    string               `json:"linkUrl"    description:"链接地址"`
	Status     int                  `json:"status"     description:"状态 0 停用 1启用"`
	Button     []*UserMenuButtonRes `json:"button" description:"按钮集合"`
	Column     []*UserMenuColumnRes `json:"column" description:"列表集合"`
	Api        []*UserApiRes        `json:"api" description:"接口API集合"`
	Children   []*UserMenuTreeRes   `json:"children" description:"子集"`
}

type UserMenuTreeOut struct {
	Id         uint                 `json:"id"         description:""`
	ParentId   int                  `json:"parentId"   description:"父ID"`
	Name       string               `json:"name"       description:"规则名称"`
	Title      string               `json:"title"      description:"规则名称"`
	Icon       string               `json:"icon"       description:"图标"`
	Condition  string               `json:"condition"  description:"条件"`
	Remark     string               `json:"remark"     description:"备注"`
	MenuType   uint                 `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int                  `json:"weigh"      description:"权重"`
	IsHide     uint                 `json:"isHide"     description:"显示状态"`
	Path       string               `json:"path"       description:"路由地址"`
	Component  string               `json:"component"  description:"组件路径"`
	IsLink     uint                 `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string               `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    uint                 `json:"modelId"    description:"模型ID"`
	IsIframe   uint                 `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   uint                 `json:"isCached"   description:"是否缓存"`
	Redirect   string               `json:"redirect"   description:"路由重定向地址"`
	IsAffix    uint                 `json:"isAffix"    description:"是否固定"`
	LinkUrl    string               `json:"linkUrl"    description:"链接地址"`
	Status     int                  `json:"status"     description:"状态 0 停用 1启用"`
	Button     []*UserMenuButtonOut `json:"button" description:"按钮集合"`
	Column     []*UserMenuColumnOut `json:"column" description:"列表集合"`
	Api        []*UserApiOut        `json:"api" description:"接口API集合"`
	Children   []*UserMenuTreeOut   `json:"children" description:"子集"`
}

type MenuJoinRes struct {
	Id         int64                `json:"id"         description:""`
	ParentId   int64                `json:"parentId"   description:"父ID"`
	Name       string               `json:"name"       description:"规则名称"`
	Title      string               `json:"title"      description:"规则名称"`
	Icon       string               `json:"icon"       description:"图标"`
	Condition  string               `json:"condition"  description:"条件"`
	Remark     string               `json:"remark"     description:"备注"`
	MenuType   int64                `json:"menuType"   description:"类型 0目录 1菜单 2按钮"`
	Weigh      int                  `json:"weigh"      description:"权重"`
	IsHide     int64                `json:"isHide"     description:"显示状态"`
	Path       string               `json:"path"       description:"路由地址"`
	Component  string               `json:"component"  description:"组件路径"`
	IsLink     int64                `json:"isLink"     description:"是否外链 1是 0否"`
	ModuleType string               `json:"moduleType" description:"所属模块 system 运维 company企业"`
	ModelId    int64                `json:"modelId"    description:"模型ID"`
	IsIframe   int64                `json:"isIframe"   description:"是否内嵌iframe"`
	IsCached   int64                `json:"isCached"   description:"是否缓存"`
	Redirect   string               `json:"redirect"   description:"路由重定向地址"`
	IsAffix    int64                `json:"isAffix"    description:"是否固定"`
	LinkUrl    string               `json:"linkUrl"    description:"链接地址"`
	Status     int                  `json:"status"     description:"状态 0 停用 1启用"`
	Button     []*UserMenuButtonRes `json:"button" description:"按钮集合"`
	Column     []*UserMenuColumnRes `json:"column" description:"列表集合"`
	Api        []*SysMenuApiRes     `json:"api" description:"接口API集合"`
}
