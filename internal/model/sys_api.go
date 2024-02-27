package model

type SysApiTreeRes struct {
	Id       uint             `json:"id"        description:""`
	ParentId int              `json:"parentId"  description:""`
	Name     string           `json:"name"      description:"名称"`
	Types    int              `json:"types"     description:"1 分类 2接口"`
	ApiTypes string           `json:"apiTypes"  description:"数据字典维护"`
	Method   string           `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string           `json:"address"   description:"接口地址"`
	Remark   string           `json:"remark"    description:"备注"`
	Status   int              `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int              `json:"sort"      description:"排序"`
	MenuIds  []int            `json:"menuIds"        description:"菜单Id数组"`
	Children []*SysApiTreeRes `json:"children" description:"子集"`
}

type SysApiTreeOut struct {
	Id       uint             `json:"id"        description:""`
	ParentId int              `json:"parentId"  description:""`
	Name     string           `json:"name"      description:"名称"`
	Types    int              `json:"types"     description:"1 分类 2接口"`
	ApiTypes string           `json:"apiTypes"  description:"数据字典维护"`
	Method   string           `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string           `json:"address"   description:"接口地址"`
	Remark   string           `json:"remark"    description:"备注"`
	Status   int              `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int              `json:"sort"      description:"排序"`
	MenuIds  []int            `json:"menuIds"        description:"菜单Id数组"`
	Children []*SysApiTreeOut `json:"children" description:"子集"`
}

type SysApiAllRes struct {
	Id       uint   `json:"id"        description:""`
	ParentId int    `json:"parentId"  description:""`
	Name     string `json:"name"      description:"名称"`
	Types    int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址"`
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
}

type SysApiAllOut struct {
	Id       uint   `json:"id"        description:""`
	ParentId int    `json:"parentId"  description:""`
	Name     string `json:"name"      description:"名称"`
	Types    int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址"`
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
}

type SysApiRes struct {
	Id       uint   `json:"id"        description:""`
	ParentId int    `json:"parentId"  description:""`
	Name     string `json:"name"      description:"名称"`
	Types    int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址"`
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
	MenuIds  []int  `json:"menuIds"        description:"菜单Id数组" v:"required#菜单ID不能为空"`
}

type SysApiOut struct {
	Id       int    `json:"id"        description:""`
	ParentId int    `json:"parentId"  description:""`
	Name     string `json:"name"      description:"名称"`
	Types    int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址"`
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
	MenuIds  []int  `json:"menuIds"        description:"菜单Id数组"`
}

type AuthorizeQueryApiRes struct {
	Id       int    `json:"id"        description:"此ID为菜单与API的关联ID"`
	ApiId    int    `json:"apiId"        description:"接口ID"`
	ParentId int    `json:"parentId"  description:""`
	Title    string `json:"title"      description:"标题"`
	Name     string `json:"name"      description:"名称"`
	Types    int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址"`
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
}

type AuthorizeQueryApiOut struct {
	Id       int    `json:"id"        description:"此ID为菜单与API的关联ID"`
	ApiId    int    `json:"apiId"        description:"接口ID"`
	ParentId int    `json:"parentId"  description:""`
	Title    string `json:"title"      description:"标题"`
	Name     string `json:"name"      description:"名称"`
	Types    int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址"`
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
}

type UserApiRes struct {
	Id        int    `json:"id"        description:""`
	MenuApiId int    `json:"menuApiId"        description:""`
	ParentId  int    `json:"parentId"  description:""`
	Name      string `json:"name"      description:"名称"`
	Types     int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes  string `json:"apiTypes"  description:"数据字典维护"`
	Method    string `json:"method"    description:"请求方式(数据字典维护)"`
	Address   string `json:"address"   description:"接口地址"`
	Remark    string `json:"remark"    description:"备注"`
	Status    int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort      int    `json:"sort"      description:"排序"`
}

type UserApiOut struct {
	Id        int    `json:"id"        description:""`
	MenuApiId int    `json:"menuApiId"        description:""`
	ParentId  int    `json:"parentId"  description:""`
	Name      string `json:"name"      description:"名称"`
	Types     int    `json:"types"     description:"1 分类 2接口"`
	ApiTypes  string `json:"apiTypes"  description:"数据字典维护"`
	Method    string `json:"method"    description:"请求方式(数据字典维护)"`
	Address   string `json:"address"   description:"接口地址"`
	Remark    string `json:"remark"    description:"备注"`
	Status    int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort      int    `json:"sort"      description:"排序"`
}

type AddApiInput struct {
	ParentId int    `json:"parentId"`
	Name     string `json:"name"`
	Types    int    `json:"types"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"`
	Remark   string `json:"remark"`
	Status   int    `json:"status"`
	Sort     int    `json:"sort"`
	MenuIds  []int  `json:"menuIds"`
}

type EditApiInput struct {
	Id       int    `json:"id"`
	ParentId int    `json:"parentId"`
	Name     string `json:"name"`
	Types    int    `json:"types"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"`
	Remark   string `json:"remark"`
	Status   int    `json:"status"`
	Sort     int    `json:"sort"`
	MenuIds  []int  `json:"menuIds"`
}

type BindMenusReq struct {
	Id      int   `json:"id"`
	MenuIds []int `json:"menuIds"`
}

type BindMenusInput struct {
	Id      int   `json:"id" v:"required#API ID不能为空"`
	MenuIds []int `json:"menuIds" v:"required#菜单ID不能为空"`
}
