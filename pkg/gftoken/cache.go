package gftoken

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

func (m *GfToken) contains(ctx context.Context, key string) bool {
	ok, _ := m.cache.Contains(ctx, key)
	return ok
}

func (m *GfToken) setCache(ctx context.Context, key string, value interface{}) error {
	return m.cache.Set(ctx, key, value, time.Duration(m.Timeout+m.MaxRefresh)*time.Second)
}

func (m *GfToken) getCache(ctx context.Context, key string) (tData *TokenData, err error) {
	var result *gvar.Var
	result, err = m.cache.Get(ctx, key)
	if err != nil {
		return
	}
	if result.Val() != nil {
		err = gconv.Struct(result, &tData)
	}
	return
}

func (m *GfToken) removeCache(ctx context.Context, key string) (err error) {
	_, err = m.cache.Remove(ctx, key)
	return
}
