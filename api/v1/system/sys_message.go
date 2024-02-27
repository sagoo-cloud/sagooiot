package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

type GetMessageListReq struct {
	g.Meta `path:"/message/list" method:"get" summary:"获取消息列表" tags:"消息管理"`
	Types  string `json:"types"        description:"消息类型"`
	Title  string `json:"title"     description:"标题"`
	*common.PaginationReq
}
type GetMessageListRes struct {
	Info []*model.MessageListRes
	common.PaginationRes
}

type GetUnReadMessageAllReq struct {
	g.Meta `path:"/message/allUnRead" method:"get" summary:"获取所有未读消息" tags:"消息管理"`
	*common.PaginationReq
}
type GetUnReadMessageAllRes struct {
	Info []*model.MessageListRes
	common.PaginationRes
}

type GetUnReadMessageCountReq struct {
	g.Meta `path:"/message/unReadCount" method:"get" summary:"获取所有未读消息数量" tags:"消息管理"`
}
type GetUnReadMessageCountRes struct {
	Count int
}

type DelMessageReq struct {
	g.Meta `path:"/message/del" method:"delete" summary:"批量删除消息" tags:"消息管理"`
	Ids    []int `json:"ids"        description:"ID" v:"required#ID不能为空"`
}
type DelMessageRes struct {
}

type ClearMessageReq struct {
	g.Meta `path:"/message/clear" method:"delete" summary:"一键清空消息" tags:"消息管理"`
}
type ClearMessageRes struct {
}

type ReadMessageReq struct {
	g.Meta `path:"/message/read" method:"put" summary:"阅读消息" tags:"消息管理"`
	Id     int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type ReadMessageRes struct {
}

type ReadMessageAllReq struct {
	g.Meta `path:"/message/readAll" method:"put" summary:"全部已读" tags:"消息管理"`
}
type ReadMessageAllRes struct {
}
