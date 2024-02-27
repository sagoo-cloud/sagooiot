package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

// 获取列表api
type GetNotificationsListReq struct {
	g.Meta `path:"/notifications/list" method:"get" summary:"获取消息列表" tags:"通知中心管理"`
	common.PaginationReq
}
type GetNotificationsListRes struct {
	Data []*model.NotificationsRes
	common.PaginationRes
}

// 获取指定ID的数据api
type GetNotificationsByIdReq struct {
	g.Meta `path:"/notifications/get" method:"get" summary:"获取消息列表" tags:"通知中心管理"`
	Id     int `json:"id"        description:"id" v:"required#id不能为空"`
}
type GetNotificationsByIdRes struct {
	Data *model.NotificationsRes
}

// 添加数据api
type AddNotificationsReq struct {
	g.Meta    `path:"/notifications/add" method:"post" summary:"添加消息" tags:"通知中心管理"`
	Types     string `json:"types"          description:"类型"`
	CreatedAt string `json:"createdAt"          description:"发送时间"`
	Status    string `json:"status"          description:"0，未读，1，已读"`
	Title     string `json:"title"          description:"标题"`
	Doc       string `json:"doc"          description:"描述"`
	Source    string `json:"source"          description:"消息来源"`
}
type AddNotificationsRes struct{}

// 编辑数据api
type EditNotificationsReq struct {
	g.Meta    `path:"/notifications/edit" method:"put" summary:"编辑消息" tags:"通知中心管理"`
	Id        string `json:"id"          description:""`
	Title     string `json:"title"          description:"标题"`
	Doc       string `json:"doc"          description:"描述"`
	Source    string `json:"source"          description:"消息来源"`
	Types     string `json:"types"          description:"类型"`
	CreatedAt string `json:"createdAt"          description:"发送时间"`
	Status    string `json:"status"          description:"0，未读，1，已读"`
}
type EditNotificationsRes struct{}

// 删除数据api
type DeleteNotificationsReq struct {
	g.Meta `path:"/notifications/delete" method:"delete" summary:"删除消息" tags:"通知中心管理"`
	Ids    []int `json:"ids"        description:"ids" v:"required#ids不能为空"`
}
type DeleteNotificationsRes struct{}
