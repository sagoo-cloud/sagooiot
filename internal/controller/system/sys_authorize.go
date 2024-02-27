package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

var SysAuthorize = cSysAuthorize{}

type cSysAuthorize struct{}

// AuthorizeQuery 权限查询
func (c *cSysAuthorize) AuthorizeQuery(ctx context.Context, req *system.AuthorizeQueryReq) (res *system.AuthorizeQueryRes, err error) {
	out, err := service.SysAuthorize().AuthorizeQuery(ctx, req.ItemsType, req.MenuIds)
	if err != nil {
		return
	}
	var authorizeQueryTree []*model.AuthorizeQueryTreeRes
	if out != nil {
		if err = gconv.Scan(out, &authorizeQueryTree); err != nil {
			return
		}
	}
	res = &system.AuthorizeQueryRes{
		Data: authorizeQueryTree,
	}
	return
}

// AddAuthorize 授权
func (c *cSysAuthorize) AddAuthorize(ctx context.Context, req *system.AddAuthorizeReq) (res *system.AddAuthorizeRes, err error) {
	err = service.SysAuthorize().AddAuthorize(ctx, req.RoleId, req.MenuIds, req.ButtonIds, req.ColumnIds, req.ApiIds)
	return
}

// IsAllowAuthorize 判断是否允许授权
func (c *cSysAuthorize) IsAllowAuthorize(ctx context.Context, req *system.IsAllowAuthorizeReq) (res *system.IsAllowAuthorizeRes, err error) {
	isAllow, err := service.SysAuthorize().IsAllowAuthorize(ctx, req.RoleId)
	if err != nil {
		return
	}
	res = &system.IsAllowAuthorizeRes{
		IsAllow: isAllow,
	}
	return
}
