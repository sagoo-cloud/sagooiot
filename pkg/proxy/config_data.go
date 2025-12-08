package proxy

import (
	"context"
	"sagooiot/internal/service"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetConfigMapByKeys 通过key数组获取参数（从缓存获取）
func GetConfigMapByKeys(ctx context.Context, keys []string) (out map[string]string, err error) {
	configInfos, err := service.ConfigData().GetConfigByKeys(ctx, keys)
	if err != nil {
		return nil, err
	}
	if err = gconv.Scan(configInfos, &out); err != nil {
		return
	}
	return
}

// GetConfigValueByKey 通过key获取配置值（从缓存获取）
func GetConfigValueByKey(ctx context.Context, key string) (value string, err error) {
	config, err := service.ConfigData().GetByKey(ctx, key)
	if err != nil {
		return
	}
	value = config.ConfigValue
	return
}
