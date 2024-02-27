package cache

import (
	"context"
	"sagooiot/pkg/cache/file"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
)

// cache 缓存驱动
var cache *gcache.Cache

// Instance 缓存实例
func Instance() *gcache.Cache {
	if cache == nil {
		panic("cache uninitialized.")
	}
	return cache
}

// SetAdapter 设置缓存适配器
func SetAdapter(ctx context.Context) {
	var cacheAdapter gcache.Adapter
	adapter := g.Cfg().MustGet(ctx, "cache.adapter").String()
	fileDir := g.Cfg().MustGet(ctx, "cache.fileDir").String()

	switch adapter {
	case "redis":
		cacheAdapter = gcache.NewAdapterRedis(g.Redis())
	case "file":
		if fileDir == "" {
			g.Log().Fatal(ctx, "file path must be configured for file caching.")
			return
		}

		if !gfile.Exists(fileDir) {
			if err := gfile.Mkdir(fileDir); err != nil {
				g.Log().Fatalf(ctx, "failed to create the cache directory. procedure, err:%+v", err)
				return
			}
		}
		cacheAdapter = file.NewAdapterFile(fileDir)
	default:
		cacheAdapter = gcache.NewAdapterMemory()
	}

	// 数据库缓存，默认和通用缓冲驱动一致，如果你不想使用默认的，可以自行调整
	g.DB().GetCache().SetAdapter(cacheAdapter)

	// 通用缓存
	cache = gcache.New()
	cache.SetAdapter(cacheAdapter)
}
