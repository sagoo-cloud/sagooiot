package system

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"github.com/tiger1103/gfast-token/gftoken"
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

func (m *sSysToken) GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error) {
	keys, err = GfToken().GenerateToken(ctx, key, data)
	return keys, err
}

func (m *sSysToken) ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error) {
	return GfToken().ParseToken(r)
}

func GfToken() *gftoken.GfToken {
	ctx := gctx.New()
	err := g.Cfg().MustGet(ctx, "gfToken").Struct(&gftService.options)
	if err != nil {
		panic(err.Error())
	}
	prefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	m := g.Cfg().MustGet(ctx, "system.cache.model").String()
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
