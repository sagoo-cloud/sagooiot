package nx

import (
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

type Options struct {
	redis    redis.UniversalClient
	key      string        //redis缓存密钥，默认nx.lock
	expire   int           //密钥过期时间，默认为1分钟，避免死锁，不应设置太长
	retry    int           //重试次数，默认10次
	interval time.Duration //重试间隔，默认25毫秒
}

// Validate 校验 Options 结构体的参数是否合法
func (o *Options) Validate() error {
	if o.redis == nil {
		return errors.New("redis must not be nil")
	}
	return nil
}

// getOptionsOrSetDefault 获取 Options 结构体的指针，如果参数为 nil，则返回默认值
func getOptionsOrSetDefault(options *Options) *Options {
	if options == nil {
		return &Options{
			key:      "nx.lock",
			expire:   60,
			retry:    10,
			interval: 25 * time.Millisecond,
		}
	}
	return options
}

// WithRedis 设置 redis 客户端
func WithRedis(rd redis.UniversalClient) func(*Options) {
	return func(options *Options) {
		if rd != nil {
			getOptionsOrSetDefault(options).redis = rd
		}
	}
}

// WithKey 设置 redis 缓存密钥
func WithKey(key string) func(*Options) {
	return func(options *Options) {
		if key != "" {
			getOptionsOrSetDefault(options).key = key
		}
	}
}

// WithExpire 设置 redis 缓存密钥过期时间
func WithExpire(second int) func(*Options) {
	return func(options *Options) {
		if second > 0 {
			getOptionsOrSetDefault(options).expire = second
		}
	}
}
