package command

import (
	"context"
	"fmt"
	"github.com/arl/statsviz"
	"github.com/gogf/gf/v2/frame/g"
	"net/http"
	"os"
	"runtime"
	"syscall"
	"time"
)

func RunSystemAnalysis(stopSignal chan os.Signal, pprofPort string) {
	// 开启性能分析
	runtime.SetMutexProfileFraction(1) // (非必需)开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // (非必需)开启对阻塞操作的跟踪
	// 将Go程序运行时的各种内部数据进行可视化的展示，如可以展示：堆、对象、协程、GC等信息
	err := statsviz.RegisterDefault()
	if err == nil {
		g.Log().Infof(context.Background(), "Point your browser to http://localhost:%s/debug/statsviz/", pprofPort)
		go func() {
			s := &http.Server{
				Addr:         ":" + pprofPort,
				ReadTimeout:  30 * time.Second,
				WriteTimeout: 30 * time.Second,
				IdleTimeout:  30 * time.Second,
			}
			if err := s.ListenAndServe(); err != nil {
				fmt.Println(err)
			}
			stopSignal <- syscall.SIGQUIT
		}()
	}
}
