package notice

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
)

// GetNoticeInfoListReq 获取数据列表
type GetNoticeInfoListReq struct {
	g.Meta   `path:"/info/list" method:"get" summary:"获取通知信息列表" tags:"通知服务管理"`
	ConfigId string `json:"configId"   description:""`
	ComeFrom string `json:"comeFrom"   description:""`
	Method   string `json:"method"     description:""`
	Status   int    `json:"status"     description:""`
	common.PaginationReq
}
type GetNoticeInfoListRes struct {
	Data []GetNoticeInfoByIdRes
	common.PaginationRes
}

// GetNoticeInfoByIdReq 获取指定ID的数据
type GetNoticeInfoByIdReq struct {
	g.Meta `path:"/info/get" method:"get" summary:"获取通知信息" tags:"通知服务管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetNoticeInfoByIdRes struct {
	OrgIds     string `json:"orgIds"          description:""`
	MethodNum  string `json:"methodNum"          description:""`
	CreatedAt  string `json:"createdAt"          description:""`
	MsgTitle   string `json:"msgTitle"          description:""`
	Totag      string `json:"totag"          description:""`
	Status     string `json:"status"          description:""`
	MethodCron string `json:"methodCron"          description:""`
	Id         string `json:"id"          description:""`
	ComeFrom   string `json:"comeFrom"          description:""`
	Method     string `json:"method"          description:""`
	MsgBody    string `json:"msgBody"          description:""`
	MsgUrl     string `json:"msgUrl"          description:""`
	ConfigId   string `json:"configId"          description:""`
	UserIds    string `json:"userIds"          description:""`
}

// AddNoticeInfoReq 添加数据
type AddNoticeInfoReq struct {
	g.Meta     `path:"/info/add" method:"post" summary:"添加通知信息" tags:"通知服务管理"`
	Totag      string `json:"totag"          description:""`
	Status     string `json:"status"          description:""`
	MethodCron string `json:"methodCron"          description:""`
	ComeFrom   string `json:"comeFrom"          description:""`
	Method     string `json:"method"          description:""`
	MsgBody    string `json:"msgBody"          description:""`
	MsgUrl     string `json:"msgUrl"          description:""`
	ConfigId   string `json:"configId"          description:""`
	UserIds    string `json:"userIds"          description:""`
	OrgIds     string `json:"orgIds"          description:""`
	MethodNum  string `json:"methodNum"          description:""`
	CreatedAt  string `json:"createdAt"          description:""`
	MsgTitle   string `json:"msgTitle"          description:""`
}
type AddNoticeInfoRes struct{}

// EditNoticeInfoReq 编辑数据api
type EditNoticeInfoReq struct {
	g.Meta     `path:"/info/edit" method:"put" summary:"编辑通知信息" tags:"通知服务管理"`
	UserIds    string `json:"userIds"          description:""`
	MsgTitle   string `json:"msgTitle"          description:""`
	OrgIds     string `json:"orgIds"          description:""`
	MethodNum  string `json:"methodNum"          description:""`
	CreatedAt  string `json:"createdAt"          description:""`
	Id         string `json:"id"          description:""`
	Totag      string `json:"totag"          description:""`
	Status     string `json:"status"          description:""`
	MethodCron string `json:"methodCron"          description:""`
	MsgUrl     string `json:"msgUrl"          description:""`
	ConfigId   string `json:"configId"          description:""`
	ComeFrom   string `json:"comeFrom"          description:""`
	Method     string `json:"method"          description:""`
	MsgBody    string `json:"msgBody"          description:""`
}
type EditNoticeInfoRes struct{}

// DeleteNoticeInfoReq 删除数据
type DeleteNoticeInfoReq struct {
	g.Meta `path:"/info/delete" method:"delete" summary:"删除通知信息" tags:"通知服务管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteNoticeInfoRes struct{}
