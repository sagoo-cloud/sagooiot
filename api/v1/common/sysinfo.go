package common

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetSysInfoReq struct {
	g.Meta `path:"/sysinfo" tags:"公共方法" method:"get" summary:"获取系统相关的信息"`
}

type GetSysInfoRes g.Map
