package file

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type (
	// AdapterFile gcache适配器使用文件的实现
	AdapterFile struct {
		dir   string
		mutex sync.RWMutex // 添加互斥锁以支持并发操作
	}

	fileContent struct {
		Duration int64       `json:"duration"`
		Data     interface{} `json:"data,omitempty"`
	}
)

const perm = 0o666 //设置文件所有者、同用户组成员和其他用户都可以读写该文件，但都不能执。

var (
	CacheExpiredErr = errors.New("cache expired")
)

// NewAdapterFile creates and returns a new memory cache object.
func NewAdapterFile(dir string) gcache.Adapter {
	return &AdapterFile{
		dir: dir,
	}
}

func (c *AdapterFile) Set(ctx context.Context, key interface{}, value interface{}, lifeTime time.Duration) (err error) {
	c.mutex.Lock()         // 加锁
	defer c.mutex.Unlock() // 解锁
	fileKey := gconv.String(key)
	if value == nil || lifeTime < 0 {
		return c.Delete(fileKey)
	}
	return c.Save(fileKey, gconv.String(value), lifeTime)
}

// SetMap 批量设置多个键值对，并且每个键值对的有效期是一样的
func (c *AdapterFile) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) (err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, value := range data {
		fileKey := gconv.String(key)
		if err = c.Save(fileKey, gconv.String(value), duration); err != nil {
			return err
		}
	}
	return nil
}

// SetIfNotExist 在指定的键不存在时设置其值
func (c *AdapterFile) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fileKey := gconv.String(key)
	if c.Has(fileKey) {
		return false, nil
	}

	err = c.Save(fileKey, gconv.String(value), duration)
	return true, err
}

// SetIfNotExistFunc 在指定的键不存在时，通过一个函数来生成值并设置它
func (c *AdapterFile) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fileKey := gconv.String(key)
	if c.Has(fileKey) {
		return false, nil
	}

	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	err = c.Save(fileKey, gconv.String(value), duration)
	return true, err
}

// SetIfNotExistFuncLock 与 SetIfNotExistFunc 类似，但它在调用生成值的函数时提供了额外的锁机制，以避免在生成值期间的并发问题。
func (c *AdapterFile) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fileKey := gconv.String(key)
	if c.Has(fileKey) {
		return false, nil
	}

	// 在这里加锁是为了确保在值生成期间不会有并发的写入操作
	value, err := f(ctx)
	if err != nil {
		return false, err
	}
	err = c.Save(fileKey, gconv.String(value), duration)
	return true, err
}

func (c *AdapterFile) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	fetch, err := c.Fetch(gconv.String(key))
	if err != nil {
		return nil, err
	}
	return gvar.New(fetch), nil
}

func (c *AdapterFile) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	result, err = c.Get(ctx, key)
	if err != nil && !errors.Is(err, CacheExpiredErr) {
		return nil, err
	}
	if result.IsNil() {
		return gvar.New(value), c.Set(ctx, key, value, duration)
	}
	return
}

func (c *AdapterFile) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	v, err := c.Get(ctx, key)
	if err != nil && !errors.Is(err, CacheExpiredErr) {
		return nil, err
	}
	if v.IsNil() {
		value, err := f(ctx)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil
		}
		return gvar.New(value), c.Set(ctx, key, value, duration)
	} else {
		return v, nil
	}
}

func (c *AdapterFile) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	return c.GetOrSetFunc(ctx, key, f, duration)
}

func (c *AdapterFile) Contains(ctx context.Context, key interface{}) (bool, error) {
	return c.Has(gconv.String(key)), nil
}

func (c *AdapterFile) Size(ctx context.Context) (size int, err error) {
	return 0, nil
}

// Data 获取所有缓存的键值对
func (c *AdapterFile) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// 初始化返回的数据结构
	data = make(map[interface{}]interface{})

	// 读取文件目录中的所有缓存文件
	files, err := os.ReadDir(c.dir)
	if err != nil {
		return nil, err
	}

	// 遍历缓存文件，提取数据
	for _, file := range files {
		content, err := c.read(file.Name())
		if err != nil {
			continue // 忽略无法读取的文件
		}
		data[file.Name()] = content.Data
	}
	return data, nil
}

// Keys 获取所有缓存的键
func (c *AdapterFile) Keys(ctx context.Context) (keys []interface{}, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	files, err := os.ReadDir(c.dir)
	if err != nil {
		return nil, err
	}

	keys = make([]interface{}, 0, len(files))
	for _, file := range files {
		keys = append(keys, file.Name())
	}
	return keys, nil
}

// Values 获取所有缓存的值
func (c *AdapterFile) Values(ctx context.Context) (values []interface{}, err error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	files, err := os.ReadDir(c.dir)
	if err != nil {
		return nil, err
	}

	values = make([]interface{}, 0, len(files))
	for _, file := range files {
		content, err := c.read(file.Name())
		if err != nil {
			continue // 忽略无法读取的文件
		}
		values = append(values, content.Data)
	}
	return values, nil
}

