package oauth

import (
	"github.com/gogf/gf/v2/frame/g"
)

// 获取 平台 sso 授权配置信息
type ListProviderReq struct {
	g.Meta `path:"/provider" method:"get" summary:"sso授权配置信息" tags:"授权登录"`
	//Name   string `json:"name"  description:"sso授权名称"`
	//Status int    `json:"status" description:"是否开启"`
}

type Provider struct {
	Name   string `json:"name"      description:"sso授权名称"`
	Logo   string `json:"logo"      description:"图标logo地址"`
	Status int    `json:"status"    description:"是否开启 0 关闭 1 启用"`
}

type ListProviderRes struct {
	g.Meta    `mime:"application/json"`
	Providers []*Provider `json:"providers" dc:"授权配置"`
}
