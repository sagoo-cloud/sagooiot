package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
)

// GetBlacklistListReq 获取数据列表
type GetBlacklistListReq struct {
	g.Meta `path:"/blacklist/list" method:"get" summary:"获取黑名单列表" tags:"黑名单管理"`
	common.PaginationReq
}
type GetBlacklistListRes struct {
	Data []GetBlacklistByIdRes
	common.PaginationRes
}

// GetBlacklistByIdReq 获取指定ID的数据
type GetBlacklistByIdReq struct {
	g.Meta `path:"/blacklist/get" method:"get" summary:"获取黑名单" tags:"黑名单管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetBlacklistByIdRes struct {
	UpdatedAt string `json:"updatedAt"          description:"更新时间"`
	Id        string `json:"id"          description:"黑名单ID"`
	Ip        string `json:"ip"          description:"IP地址"`
	Remark    string `json:"remark"          description:"备注"`
	Status    string `json:"status"          description:"状态"`
	CreatedAt string `json:"createdAt"          description:"创建时间"`
}

// AddBlacklistReq 添加数据
type AddBlacklistReq struct {
	g.Meta    `path:"/blacklist/add" method:"post" summary:"添加黑名单" tags:"黑名单管理"`
	Ip        string `json:"ip"          description:"IP地址"`
	Status    string `json:"status"          description:"状态"`
	Remark    string `json:"remark"          description:"备注"`
	CreatedAt string `json:"createdAt"          description:"创建时间"`
	UpdatedAt string `json:"updatedAt"          description:"更新时间"`
}
type AddBlacklistRes struct{}

// EditBlacklistReq 编辑数据api
type EditBlacklistReq struct {
	g.Meta    `path:"/blacklist/edit" method:"put" summary:"编辑黑名单" tags:"黑名单管理"`
	Id        string `json:"id"          description:"黑名单ID"`
	Ip        string `json:"ip"          description:"IP地址"`
	Remark    string `json:"remark"          description:"备注"`
	Status    string `json:"status"          description:"状态"`
	CreatedAt string `json:"createdAt"          description:"创建时间"`
	UpdatedAt string `json:"updatedAt"          description:"更新时间"`
}
type EditBlacklistRes struct{}

// DeleteBlacklistReq 删除数据
type DeleteBlacklistReq struct {
	g.Meta `path:"/blacklist/delete" method:"delete" summary:"删除黑名单" tags:"黑名单管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteBlacklistRes struct{}

// StatusReq 更新状态
type StatusReq struct {
	g.Meta `path:"/blacklist/status" method:"post" summary:"更新黑名单状态" tags:"黑名单管理"`
	Id     int `json:"id"          description:"黑名单ID"`
	Status int `json:"status"          description:"状态"`
}
type StatusRes struct{}
