package service

import (
	"context"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
)

var (
	localIOauthUser     IOauthUser
	localIOauthProvider IOauthProvider
)

type IOauthUser interface {
	GetByUserId(ctx context.Context, userId uint64) (out []*entity.OauthUser, err error)
	SaveUser(ctx context.Context, in *entity.OauthUser) (out *entity.OauthUser, err error)
	GetUser(ctx context.Context, provider string, openid string) (out *entity.OauthUser, err error)
	UserBinding(ctx context.Context, provider string, openid string) (out *entity.OauthUser, err error)
	AuthLogin(ctx context.Context, in *entity.OauthUser) (*model.LoginUserOut, string, error)
	List(ctx context.Context, in *entity.OauthUser) ([]*entity.OauthUser, error)
}

type IOauthProvider interface {
	UseProvider(ctx context.Context, name string) error
	List(ctx context.Context, in *entity.OauthProvider) (out []*entity.OauthProvider, err error)
}

func OauthUser() IOauthUser {
	if localIOauthUser == nil {
		panic("implement not found for interface localIOauthUser, forgot register?")
	}
	return localIOauthUser
}

func OauthProvider() IOauthProvider {
	if localIOauthProvider == nil {
		panic("implement not found for interface localIOauthProvider, forgot register?")
	}
	return localIOauthProvider
}

func RegisterOauthUser(i IOauthUser) {
	localIOauthUser = i
}

func RegisterOauthProvider(i IOauthProvider) {
	localIOauthProvider = i
}
