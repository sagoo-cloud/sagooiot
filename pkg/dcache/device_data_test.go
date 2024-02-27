package dcache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"log"
	"sagooiot/pkg/iotModel"
	"testing"
	"time"
)

// TestInsertData 插入数据示例
func TestInsertBatchData(t *testing.T) {
	ctx := context.Background()
	//manager := GetRedisManager("localhost:6379")
	// 插入数据示例
	// 创建DeviceLogData类型的数据
	var logList []interface{}
	for i := 0; i < 3000+1; i++ {
		devLog := iotModel.DeviceLog{
			Ts:      gtime.Now(),
			Device:  "device001",
			Type:    "info",
			Content: fmt.Sprintf("Log message %d", i),
		}
		logList = append(logList, devLog)
	}

	if err := DB().InsertBatchData(ctx, "mylist222", logList); err != nil {
		t.Error("Error inserting data:", err)
		return
	}
}

// TestInsertData 插入数据示例
func TestInsertData(t *testing.T) {
	data := map[string]interface{}{
		"name": grand.S(6),
		"age":  grand.Intn(100),
		"time": time.Now(),
	}
	t.Log(data)
	// 插入单条数据
	if err := DB().InsertData(context.Background(), "userKey", data); err != nil {
		log.Printf("Failed to insert data: %s\n", err)
	}
}

// TestGetData 获取数据示例
func TestGetData(t *testing.T) {
	ctx := context.Background()
	// 获取数据示例
	result, err := DB().GetData(ctx, "mylist222")
	if err != nil {
		t.Error("Error getting data:", err)
		return
	}
	var logs []iotModel.DeviceLog
	if err := gconv.Scan(result, &logs); err != nil {
		t.Error(err)
	}
	t.Log("记录数：", len(logs))
	for _, devLog := range logs {
		fmt.Printf("Device: %s, Type: %s, Content: %s, Timestamp: %v\n", devLog.Device, devLog.Type, devLog.Content, devLog.Ts)
	}

}

// TestListenForNewData 监听新数据示例
func TestListenForNewData(t *testing.T) {
	// 创建一个可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())
	// 启动协程以监听 Redis key, 1秒钟检查一次
	var tmpData []interface{}
	go DB().ListenForNewData(ctx, "userKey", func(data string) {
		fmt.Println("处理新数据：", data)
		tmpData = append(tmpData, data)

	}, 1*time.Second)

	//延迟落库处理
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for {
			select {
			case <-ctx.Done():
				//fmt.Println("延迟落库处理监听结束")
				return
			case <-ticker.C:
				//fmt.Println("=====监听中=========")
				//如果数据为空退出
				if len(tmpData) == 0 {
					continue
				}
				fmt.Println(tmpData)
				tmpData = nil
			}
		}
	}()

	// 做一些操作，例如等待用户输入或达到某个条件
	time.Sleep(1600 * time.Second) // 示例：等待10秒

	// 当需要停止监听时，调用 cancel 函数
	cancel()
}

// TestGetDeviceDetailData 获取设备数据示例
func TestGetDeviceDetailData(t *testing.T) {
	// 获取数据示例
	resultList := GetDeviceDetailData(context.Background(), "t20221222")
	for _, value := range resultList {
		t.Log(value)

	}
	t.Log("记录数：", len(resultList))
}

// TestGetDeviceDetailDataByPage 获取设备数据示例
func TestGetDeviceDetailDataByPage(t *testing.T) {
	dataList, total, currentPage := GetDeviceDetailDataByPage(context.Background(), "t20221222", 2, 10)
	for _, d := range dataList {
		t.Log(d)
	}
	t.Log("记录数：", len(dataList), total, currentPage)
}
