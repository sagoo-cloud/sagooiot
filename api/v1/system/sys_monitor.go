package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ListLogsReq struct {
	g.Meta `path:"/monitor/listLogs" tags:"系统日志" method:"get" summary:"日志列表"`
	Types  string `json:"types"  description:"日志类型" v:"required#types不能为空"`
}
type ListLogsRes struct {
	List []LogInfo `json:"list" `
}

type LogInfo struct {
	Name     string `json:"name" description:"日志文件名" `
	Size     string `json:"size" description:"日志大小" `
	ChangeAt string `json:"changeAt" description:"修改时间" `
}

type LastLinesLogReq struct {
	g.Meta `path:"/monitor/lastLinesLog" tags:"系统日志" method:"get" summary:"查看日志内容"`
	Types  string `json:"types"  description:"日志类型" v:"required#types不能为空"`
	Name   string `json:"name"  description:"日志文件名"  v:"required#name不能为空"`
}

type LastLinesLogRes struct {
	List []string `json:"list"`
}

// DeleteLogReq 删除日志
type DeleteLogReq struct {
	g.Meta `path:"/monitor/lastLinesLog/delete" method:"delete" tags:"系统日志" summary:"删除日志"`
	Types  string `json:"types"  description:"日志类型" v:"required#types不能为空"`
	Name   string `json:"name" description:"日志文件名" v:"required#name不能为空"`
}
type DeleteLogRes struct{}

type DownloadSysLogReq struct {
	g.Meta `path:"/monitor/downloadLog"       tags:"系统日志" method:"get" summary:"系统日志下载"`
	Name   string `json:"name" description:"日志文件名" v:"required#name不能为空"`
	Types  string `json:"types"  description:"日志类型" v:"required#types不能为空"`
}
type DownloadSysLogRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