// Update 更新一个指定键的缓存值，如果这个键不存在，则返回错误
func (c *AdapterFile) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fileKey := gconv.String(key)
	content, err := c.read(fileKey)
	if err != nil {
		return nil, false, err
	}
	if content == nil {
		return nil, false, nil
	}

	oldValue = gvar.New(content.Data)
	err = c.Save(fileKey, gconv.String(value), time.Duration(content.Duration)*time.Second)
	return oldValue, true, err
}

func (c *AdapterFile) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	var (
		v       *gvar.Var
		oldTTL  int64
		fileKey = gconv.String(key)
	)
	// TTL.
	expire, err := c.GetExpire(ctx, fileKey)
	if err != nil {
		return
	}
	oldTTL = int64(expire)
	if oldTTL == -2 {
		// It does not exist.
		return
	}
	oldDuration = time.Duration(oldTTL) * time.Second
	// DEL.
	if duration < 0 {
		err = c.Delete(fileKey)
		return
	}
	v, err = c.Get(ctx, fileKey)
	if err != nil {
		return
	}
	err = c.Set(ctx, fileKey, v.Val(), duration)
	return
}

func (c *AdapterFile) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	content, err := c.read(gconv.String(key))
	if err != nil {
		return -1, nil
	}

	if content.Duration <= time.Now().Unix() {
		return -1, nil
	}
	return time.Duration(time.Now().Unix()-content.Duration) * time.Second, nil
}

func (c *AdapterFile) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	if len(keys) == 0 {
		return nil, nil
	}
	// Retrieves the last key value.
	if lastValue, err = c.Get(ctx, gconv.String(keys[len(keys)-1])); err != nil {
		return nil, err
	}
	// Deletes all given keys.
	err = c.DeleteMulti(gconv.Strings(keys)...)
	return
}

func (c *AdapterFile) Clear(ctx context.Context) error {
	return c.Flush()
}

func (c *AdapterFile) Close(ctx context.Context) error {
	return nil
}

func (c *AdapterFile) createName(key string) string {
	h := sha256.New()
	_, _ = h.Write([]byte(key))
	hash := hex.EncodeToString(h.Sum(nil))
	return filepath.Join(c.dir, fmt.Sprintf("%s.cache", hash))
}

func (c *AdapterFile) read(key string) (*fileContent, error) {
	rp := gfile.RealPath(c.createName(key))
	if rp == "" {
		return nil, nil
	}

	value, err := os.ReadFile(rp)
	if err != nil {
		return nil, err
	}

	content := &fileContent{}
	if err := json.Unmarshal(value, content); err != nil {
		return nil, err
	}

	if content.Duration == 0 {
		return content, nil
	}

	if content.Duration <= time.Now().Unix() {
		_ = c.Delete(key)
		return nil, CacheExpiredErr
	}
	return content, nil
}

// Has checks if the cached key exists into the File storage
func (c *AdapterFile) Has(key string) bool {
	fc, err := c.read(key)
	return err == nil && fc != nil
}

// Delete the cached key from File storage
func (c *AdapterFile) Delete(key string) error {
	_, err := os.Stat(c.createName(key))
	if err != nil && os.IsNotExist(err) {
		return nil
	}
	return os.Remove(c.createName(key))
}

// DeleteMulti the cached key from File storage
func (c *AdapterFile) DeleteMulti(keys ...string) (err error) {
	for _, key := range keys {
		if err = c.Delete(key); err != nil {
			return
		}
	}
	return
}

// Fetch retrieves the cached value from key of the File storage
func (c *AdapterFile) Fetch(key string) (interface{}, error) {
	content, err := c.read(key)
	if err != nil {
		return nil, err
	}

	if content == nil {
		return nil, nil
	}
	return content.Data, nil
}

// FetchMulti retrieve multiple cached values from keys of the File storage
func (c *AdapterFile) FetchMulti(keys []string) map[string]interface{} {
	result := make(map[string]interface{})
	for _, key := range keys {
		if value, err := c.Fetch(key); err == nil {
			result[key] = value
		}
	}
	return result
}

// Flush removes all cached keys of the File storage
func (c *AdapterFile) Flush() error {
	dir, err := os.Open(c.dir)
	if err != nil {
		return err
	}

	defer func() {
		_ = dir.Close()
	}()

	names, _ := dir.Readdirnames(-1)

	for _, name := range names {
		_ = os.Remove(filepath.Join(c.dir, name))
	}
	return nil
}

// Save a value in File storage by key
func (c *AdapterFile) Save(key string, value string, lifeTime time.Duration) error {
	duration := int64(0)

	if lifeTime > 0 {
		duration = time.Now().Unix() + int64(lifeTime.Seconds())
	}

	content := &fileContent{duration, value}

	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	err = os.WriteFile(c.createName(key), data, perm)
	return err
}
