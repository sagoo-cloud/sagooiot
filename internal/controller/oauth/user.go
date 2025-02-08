package oauth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	oauthV1 "sagooiot/api/v1/oauth"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	oauth2 "sagooiot/pkg/oauth"
)

var OUser = cUser{}

type cUser struct{}

// 授权登录
func (c *cUser) Login(ctx context.Context, in *oauthV1.AuthLoginReq) (res *oauthV1.AuthLoginRes, err error) {
	r := g.RequestFromCtx(ctx)
	err = service.OauthProvider().UseProvider(ctx, in.Provider)
	if err != nil {
		return nil, err
	}
	oauth2.BeginAuthHandler(r.Response.RawWriter(), r.Request)
	return
}

// 授权回调
func (c *cUser) Callback(ctx context.Context, req *oauthV1.AuthCallbackReq) (res *oauthV1.AuthCallbackRes, err error) {
	r := g.RequestFromCtx(ctx)
	userInfo, err := oauth2.CompleteUserAuth(r.Response.RawWriter(), r.Request)
	if err != nil {
		return nil, err
	}
	user, err := service.OauthUser().SaveUser(ctx, &entity.OauthUser{
		Nickname:  userInfo.NickName,
		Provider:  req.Provider,
		AvatarUrl: userInfo.AvatarURL,
		Openid:    userInfo.UserID,
	})
	if err != nil {
		return nil, err
	}
	// 登录
	var loginUser *oauthV1.LoginUserRes
	if user.UserId > 0 {
		sysUser, token, err := service.OauthUser().AuthLogin(ctx, user)
		if err != nil {
			return nil, err
		}
		loginUser = &oauthV1.LoginUserRes{
			UserNickname: sysUser.UserNickname,
			Avatar:       sysUser.Avatar,
			Token:        token,
		}
	}
	return &oauthV1.AuthCallbackRes{
		UserInfo:  loginUser,
		OauthUser: user,
	}, err
}

func (c *cUser) ListUser(ctx context.Context, in *oauthV1.GetUserBindingReq) (res *oauthV1.GetUserBindingRes, err error) {
	users, err := service.OauthUser().List(ctx, &entity.OauthUser{
		UserId:   in.UserId,
		Provider: in.Provider,
	})
	if err != nil {
		return nil, err
	}
	return &oauthV1.GetUserBindingRes{
		Users: users,
	}, nil
}

func (c *cUser) UserBinding(ctx context.Context, in *oauthV1.UserBindingReq) (res *oauthV1.UserBindingRes, err error) {
	user, err := service.OauthUser().UserBinding(ctx, in.Provider, in.Openid)
	if err != nil {
		return nil, err
	}
	return &oauthV1.UserBindingRes{
		OauthUser: user,
	}, nil
}
