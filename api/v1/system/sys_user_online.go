package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/api/v1/common"
	"sagooiot/internal/model"
)

type UserOnlineListReq struct {
	g.Meta `path:"/userOnline/list" method:"get" summary:"列表" tags:"在线用户管理"`
	*common.PaginationReq
}
type UserOnlineListRes struct {
	Data []*model.UserOnlineListRes
	common.PaginationRes
}

type UserOnlineStrongBackReq struct {
	g.Meta `path:"/userOnline/strongBack" method:"delete" summary:"强退" tags:"在线用户管理"`
	Id     int `json:"id"        description:"ID" v:"required#ID不能为空"`
}
type UserOnlineStrongBackRes struct {
}
