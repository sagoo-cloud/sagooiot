package notice

import (
	"sagooiot/api/v1/common"

	"github.com/gogf/gf/v2/frame/g"
)

// GetNoticeTemplateListReq 获取数据列表
type GetNoticeTemplateListReq struct {
	g.Meta      `path:"/template/list" method:"get" summary:"获取通知模版列表" tags:"通知服务管理"`
	ConfigId    string `json:"configId"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Code        string `json:"code"          description:""`
	common.PaginationReq
}
type GetNoticeTemplateListRes struct {
	Data []GetNoticeTemplateByIdRes
	common.PaginationRes
}

// GetNoticeTemplateByIdReq 获取指定ID的数据
type GetNoticeTemplateByIdReq struct {
	g.Meta `path:"/template/get" method:"get" summary:"获取通知模版" tags:"通知服务管理"`
	Id     string `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetNoticeTemplateByIdRes struct {
	ConfigId    string `json:"configId"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Code        string `json:"code"          description:""`
	Title       string `json:"title"          description:""`
	Content     string `json:"content"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
	Id          string `json:"id"          description:""`
}

// GetNoticeTemplateByConfigIdReq 获取指定ConfigId的数据
type GetNoticeTemplateByConfigIdReq struct {
	g.Meta   `path:"/template/getbyconfig" method:"get" summary:"获取通知模版" tags:"通知服务管理"`
	ConfigId string `json:"configId"        description:"configId" v:"required#通知配置ID不能为空"`
}
type GetNoticeTemplateByConfigIdRes struct {
	ConfigId    string `json:"configId"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Code        string `json:"code"          description:""`
	Title       string `json:"title"          description:""`
	Content     string `json:"content"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
	Id          string `json:"id"          description:""`
}

// AddNoticeTemplateReq 添加数据
type AddNoticeTemplateReq struct {
	g.Meta      `path:"/template/add" method:"post" summary:"添加通知模版" tags:"通知服务管理"`
	Code        string `json:"code"          description:""`
	Title       string `json:"title"          description:""`
	Content     string `json:"content"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
	ConfigId    string `json:"configId"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
}
type AddNoticeTemplateRes struct{}

// SaveNoticeTemplateReq 添加数据
type SaveNoticeTemplateReq struct {
	g.Meta      `path:"/template/save" method:"post" summary:"直接更新通知模版数据" tags:"通知服务管理"`
	Id          string `json:"id"          description:""`
	Code        string `json:"code"          description:""`
	Title       string `json:"title"          description:""`
	Content     string `json:"content"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
	ConfigId    string `json:"configId"            description:"configId" v:"required#configId不能为空"`
	SendGateway string `json:"sendGateway"          description:""`
}
type SaveNoticeTemplateRes struct{}

// EditNoticeTemplateReq 编辑数据api
type EditNoticeTemplateReq struct {
	g.Meta      `path:"/template/edit" method:"put" summary:"编辑通知模版" tags:"通知服务管理"`
	Id          string `json:"id"          description:""`
	SendGateway string `json:"sendGateway"          description:""`
	Code        string `json:"code"          description:""`
	Title       string `json:"title"          description:""`
	Content     string `json:"content"          description:""`
	CreatedAt   string `json:"createdAt"          description:""`
	ConfigId    string `json:"configId"          description:""`
}
type EditNoticeTemplateRes struct{}

// DeleteNoticeTemplateReq 删除数据
type DeleteNoticeTemplateReq struct {
	g.Meta `path:"/template/delete" method:"delete" summary:"删除通知模版" tags:"通知服务管理"`
	Ids    []string `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteNoticeTemplateRes struct{}
