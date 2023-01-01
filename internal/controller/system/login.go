package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/sagoo-cloud/sagooiot/api/v1/system"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

// Login 登录管理
var Login = cLogin{}

type cLogin struct{}

// Login 登录
func (a *cLogin) Login(ctx context.Context, req *system.LoginDoReq) (res *system.LoginDoRes, err error) {
	out, token, err := service.Login().Login(ctx, req.VerifyKey, req.Captcha, req.UserName, req.Password)
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
		UserInfo: loginUserRes,
		Token:    token,
	}
	return
}

// LoginOut 退出登录
func (a *cLogin) LoginOut(ctx context.Context, req *system.LoginOutReq) (res *system.LoginOutRes, err error) {
	err = service.Login().LoginOut(ctx)
	return
}
