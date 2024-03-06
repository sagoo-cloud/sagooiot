package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"sagooiot/internal/consts"
	"sagooiot/internal/model"
	"sagooiot/internal/model/entity"
	"sagooiot/internal/service"
	"sagooiot/pkg/gftoken"
	"strconv"
	"strings"
	"sync"
)

type sSysToken struct {
}

func init() {
	service.RegisterSysToken(sysTokenNew())
}

func sysTokenNew() *sSysToken {
	return &sSysToken{}
}

type gft struct {
	options *model.TokenOptions
	lock    *sync.Mutex
}

var gftService = &gft{
	options: nil,
	lock:    &sync.Mutex{},
}

func (s *sSysToken) GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error) {
	keys, err = s.GfToken().GenerateToken(ctx, key, data)
	return keys, err
}

func (s *sSysToken) ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error) {
	return s.GfToken().ParseToken(r)
}

func (s *sSysToken) GfToken() *gftoken.GfToken {
	ctx := gctx.New()

	//判断控制是否生效
	configDataByIsSecurityControlEnabled, err := service.ConfigData().GetConfigByKey(ctx, consts.SysIsSecurityControlEnabled)
	if err != nil {
		panic(err.Error())
	}
	var configDataByTokenExpiryDate *entity.SysConfig
	var configDataBySingleLogin *entity.SysConfig
	if configDataByIsSecurityControlEnabled != nil && strings.EqualFold(configDataByIsSecurityControlEnabled.ConfigValue, "1") {
		//获取token过期时间
		configDataByTokenExpiryDate, err = service.ConfigData().GetConfigByKey(ctx, consts.SysTokenExpiryDate)
		if err != nil {
			panic(err.Error())
		}

		//获取是否单一登录系统参数
		configDataBySingleLogin, err = service.ConfigData().GetConfigByKey(ctx, consts.SysIsSingleLogin)
		if err != nil {
			panic(err.Error())
		}
	}

	err = g.Cfg().MustGet(ctx, "gfToken").Struct(&gftService.options)
	if err != nil {
		panic(err.Error())
	}

	//设置token过期时间
	if configDataByTokenExpiryDate != nil {
		// 将字符串转换为整数
		var minutes int
		minutes, err = strconv.Atoi(configDataByTokenExpiryDate.ConfigValue)
		if err != nil {
			glog.Debugf(ctx, "无法将字符串转换为整数:%s", err)
			panic(err.Error())
		}
		gftService.options.Timeout = int64(minutes * 60)
	}
	//是否开启单一登录
	if configDataBySingleLogin != nil {
		if strings.EqualFold(configDataBySingleLogin.ConfigValue, "0") {
			gftService.options.MultiLogin = true //允许一个账户多设备登录
		} else if strings.EqualFold(configDataBySingleLogin.ConfigValue, "1") {
			gftService.options.MultiLogin = false //禁止一个账户多设备登录
		}
	}
	prefix := g.Cfg().MustGet(ctx, "cache.prefix").String()
	m := g.Cfg().MustGet(ctx, "cache.adapter").String()
	return GfTokenOption(gftService.options, prefix, m)
}

func GfTokenOption(options *model.TokenOptions, prefix string, model string) *gftoken.GfToken {
	var fun gftoken.OptionFunc
	if model == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}

	gfToken := gftoken.NewGfToken(
		gftoken.WithCacheKey(prefix),
		gftoken.WithTimeout(options.Timeout),
		gftoken.WithMaxRefresh(options.MaxRefresh),
		gftoken.WithMultiLogin(options.MultiLogin),
		gftoken.WithExcludePaths(options.ExcludePaths),
		fun,
	)
	return gfToken
}
