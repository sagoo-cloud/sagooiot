package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type GetApiAllReq struct {
	g.Meta `path:"/api/GetAll" method:"get" summary:"获取所有接口" tags:"接口API管理"`
	Method string `json:"method"    description:"请求方式(数据字典维护)"`
}
type GetApiAllRes struct {
	Data []*model.SysApiAllRes
}

type GetApiTreeReq struct {
	g.Meta  `path:"/api/tree" method:"get" summary:"获取接口列表" tags:"接口API管理"`
	Name    string `json:"name"        description:"名称"`
	Address string `json:"address"        description:"接口地址"`
	Status  int    `json:"status"      description:"状态 0 停用 1启用"`
	Types   int    `json:"types"      description:"类型 1 分类 2接口"`
}
type GetApiTreeRes struct {
	Info []*model.SysApiTreeRes
}

type AddApiReq struct {
	g.Meta   `path:"/api/add" method:"post" summary:"添加Api" tags:"接口API管理"`
	ParentId int    `json:"parentId"  description:""`
	Name     string `json:"name"      description:"名称"     v:"required#请输入名称"`
	Types    int    `json:"types"     description:"1 分类 2接口" v:"required#请选择类型"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护" v:"required#请选择接口类型"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址"`
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
	MenuIds  []int  `json:"menuIds"        description:"菜单Id数组"`
}

type AddApiRes struct {
}

type EditApiReq struct {
	g.Meta   `path:"/api/edit" method:"put" summary:"编辑Api" tags:"接口API管理"`
	Id       int    `json:"id"        description:""       v:"required#ApiId不能为空"`
	ParentId int    `json:"parentId"  description:""`
	Name     string `json:"name"      description:"名称"     v:"required#请输入名称"`
	Types    int    `json:"types"     description:"1 分类 2接口" v:"required#请选择类型"`
	ApiTypes string `json:"apiTypes"  description:"数据字典维护" v:"required#请选择接口类型"`
	Method   string `json:"method"    description:"请求方式(数据字典维护)"`
	Address  string `json:"address"   description:"接口地址" `
	Remark   string `json:"remark"    description:"备注"`
	Status   int    `json:"status"    description:"状态 0 停用 1启用"`
	Sort     int    `json:"sort"      description:"排序"`
	MenuIds  []int  `json:"menuIds"        description:"菜单Id数组"`
}
type EditApiRes struct {
}

type DetailApiReq struct {
	g.Meta `path:"/api/detail" method:"get" summary:"根据ID获取Api详情" tags:"接口API管理"`
	Id     int `json:"id" description:"ApiID"  v:"required#ID不能为空"`
}
type DetailApiRes struct {
	Data *model.SysApiRes
}

type DelApiReq struct {
	g.Meta `path:"/api/del" method:"delete" summary:"根据ID删除Api" tags:"接口API管理"`
	Id     int `json:"id" description:"ApiID"  v:"required#ID不能为空"`
}
type DelApiRes struct {
}

type EditApiStatusReq struct {
	g.Meta `path:"/api/editStatus" method:"put" summary:"编辑Api状态" tags:"接口API管理"`
	Id     int `json:"id" description:"接口ID"  v:"required#ID不能为空"`
	Status int `json:"status"    description:"状态 0 停用 1启用"`
}
type EditApiStatusRes struct {
}

type ImportApiFileReq struct {
	g.Meta `path:"/api/import" method:"post" summary:"导入Api文件" tags:"接口API管理"`
}
type ImportApiFileRes struct {
}

type BindApiMenusReq struct {
	g.Meta    `path:"/api/bindMenus" method:"post" summary:"批量绑定菜单" tags:"接口API管理"`
	BindMenus []*model.BindMenusReq `json:"bindMenus" description:"接口ID"`
}
type BindApiMenusRes struct {
}
