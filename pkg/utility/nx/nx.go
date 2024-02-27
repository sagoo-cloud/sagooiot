package nx

import (
	"context"
	"errors"
	"time"
)

type Nx struct {
	ops   Options
	valid bool
}

// New 创建一个新的 nx 锁实例
func New(options ...func(*Options)) (nx *Nx) {
	ops := getOptionsOrSetDefault(nil)
	for _, f := range options {
		f(ops)
	}
	nx = &Nx{
		ops:   *ops,
		valid: ops.redis != nil,
	}
	if err := ops.Validate(); err != nil {
		return
	}
	return
}

// MustLock 尝试获取锁，如果获取失败，则在 interval 秒内自动重试 retry 次以获得锁定，如果失败，则返回错误
func (nx Nx) MustLock(c ...context.Context) (err error) {
	if !nx.valid {
		return
	}
	ctx := context.Background()
	if len(c) > 0 {
		ctx = c[0]
	}
	var attempts int
	for {
		ok, err := nx.tryLock(ctx)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		//MustLock将在 interval 秒内自动重试 retry 次以获得锁定，如果失败，则返回错误
		if attempts >= nx.ops.retry {
			err = errors.New("lock timeout")
			return err
		}
		time.Sleep(nx.ops.interval)
		attempts++
	}
}

// tryLock 尝试获取锁。如果获取成功，则返回 true，否则返回 false
func (nx Nx) tryLock(ctx context.Context) (bool, error) {
	cmd := nx.ops.redis.SetNX(ctx, nx.ops.key, 1, time.Duration(nx.ops.expire)*time.Second)
	_, err := cmd.Result()
	if err != nil {
		if err == nil || err.Error() == "redis: nil" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Lock 尝试获取锁。如果获取成功，则返回 true，否则返回 false
func (nx Nx) Lock(c ...context.Context) (ok bool) {
	if !nx.valid {
		return
	}
	ctx := context.Background()
	if len(c) > 0 {
		ctx = c[0]
	}
	ok, _ = nx.tryLock(ctx)
	return
}

// Unlock 释放锁
func (nx Nx) Unlock(ctx ...context.Context) {
	if !nx.valid {
		return
	}
	c := context.Background()
	if len(ctx) > 0 {
		c = ctx[0]
	}
	nx.ops.redis.Del(c, nx.ops.key)
	return
}
