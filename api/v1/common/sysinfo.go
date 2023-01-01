package common

import "github.com/gogf/gf/v2/frame/g"

type GetSysInfoReq struct {
	g.Meta `path:"/sysinfo" tags:"公共方法" method:"get" summary:"获取系统相关的信息"`
}

type GetSysInfoRes g.Map

type IsTokenReq struct {
	g.Meta `path:"/isToken" method:"get" summary:"验证token是否正确" tags:"公共方法"`
}
type IsTokenRes struct {
	ExpiresAt int64 `json:"expiresAt"`
	IsToken   bool  `json:"isToken"`
}
