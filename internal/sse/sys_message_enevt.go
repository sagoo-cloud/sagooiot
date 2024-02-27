package sse

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/xinjiayu/sse"
	"sagooiot/internal/sse/sysenv"
	"time"
)

func SysMessageEntvt(r *ghttp.Request) {
	sseServer := sseserver.NewServer()
	userId := gconv.Int(r.GetQuery("userId"))
	isCLose := make(chan bool)
	go func() {
		for {
			select {
			case <-isCLose:
				return
			default:
				data := sysenv.GetUnReadMessageLast(userId)

				if data != nil && len(data) > 0 {
					sysMessage := sseserver.SSEMessage{}
					sysMessage.Event = "lastUnRead"
					sysMessage.Data = data
					sysMessage.Namespace = "/sysMessage/lastUnRead"
					sseServer.Broadcast <- sysMessage
				}
				time.Sleep(time.Duration(5) * time.Second)
			}
		}
	}()
	sseServer.ServeHTTP(r.Response.RawWriter(), r.Request)
	isCLose <- true
}
