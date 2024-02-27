package common

import "github.com/gogf/gf/v2/frame/g"

type IsTokenReq struct {
	g.Meta `path:"/isToken" method:"get" summary:"验证token是否正确" tags:"公共方法"`
}
type IsTokenRes struct {
	ExpiresAt int64  `json:"expiresAt"`
	IsToken   bool   `json:"isToken"`
	Auth      string `json:"auth"`
}

type CheckAccessAuthReq struct {
	g.Meta  `path:"/checkAccessAuth" method:"get" summary:"验证接口是否具有访问权限" tags:"公共方法"`
	Address string `p:"address"   description:"接口地址" v:"required#接口地址不能为空"`
}
type CheckAccessAuthRes struct {
	IsAllow bool `json:"isAllow"`
}
