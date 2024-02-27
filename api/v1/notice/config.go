package notice

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
)

// GetNoticeConfigListReq 获取数据列表
type GetNoticeConfigListReq struct {
	g.Meta      `path:"/config/list" method:"get" summary:"获取通知配置列表" tags:"通知服务管理"`
	SendGateway string `json:"sendGateway"          description:""`
	Types       string `json:"types"          description:""`
	common.PaginationReq
}
type GetNoticeConfigListRes struct {
	Data []GetNoticeConfigByIdRes
	common.PaginationRes
}

// GetNoticeConfigByIdReq 获取指定ID的数据
type GetNoticeConfigByIdReq struct {
	g.Meta `path:"/config/get" method:"get" summary:"获取通知配置" tags:"通知服务管理"`
	Id     string `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetNoticeConfigByIdRes struct {
	SendGateway string `json:"sendGateway"          description:""`
	Types       string `json:"types"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
	Id          string `json:"id"          description:""`
	Title       string `json:"title"          description:""`
}

// AddNoticeConfigReq 添加数据
type AddNoticeConfigReq struct {
	g.Meta      `path:"/config/add" method:"post" summary:"添加通知配置" tags:"通知服务管理"`
	Title       string `json:"title"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Types       string `json:"types"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
}
type AddNoticeConfigRes struct{}

// EditNoticeConfigReq 编辑数据api
type EditNoticeConfigReq struct {
	g.Meta      `path:"/config/edit" method:"put" summary:"编辑通知配置" tags:"通知服务管理"`
	Id          string `json:"id"          description:""`
	Title       string `json:"title"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Types       string `json:"types"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
}
type EditNoticeConfigRes struct{}

// DeleteNoticeConfigReq 删除数据
type DeleteNoticeConfigReq struct {
	g.Meta `path:"/config/delete" method:"delete" summary:"删除通知配置" tags:"通知服务管理"`
	Ids    []string `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteNoticeConfigRes struct{}
