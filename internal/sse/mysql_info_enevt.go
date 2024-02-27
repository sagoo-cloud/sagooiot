package sse

import (
	"github.com/gogf/gf/v2/net/ghttp"
	sseserver "github.com/xinjiayu/sse"
	"sagooiot/internal/sse/sysenv"
)

var mysqlInfoEvent = sseserver.NewServer()

func runMysqlInfo() {
	//所有mysql版本信息
	versionMsg := sseserver.SSEMessage{}
	versionMsg.Event = "version"
	versionMsg.Data = sysenv.GetMysqlVersionInfo()
	versionMsg.Namespace = "/mysqlinfo/version"
	mysqlInfoEvent.Broadcast <- versionMsg

	//所有mysql的status信息
	statusMsg := sseserver.SSEMessage{}
	statusMsg.Event = "status"
	statusMsg.Data = sysenv.GetMysqlStatusInfo()
	statusMsg.Namespace = "/mysqlinfo/status"
	mysqlInfoEvent.Broadcast <- statusMsg

	//所有mysql统计量信息
	variablesMsg := sseserver.SSEMessage{}
	variablesMsg.Event = "variables"
	variablesMsg.Data = sysenv.GetMysqlVariablesInfo()
	variablesMsg.Namespace = "/mysqlinfo/variables"
	mysqlInfoEvent.Broadcast <- variablesMsg
}

func MysqlInfoMessageEvent(r *ghttp.Request) {
	mysqlInfoEvent.ServeHTTP(r.Response.RawWriter(), r.Request)

}
