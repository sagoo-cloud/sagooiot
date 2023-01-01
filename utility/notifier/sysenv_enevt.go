package notifier

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/sagoo-cloud/sagooiot/utility/notifier/sysenv"
	"github.com/xinjiayu/sse"
	"time"
)

func SysenvMessageEvent(r *ghttp.Request) {

	sseServer := sseserver.NewServer()
	go func() {
		for {
			hostMsg := sseserver.SSEMessage{}
			hostMsg.Event = "host"
			hostMsg.Data = sysenv.GetHostInfo()
			hostMsg.Namespace = "/sysenv/host"
			sseServer.Broadcast <- hostMsg

			sysLoadMsg := sseserver.SSEMessage{}
			sysLoadMsg.Event = "sysLoad"
			sysLoadMsg.Data = sysenv.GetSysLoad()
			sysLoadMsg.Namespace = "/sysenv/sysLoad"
			sseServer.Broadcast <- sysLoadMsg

			cpuMsg := sseserver.SSEMessage{}
			cpuMsg.Event = "cpu"
			cpuMsg.Data = sysenv.GetCpuInfo()
			cpuMsg.Namespace = "/sysenv/cpu"
			sseServer.Broadcast <- cpuMsg

			memMsg := sseserver.SSEMessage{}
			memMsg.Event = "mem"
			memMsg.Data = sysenv.GetMemInfo()
			memMsg.Namespace = "/sysenv/mem"
			sseServer.Broadcast <- memMsg

			diskMsg := sseserver.SSEMessage{}
			diskMsg.Event = "disk"
			diskMsg.Data = sysenv.GetDiskInfo()
			diskMsg.Namespace = "/sysenv/disk"
			sseServer.Broadcast <- diskMsg

			netMsg := sseserver.SSEMessage{}
			netMsg.Event = "net"
			netMsg.Data = sysenv.GetNetStatusInfo()
			netMsg.Namespace = "/sysenv/net"
			sseServer.Broadcast <- netMsg

			time.Sleep(time.Duration(2) * time.Second)
		}
	}()
	sseServer.ServeHTTP(r.Response.RawWriter(), r.Request)

}
