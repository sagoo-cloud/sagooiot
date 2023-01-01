package websocket

import "github.com/gogf/gf/v2/frame/g"

type ConfigureDiagramWebsocketReq struct {
	g.Meta `path:"/configureDiagram/ws" method:"get"  tags:"websocket管理" summary:"组态拓扑图发送消息"`
	Id     string `v:"required"`
}
type ConfigureDiagramWebsocketRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>" dc:"组态拓扑图"`
}

type MonitorSearchReq struct {
	g.Meta `path:"/monitorServer/ws" method:"get" tags:"websocket管理"  summary:"服务监控"`
}

type MonitorSearchRes struct {
	g.Meta `mime:"text/html" type:"string" example:"<html/>" dc:"服务监控"`
}
