package common

import (
	"github.com/sagoo-cloud/sagooiot/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type BaseDbLinkDoReq struct {
	g.Meta   `path:"/base/db/list"    tags:"数据源管理" method:"get" summary:"数据源列表"`
	Name     string `p:"name"             description:"数据源名称"`
	Types    string `p:"types"            description:"驱动类型 mysql或oracle"`
	Host     string `p:"host"             description:"主机地址"`
	Port     string `p:"port"             description:"端口"`
	UserName string `p:"user_name"        description:"用户名称"`
	Status   int    `p:"status"           description:"状态:-1为全部,0为正常,1为停用"`
	*PaginationReq
}
type BaseDbLinkDoRes struct {
	Data []*model.BaseDbLinkRes
	PaginationRes
}

type AddBaseDbLinkReq struct {
	g.Meta      `path:"/base/db/add"     tags:"数据源管理" method:"post" summary:"添加数据源"`
	Name        string `json:"name"        description:"名称" v:"required#请输入数据源名称"`
	Types       string `json:"types"       description:"驱动类型 mysql或oracle" v:"required#请输入数据源驱动类型"`
	Host        string `json:"host"        description:"主机地址" v:"required#请输入数据源主机地址"`
	Port        int    `json:"port"        description:"端口号" v:"required#请输入数据源端口号"`
	UserName    string `json:"userName"    description:"用户名称" v:"required#请输入数据源用户名称"`
	Password    string `json:"password"    description:"密码" v:"required#请输入数据源密码"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}
type AddBaseDbLinkRes struct {
}

type DetailBaseDbLinkReq struct {
	g.Meta `path:"/base/db/detail" tags:"数据源管理"   method:"get" summary:"根据ID获取数据源详情"`
	Id     int `p:"id"               description:"数据源ID"  v:"required#ID不能为空"`
}
type DetailBaseDbLinkRes struct {
	Data *model.DetailBaseDbLinkRes
}

type EditBaseDbLinkReq struct {
	g.Meta      `path:"/base/db/edit"   method:"put"      summary:"编辑数据源" tags:"数据源管理"`
	Id          int    `json:"id"          description:"" v:"required#请输入数据源ID"`
	Name        string `json:"name"        description:"名称" v:"required#请输入数据源名称"`
	Types       string `json:"types"       description:"驱动类型 mysql或oracle" v:"required#请输入数据源驱动类型"`
	Host        string `json:"host"        description:"主机地址" v:"required#请输入数据源主机地址"`
	Port        int    `json:"port"        description:"端口号" v:"required#请输入数据源端口号"`
	UserName    string `json:"userName"    description:"用户名称" v:"required#请输入数据源用户名称"`
	Password    string `json:"password"    description:"密码" v:"required#请输入数据源密码"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}
type EditBaseDbLinkRes struct {
}

type DelBaseDbLinkReq struct {
	g.Meta `path:"/base/db/del"   method:"delete"     summary:"根据ID删除数据源" tags:"数据源管理"`
	Id     int `p:"id"              description:"数据源ID"  v:"required#ID不能为空"`
}
type DelBaseDbLinkRes struct {
}
