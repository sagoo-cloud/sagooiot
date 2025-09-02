package command

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sagooiot/internal/sse"
	"syscall"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func runMain(ctx context.Context, s *ghttp.Server) error {
	var signalChannel = make(chan os.Signal, 1)
	enablePProf := g.Cfg().MustGet(context.Background(), "system.enablePProf").Bool()
	if enablePProf {
		pprofPort := g.Cfg().MustGet(context.Background(), "system.pprofPort").String()
		if pprofPort == "" {
			pprofPort = "58089"
		}
		RunSystemAnalysis(signalChannel, pprofPort) // 运行系统分析
	}

	deferFuncListIotCore, err := InitSystemDeferFunc(ctx)
	defer func() {
		for _, f := range deferFuncListIotCore {
			if f == nil {
				continue
			}
			if deferErr := f(ctx); deferErr != nil {
				fmt.Printf("defer func error: %s\n", deferErr.Error())
			}
		}
	}()

	err = InitSystem(ctx, InitFuncNoDeferListForIotCore)
	if err != nil {
		fmt.Printf("defer func error: %s\n", err.Error())
	}

	err = InitSystem(ctx, InitFuncNoDeferListWebAdmin)
	if err != nil {
		fmt.Printf("defer func error: %s\n", err.Error())
	}

	sse.Init() // 启动SSE推送

	RunServer(ctx, signalChannel, s)

	signal.Notify(signalChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	fmt.Println("收到关闭服务信号:", <-signalChannel)
	time.Sleep(time.Second * 3)
	fmt.Println("成功关闭服务器")
	return nil
}

// NewMainCommand creates a new gcmd.Command with a custom server getter function
func NewMainCommand(s *ghttp.Server) gcmd.Command {
	return gcmd.Command{
		Name:  "main",
		Usage: "main [options]",
		Brief: "start sagoo-iot server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return runMain(ctx, s)
		},
	}
}
