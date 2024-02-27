package dcache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
	"time"
)

const DeviceDataCachePrefix = "deviceCacheData:"

// RedisManager 管理Redis操作和连接池
type RedisManager struct {
	client         *redis.Client
	recordDuration time.Duration // 记录的有效时间
	recordLimit    int64         // 记录的条数限制
}

type redisOptions struct {
	Addr           string // Redis服务器地址
	DB             int    // Redis数据库
	Password       string // Redis密码
	PoolSize       int    // 连接池大小
	RecordDuration string // 记录的有效时间
	RecordLimit    int64  // 记录的条数限制
}

var (
	managerInstance *RedisManager
	once            sync.Once
)

// DataProcessor 是一个处理新数据的函数类型
type DataProcessor func(data string)

// DB 用于全局获取RedisManager实例
func DB() *RedisManager {
	once.Do(func() {
		// 从配置文件获取redis的ip以及db
		address := g.Cfg().MustGet(context.Background(), "redis.default.address", "localhost:6379").String()
		db := g.Cfg().MustGet(context.Background(), "redis.default.db", "0").Int()
		password := g.Cfg().MustGet(context.Background(), "redis.default.pass", "").String()
		poolSize := g.Cfg().MustGet(context.Background(), "system.deviceCacheData.poolSize", "500").Int()
		recordDuration := g.Cfg().MustGet(context.Background(), "system.deviceCacheData.recordDuration", "10m").String()
		recordLimit := g.Cfg().MustGet(context.Background(), "system.deviceCacheData.recordLimit", "1000").Int64()

		// 从配置文件获取redis的ip以及db
		options := redisOptions{
			Addr:           address,
			DB:             db,
			Password:       password,       // 如果需要密码
			PoolSize:       poolSize,       // 设置连接池大小
			RecordDuration: recordDuration, // 记录的有效时间
			RecordLimit:    recordLimit,    // 记录的条数限制
		}
		managerInstance = getRedisManager(options)
	})
	return managerInstance

}

// GetRedisManager 单例模式获取RedisManager实例
func getRedisManager(options redisOptions) *RedisManager {
	var client *redis.Client
	var err error

	recordDuration, err := time.ParseDuration(options.RecordDuration)
	if err != nil {
		recordDuration = 10 * time.Minute // 或设置一个默认值
		fmt.Println("Invalid RecordDuration format, setting to default:", recordDuration)
	}

	for { // 按秒持续连接尝试
		client = redis.NewClient(&redis.Options{
			Addr:       options.Addr,
			DB:         options.DB,
			Password:   options.Password,
			PoolSize:   options.PoolSize,
			ClientName: "DeviceData", // 设置连接名称
		})
		// 尝试 Ping 操作以检查连接
		_, err = client.Ping(context.Background()).Result()
		if err == nil {
			break // 如果成功，则中断循环
		}

		// 如果连接失败，等待一段时间后重试
		fmt.Printf("Failed to connect to Redis: %v. Retrying in 1 second...\n", err)
		time.Sleep(1 * time.Second)
	}

	// 如果经过重试后仍然失败，处理错误（或退出程序）
	if err != nil {
		// 这里可以记录错误、返回nil或退出程序
		fmt.Printf("Failed to connect to Redis after retries: %v\n", err)
		return nil
	}

	return &RedisManager{
		client:         client,
		recordDuration: recordDuration,
		recordLimit:    options.RecordLimit,
	}
}

// GetClient 获取Redis客户端
func (r *RedisManager) GetClient() *redis.Client {
	return r.client
}

// InsertBatchData 批量插入数据
func (r *RedisManager) InsertBatchData(ctx context.Context, key string, data []interface{}) error {
	pipe := r.client.Pipeline()
	for _, item := range data {
		serializedValue, err := json.Marshal(item)
		if err != nil {
			return err
		}
		pipe.LPush(ctx, DeviceDataCachePrefix+key, serializedValue)
	}

	// 设置列表长度限制
	pipe.LTrim(ctx, DeviceDataCachePrefix+key, 0, r.recordLimit-1)

	// 设置过期时间
	pipe.Expire(ctx, DeviceDataCachePrefix+key, r.recordDuration)

	// 执行所有命令
	_, err := pipe.Exec(ctx)
	return err
}

// InsertData 插入单条数据
func (r *RedisManager) InsertData(ctx context.Context, key string, data interface{}) error {
	serializedValue, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = r.client.LPush(ctx, DeviceDataCachePrefix+key, serializedValue).Err()
	if err != nil {
		return err
	}

	// 设置列表长度限制
	r.client.LTrim(ctx, DeviceDataCachePrefix+key, 0, r.recordLimit-1)

	// 设置过期时间
	r.client.Expire(ctx, DeviceDataCachePrefix+key, r.recordDuration)

	return nil
}

