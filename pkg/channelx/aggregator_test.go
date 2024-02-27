package channelx

import (
	"log"
	"os"
	"testing"
	"time"
)

// mockBatchProcessFunc 是一个模拟的批处理函数，用于测试
func mockBatchProcessFunc(items []interface{}) error {
	// 在这里简单模拟处理逻辑，实际使用中应替换为具体的业务逻辑
	log.Println("批量处理:", items)
	return nil
}

// TestAggregator 测试Aggregator的基本行为
func TestAggregator(t *testing.T) {
	logger := log.New(os.Stdout, "AggregatorTest: ", log.Lshortfile)
	batchSize := 5 // 批处理大小
	workers := 5
	channelBufferSize := 5000 // 通道缓冲区大小
	lingerTime := 100 * time.Millisecond

	// 创建聚合器实例
	aggregator := NewAggregator(
		mockBatchProcessFunc,
		WithBatchSize(batchSize),
		WithWorkers(workers),
		WithChannelBufferSize(channelBufferSize),
		WithLingerTime(lingerTime),
		WithLogger(logger),
	)

	// 开始聚合器
	aggregator.Start()

	// 模拟入队操作
	itemsToEnqueue := 2000
	for i := 0; i < itemsToEnqueue+1; i++ {
		if !aggregator.TryEnqueue(i) {
			t.Logf("Failed to enqueue item: %d", i)
		}
	}

	// 等待足够的时间，以确保所有项目都被处理
	time.Sleep(2 * time.Second)

	// 安全停止聚合器
	aggregator.SafeStop()

}
