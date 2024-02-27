package sysenv

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"os"
	"runtime"
	"sagooiot/pkg/utility/utils"
)

var (
	SysStartTime *gtime.Time
	LocalIP      string //本地IP
	PublicIP     string //公网IP
	GoDiskSize   string //Go程序占磁盘空间
)

// GetGoRunInfo 获取运行信息
func GetGoRunInfo() (data []byte) {
	var (
		tmpData      = g.Map{}
		SysRunDir, _ = os.Getwd()
		gm           runtime.MemStats
	)

	runtime.ReadMemStats(&gm)
	// GO运行信息
	tmpData = g.Map{
		"goName":    "Golang",
		"goOs":      runtime.GOOS,                                       //操作系统
		"arch":      runtime.GOARCH,                                     //系统架构
		"goVersion": runtime.Version(),                                  //GO版本
		"startTime": SysStartTime,                                       //系统开始时间
		"runTime":   gtime.Now().Timestamp() - SysStartTime.Timestamp(), //运行时长
		"rootPath":  runtime.GOROOT(),
		"pwd":       SysRunDir,
		"goroutine": runtime.NumGoroutine(),
		"goMem":     utils.FileSize(int64(gm.Sys)), //运行内存
		"goSize":    GoDiskSize,                    //磁盘占用
	}

	data, _ = json.Marshal(tmpData)
	return
}
