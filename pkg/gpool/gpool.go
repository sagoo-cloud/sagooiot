package gpool

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type GPool struct {
	sem        chan struct{}
	activeJobs *int32 // 使用原子操作追踪活跃的协程数目
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	errChan    chan error // 错误通道，用于收集任务执行中的错误
}

func NewGPool(capacity int) *GPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &GPool{
		sem:        make(chan struct{}, capacity),
		activeJobs: new(int32),
		ctx:        ctx,
		cancel:     cancel,
		errChan:    make(chan error, capacity), // 错误通道容量设置为工作池的容量
	}
}

func (p *GPool) Acquire() bool {
	select {
	case p.sem <- struct{}{}:
		atomic.AddInt32(p.activeJobs, 1)
		p.wg.Add(1)
		return true
	case <-p.ctx.Done():
		return false // 如果上下文已取消，则不获取并返回false
	}
}

func (p *GPool) Release() {
	<-p.sem
	atomic.AddInt32(p.activeJobs, -1)
	p.wg.Done()
}

func (p *GPool) Go(job func(ctx context.Context) error) {
	if p.Acquire() {
		go func() {
			defer p.Release()
			if err := job(p.ctx); err != nil {
				p.errChan <- err // 将错误发送到错误通道
			}
		}()
	}
}

func (p *GPool) Wait(timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		p.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return true // 所有任务完成
	case <-time.After(timeout):
		return false // 超时
	}
}

func (p *GPool) Shutdown() {
	p.cancel()                    // 发送取消信号
	if !p.Wait(5 * time.Second) { // 设置超时时间为5秒
		// 超时处理逻辑（可根据需要自定义）
	}
	close(p.errChan) // 关闭错误通道
}

// ErrChan 提供一个错误通道的访问方法
func (p *GPool) ErrChan() <-chan error {
	return p.errChan
}
