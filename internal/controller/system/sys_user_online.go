package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysUserOnline = cSysUserOnline{}

type cSysUserOnline struct{}

// UserOnlineListReq 在线用户列表
func (c *cSysUserOnline) UserOnlineListReq(ctx context.Context, req *system.UserOnlineListReq) (res *system.UserOnlineListRes, err error) {
	var input *model.UserOnlineDoListInput
	if err = gconv.Scan(req, &input); err != nil {
		return
	}

	total, out, err := service.SysUserOnline().UserOnlineList(ctx, input)
	if err != nil {
		return
	}
	res = new(system.UserOnlineListRes)
	res.Total = total
	res.CurrentPage = req.PageNum
	if out != nil {
		if err = gconv.Scan(out, &res.Data); err != nil {
			return
		}
	}
	return
}

// UserOnlineStrongBack 强退
func (c *cSysUserOnline) UserOnlineStrongBack(ctx context.Context, req *system.UserOnlineStrongBackReq) (res *system.UserOnlineStrongBackRes, err error) {
	err = service.SysUserOnline().UserOnlineStrongBack(ctx, req.Id)
	if err != nil {
		return
	}
	return
}
