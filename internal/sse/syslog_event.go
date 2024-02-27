package sse

import (
	"context"
	"github.com/hpcloud/tail"
	"sagooiot/pkg/utility/utils"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	sseserver "github.com/xinjiayu/sse"
)

// LogInfoEvent 实时刷新日志
func LogInfoEvent(r *ghttp.Request) {

	// 获取日志文件名
	name := r.Get("name")
	// 获取日志类型
	types := r.Get("types")

	//指定文件名
	var url string
	switch types.String() {
	case "service":
		url = g.Cfg().MustGet(context.Background(), "server.logPath", "resource/log/server").String()
	case "run":
		url = g.Cfg().MustGet(context.Background(), "server.runLogPath", "var").String()
	case "sql":
		url = g.Cfg().MustGet(context.Background(), "database.logger.path", "resource/log/sql").String()
	default:
		url = g.Cfg().MustGet(context.Background(), "server.logPath").String()
	}
	filePathName := strings.Join([]string{url, "/", name.String()}, "")
	filename := filePathName

	//判断日志文件是否存在
	if !utils.FileIsExisted(filename) {
		return
	}

	//配置文件
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: true,                                 //文件不存在报错
		Poll:      true,
	}

	//以config的配置打开filename文件
	tails, err := tail.TailFile(filename, config)
	//处理错误
	if err != nil {
		g.Log().Error(r.Context(), err)
		return
	}

	var (
		line    *tail.Line
		ok      bool
		logChan = make(chan string)
	)
	//一行一行的读取日志
	go func() {
		for {
			line, ok = <-tails.Lines
			if !ok {
				g.Log().Errorf(r.Context(), "tail file close reopen, filename:%s\n", tails.Filename)
				continue
			}
			logChan <- line.Text
			time.Sleep(time.Second)
		}
	}()

	sseServer := sseserver.NewServer()
	go func() {
		for {
			hostMsg := sseserver.SSEMessage{}
			hostMsg.Event = "log"
			msg, ok := <-logChan
			if !ok {
				continue
			}
			hostMsg.Data = []byte(msg)
			hostMsg.Namespace = "/logInfo/log"
			sseServer.Broadcast <- hostMsg
		}
	}()
	sseServer.ServeHTTP(r.Response.RawWriter(), r.Request)

}