// GetData 获取最新的数据
func (r *RedisManager) GetData(ctx context.Context, key string) ([]string, error) {
	return r.client.LRange(ctx, DeviceDataCachePrefix+key, 0, r.recordLimit-1).Result()
}

// GetDataByLatest 获取最新的一条数据
func (r *RedisManager) GetDataByLatest(ctx context.Context, key string) (string, error) {
	return r.client.LIndex(ctx, DeviceDataCachePrefix+key, -1).Result()
}

// GetDataByPage 按分页获取数据，增加按字段内容搜索和时间区间搜索
func GetDataByPage(ctx context.Context, deviceKey string, pageNum, pageSize int, types, dateRange []string) (res []string, total, currentPage int, err error) {
	listName := DeviceDataCachePrefix + deviceKey

	// 先获取全部数据
	allData, err := DB().client.LRange(ctx, listName, 0, -1).Result()
	if err != nil {
		fmt.Println(err)
		return
	}

	if dateRange == nil {
		dateRange = make([]string, 2)
		dateRange[0] = "1970-01-01"
		dateRange[1] = time.Now().Format("2006-01-02")
	}

	startDate, endDate := parseDateRange(dateRange)

	var tmpDataList []string
	for _, item := range allData {
		if matchesTypes(item, types) && inDateRange(item, startDate, endDate) {
			tmpDataList = append(tmpDataList, item)
		}
	}

	total = len(tmpDataList)

	if pageNum <= 0 {
		pageNum = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}
	currentPage = pageNum
	res, _ = getPage(tmpDataList, pageNum, pageSize)

	return
}

// getPage 进行分页
func getPage(data []string, pageNum, pageSize int) ([]string, int) {
	if pageNum < 1 || pageSize < 1 {
		return []string{}, 0
	}
	// 计算总页数
	total := len(data) / pageSize
	if len(data)%pageSize != 0 {
		total += 1
	}

	// 检查页码是否超出范围
	if pageNum > total {
		return []string{}, total
	}

	// 计算分页的起始位置和结束位置
	start := (pageNum - 1) * pageSize
	end := start + pageSize

	// 调整结束位置以防止越界
	if end > len(data) {
		end = len(data)
	}

	// 截取分页数据
	res := data[start:end]

	return res, total
}

// matchesTypes 检查数据是否匹配指定的类型
func matchesTypes(data string, types []string) bool {
	for _, t := range types {
		if strings.Contains(data, t) {
			return true
		}
	}
	return len(types) == 0
}

// inDateRange 检查数据是否在指定的日期范围内
func inDateRange(data string, startDate, endDate time.Time) bool {
	// 这里假设数据中包含可解析的日期格式
	// 实际应根据数据格式进行调整
	// 例如: "2021-02-01 15:04:05 - Some data"
	dateStr := strings.Split(data, " - ")[0]
	dataDate, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return true
	}

	return (startDate.IsZero() || dataDate.After(startDate)) && (endDate.IsZero() || dataDate.Before(endDate))
}

// parseDateRange 解析日期范围
func parseDateRange(dateRange []string) (startDate, endDate time.Time) {
	if len(dateRange) >= 2 {
		startDate, _ = time.Parse("2006-01-02", dateRange[0])
		endDate, _ = time.Parse("2006-01-02", dateRange[1])
	}
	return
}

// ListenForNewData 监听指定的 Redis key，对新数据执行处理函数，interval为轮询间隔
func (r *RedisManager) ListenForNewData(ctx context.Context, key string, processor DataProcessor, interval time.Duration) {
	var lastCheckedSize int64 = 0
	for {
		select {
		case <-ctx.Done():
			//fmt.Println("监听结束")
			return
		default:
			// 获取当前 list 的大小
			currentSize, err := r.client.LLen(ctx, DeviceDataCachePrefix+key).Result()
			if err != nil {
				time.Sleep(time.Second) // 简单的错误恢复
				continue
			}

			if currentSize > lastCheckedSize {
				// 只检查自上次以来的新元素
				newElements, err := r.client.LRange(ctx, DeviceDataCachePrefix+key, 0, currentSize-lastCheckedSize-1).Result()
				if err != nil {
					fmt.Println("获取新元素失败:", err)
					time.Sleep(time.Second) // 简单的错误恢复
					continue
				}

				// 使用传入的处理函数处理新元素
				for _, element := range newElements {
					processor(element)
				}

				// 更新上次检查的位置
				lastCheckedSize = currentSize
			}

			// 休眠以减少轮询频率
			time.Sleep(interval)
		}
	}
}
