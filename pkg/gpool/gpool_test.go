package gpool

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	// 初始化一个容量为5的工作池
	pool := NewGPool(5)

	// 模拟10个任务
	for i := 0; i < 10; i++ {
		jobIndex := i
		pool.Go(func(ctx context.Context) error {
			// 随机等待时间，模拟任务执行
			waitTime := time.Duration(rand.Intn(1000)) * time.Millisecond
			select {
			case <-time.After(waitTime):
				// 模拟任务有一定几率失败
				if rand.Float32() < 0.3 {
					return fmt.Errorf("job %d failed", jobIndex)
				}
				fmt.Printf("job %d completed successfully\n", jobIndex)
				return nil
			case <-ctx.Done():
				// 上下文被取消，任务终止
				fmt.Printf("job %d cancelled\n", jobIndex)
				return ctx.Err()
			}
		})
	}

	// 等待所有任务完成或超时
	if !pool.Wait(3 * time.Second) {
		t.Log("Waited for 3 seconds, not all jobs completed")
	}

	// 关闭工作池并处理错误
	pool.Shutdown()

	// 从错误通道处理和打印错误
	for err := range pool.ErrChan() {
		t.Log("Received error:", err)
	}
}
