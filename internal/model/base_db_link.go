package model

import "github.com/gogf/gf/v2/os/gtime"

type BaseDbLinkDoInput struct {
	Name     string `p:"name"             description:"数据源名称"`
	Types    string `p:"types"            description:"驱动类型 mysql或oracle"`
	Host     string `p:"host"             description:"主机地址"`
	Port     string `p:"port"             description:"端口"`
	UserName string `p:"user_name"        description:"用户名称"`
	Status   int    `p:"status"           description:"状态:-1为全部,0为正常,1为停用"`
	*PaginationInput
}

type BaseDbLinkOut struct {
	Id          int         `json:"id"          description:""`
	Name        string      `json:"name"        description:"名称"`
	Types       string      `json:"types"       description:"驱动类型 mysql或oracle"`
	Host        string      `json:"host"        description:"主机地址"`
	Port        int         `json:"port"        description:"端口号"`
	UserName    string      `json:"userName"    description:"用户名称"`
	Password    string      `json:"password"    description:"密码"`
	Description string      `json:"description" description:"描述"`
	Status      int         `json:"status"      description:"状态 0 停用 1启用"`
	IsDeleted   int         `json:"isDeleted"   description:"是否删除 0未删除 1已删除"`
	CreatedBy   uint        `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
}

// BaseDbLinkRes 数据源列表返回字段
type BaseDbLinkRes struct {
	Id          int         `json:"id"          description:""`
	Name        string      `json:"name"        description:"名称"`
	Types       string      `json:"types"       description:"驱动类型 mysql或oracle"`
	Host        string      `json:"host"        description:"主机地址"`
	Port        int         `json:"port"        description:"端口号"`
	UserName    string      `json:"userName"    description:"用户名称"`
	Password    string      `json:"password"    description:"密码"`
	Description string      `json:"description" description:"描述"`
	Status      int         `json:"status"      description:"状态 0 停用 1启用"`
	IsDeleted   int         `json:"isDeleted"   description:"是否删除 0未删除 1已删除"`
	CreatedBy   uint        `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
}

type DetailBaseDbLinkRes struct {
	Id          int         `json:"id"          description:""`
	Name        string      `json:"name"        description:"名称"`
	Types       string      `json:"types"       description:"驱动类型 mysql或oracle"`
	Host        string      `json:"host"        description:"主机地址"`
	Port        int         `json:"port"        description:"端口号"`
	UserName    string      `json:"userName"    description:"用户名称"`
	Password    string      `json:"password"    description:"密码"`
	Description string      `json:"description" description:"描述"`
	Status      int         `json:"status"      description:"状态 0 停用 1启用"`
	IsDeleted   int         `json:"isDeleted"   description:"是否删除 0未删除 1已删除"`
	CreatedBy   uint        `json:"createdBy"   description:"创建人"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
}

type AddBaseDbLinkInput struct {
	Name        string `json:"name"        description:"名称" v:"required#请输入数据源名称"`
	Types       string `json:"types"       description:"驱动类型 mysql或oracle" v:"required#请输入数据源驱动类型"`
	Host        string `json:"host"        description:"主机地址" v:"required#请输入数据源主机地址"`
	Port        int    `json:"port"        description:"端口号" v:"required#请输入数据源端口号"`
	UserName    string `json:"userName"    description:"用户名称" v:"required#请输入数据源用户名称"`
	Password    string `json:"password"    description:"密码" v:"required#请输入数据源密码"`
	Description string `json:"description" description:"描述"`
	Status      int    `json:"status"      description:"状态 0 停用 1启用"`
}

type EditBaseDbLinkInput struct {
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
