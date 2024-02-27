package common

import (
	"context"
	"sagooiot/api/v1/common"
	"sagooiot/internal/service"
)

var CheckAuth = cCheckAuth{}

type cCheckAuth struct{}

// CheckAccessAuth 验证访问权限
func (c *cCheckAuth) CheckAccessAuth(ctx context.Context, req *common.CheckAccessAuthReq) (res *common.CheckAccessAuthRes, err error) {
	isAllow, err := service.CheckAuth().CheckAccessAuth(ctx, req.Address)
	if err != nil {
		return
	}
	res = &common.CheckAccessAuthRes{
		IsAllow: isAllow,
	}
	return
}

// IsToken 验证token是否正确
func (c *cSysInfo) IsToken(ctx context.Context, req *common.IsTokenReq) (res *common.IsTokenRes, err error) {
	isToken, expiresAt, isAuth, err := service.CheckAuth().IsToken(ctx)
	if err != nil {
		return
	}
	res = &common.IsTokenRes{
		IsToken:   isToken,
		ExpiresAt: expiresAt,
		Auth:      isAuth,
	}
	return
}
