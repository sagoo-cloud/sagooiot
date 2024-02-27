package sse

import (
	"github.com/gogf/gf/v2/net/ghttp"
	sseserver "github.com/xinjiayu/sse"
	"sagooiot/internal/sse/sysenv"
	"time"
)

var sysEnvEvent = sseserver.NewServer()

func runSysEnv() {
	hostMsg := sseserver.SSEMessage{}
	hostMsg.Event = "host"
	hostMsg.Data = sysenv.GetHostInfo()
	hostMsg.Namespace = "/sysenv/host"
	sysEnvEvent.Broadcast <- hostMsg

	sysLoadMsg := sseserver.SSEMessage{}
	sysLoadMsg.Event = "sysLoad"
	sysLoadMsg.Data = sysenv.GetSysLoad()
	sysLoadMsg.Namespace = "/sysenv/sysLoad"
	sysEnvEvent.Broadcast <- sysLoadMsg

	cpuMsg := sseserver.SSEMessage{}
	cpuMsg.Event = "cpu"
	cpuMsg.Data = sysenv.GetCpuInfo()
	cpuMsg.Namespace = "/sysenv/cpu"
	sysEnvEvent.Broadcast <- cpuMsg

	memMsg := sseserver.SSEMessage{}
	memMsg.Event = "mem"
	memMsg.Data = sysenv.GetMemInfo()
	memMsg.Namespace = "/sysenv/mem"
	sysEnvEvent.Broadcast <- memMsg

	diskMsg := sseserver.SSEMessage{}
	diskMsg.Event = "disk"
	diskMsg.Data = sysenv.GetDiskInfo()
	diskMsg.Namespace = "/sysenv/disk"
	sysEnvEvent.Broadcast <- diskMsg

	netMsg := sseserver.SSEMessage{}
	netMsg.Event = "net"
	netMsg.Data = sysenv.GetNetStatusInfo()
	netMsg.Namespace = "/sysenv/net"
	sysEnvEvent.Broadcast <- netMsg

	goMsg := sseserver.SSEMessage{}
	goMsg.Event = "go"
	goMsg.Data = sysenv.GetGoRunInfo()
	goMsg.Namespace = "/sysenv/go"
	sysEnvEvent.Broadcast <- goMsg

	time.Sleep(time.Duration(1) * time.Second)
}

func SysenvMessageEvent(r *ghttp.Request) {
	sysEnvEvent.ServeHTTP(r.Response.RawWriter(), r.Request)
}
