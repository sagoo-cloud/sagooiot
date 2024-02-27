package statistics

import (
	"context"
	"sagooiot/internal/consts"
	"sagooiot/pkg/cache"
	"sagooiot/pkg/utility/utils"
	"sync"
	"sync/atomic"
	"time"
)

var (
	count     int64
	lastDate  string
	dateMutex sync.Mutex
)

// CountDeviceData 计数设备数据
func CountDeviceData() {
	currentDateString := utils.GetCurrentDateString()
	dateMutex.Lock()
	if lastDate != currentDateString {
		lastDate = currentDateString
		count = 0
		saveCountToDB(0)
	}
	dateMutex.Unlock()

	atomic.AddInt64(&count, 1)
}

// GetCurrentDeviceDataCount 获取当前计数
func GetCurrentDeviceDataCount() int64 {
	return atomic.LoadInt64(&count)
}

// 持久化当前计数
func saveCountToDB(currentCount int64) {

	err := cache.Instance().Set(context.Background(), consts.AnalysisDeviceCountPrefix+consts.TodayMessageVolume+utils.GetTimeTagGroup()+lastDate, currentCount, 0)
	if err != nil {
		return
	}
}

// InitCountDeviceData 初始化设备数据计数
func InitCountDeviceData() {
	lastDate = utils.GetCurrentDateString()
	// 恢复之前的计数值
	res, err := cache.Instance().Get(context.Background(), consts.AnalysisDeviceCountPrefix+consts.TodayMessageVolume+utils.GetTimeTagGroup()+lastDate)
	if err != nil {
		return
	}
	count = res.Int64()
	// 设置定时持久化一次
	ticker := time.NewTicker(3 * time.Second)
	go func() {
		for range ticker.C {
			currentCount := atomic.LoadInt64(&count)
			saveCountToDB(currentCount) // 持久化当前计数
		}
	}()
}
