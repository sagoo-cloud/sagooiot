package oauth

import (
	"context"
	"fmt"
	"sagooiot/internal/consts"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/oauth"
	"sagooiot/pkg/oauth/qq"
	"sagooiot/pkg/oauth/wechat"
	"strconv"
)

type sOauthProvider struct{}

func init() {
	oauthProvider := oauthProviderNew()
	service.RegisterOauthProvider(oauthProvider)
}

func oauthProviderNew() *sOauthProvider {
	return &sOauthProvider{}
}

func useProvider(ctx context.Context, in *entity.OauthProvider) error {
	switch in.Name {
	case "qq":
		oauth.UseProviders(
			qq.New(in.Appid, in.Appsecret, in.CallbackUrl, qq.QQ_LANG_CN),
		)
		return nil
	case "wechat":
		oauth.UseProviders(
			wechat.New(in.Appid, in.Appsecret, in.CallbackUrl, wechat.WECHAT_LANG_CN),
		)
		return nil
	}
	return fmt.Errorf("系统暂不支持 %s 授权登录", in.Name)
}

func (s *sOauthProvider) UseProvider(ctx context.Context, name string) (err error) {
	provider := &entity.OauthProvider{Name: name}
	if name == "wechat" {
		provider, err = getWechatConfig(ctx)
	} else if name == "qq" {
		provider, err = getQqConfig(ctx)
	}
	return useProvider(ctx, provider)
}

func (s *sOauthProvider) List(ctx context.Context, in *entity.OauthProvider) (out []*entity.OauthProvider, err error) {
	qqConfig, _ := getQqConfig(ctx)
	wechatConfig, _ := getWechatConfig(ctx)
	if qqConfig != nil {
		out = append(out, qqConfig)
	}
	if wechatConfig != nil {
		out = append(out, wechatConfig)
	}
	return
}

func getWechatConfig(ctx context.Context) (out *entity.OauthProvider, err error) {
	out = &entity.OauthProvider{}
	cfgConfig, err := service.ConfigData().GetConfigByKey(ctx, consts.OauthWechatName)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Name = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthWechatLogo)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Logo = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthWechatAppid)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Appid = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthWechatAppsecret)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Appsecret = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthWechatStatus)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Status, _ = strconv.Atoi(cfgConfig.ConfigValue)
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthCallbackUrl)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.CallbackUrl = fmt.Sprintf(cfgConfig.ConfigValue, out.Name)
	}
	return

}

func getQqConfig(ctx context.Context) (out *entity.OauthProvider, err error) {
	out = &entity.OauthProvider{}
	cfgConfig, err := service.ConfigData().GetConfigByKey(ctx, consts.OauthQQName)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Name = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthQQLogo)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Logo = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthQQAppid)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Appid = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthQQAppsecret)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Appsecret = cfgConfig.ConfigValue
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthQQStatus)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.Status, _ = strconv.Atoi(cfgConfig.ConfigValue)
	}
	cfgConfig, err = service.ConfigData().GetConfigByKey(ctx, consts.OauthCallbackUrl)
	if err != nil {
		return nil, err
	}
	if cfgConfig != nil {
		out.CallbackUrl = fmt.Sprintf(cfgConfig.ConfigValue, out.Name)
	}
	return
}
