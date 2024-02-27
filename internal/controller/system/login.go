package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"sagooiot/api/v1/system"
	"sagooiot/internal/model"
	"sagooiot/internal/service"
)

// Login 登录管理
var Login = cLogin{}

type cLogin struct{}

// Login 登录
func (a *cLogin) Login(ctx context.Context, req *system.LoginDoReq) (res *system.LoginDoRes, err error) {

	out, token, isChangePassword, err := service.Login().Login(ctx, req.VerifyKey, req.Captcha, req.UserName, req.Password)
	if err != nil {
		return
	}
	var loginUserRes *model.LoginUserRes
	if out != nil {
		if err = gconv.Scan(out, &loginUserRes); err != nil {
			return
		}
	}
	res = &system.LoginDoRes{
		UserInfo:    loginUserRes,
		Token:       token,
		IsChangePwd: isChangePassword,
	}
	return
}

// LoginOut 退出登录
func (a *cLogin) LoginOut(ctx context.Context, req *system.LoginOutReq) (res *system.LoginOutRes, err error) {
	err = service.Login().LoginOut(ctx)
	return
}

// EditPassword 修改密码
func (a *cLogin) EditPassword(ctx context.Context, req *system.EditPasswordReq) (res *system.EditPasswordRes, err error) {
	err = service.SysUser().EditPassword(ctx, req.UserName, req.OldUserPassword, req.UserPassword)
	return
}
