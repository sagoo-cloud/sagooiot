package sse

import (
	"github.com/gogf/gf/v2/os/gtime"
	"sagooiot/internal/sse/sysenv"
)

func Init() {
	sysenv.SysStartTime = gtime.Now() //初始化开始时间

	for _, f := range []func(){runRedisInfo, runSysEnv} {
		go func(f func()) {
			for {
				f()
			}
		}(f)
	}
}
