// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthProvider is the golang structure of table oauth_provider for DAO operations like Where/Data.
type OauthProvider struct {
	g.Meta    `orm:"table:oauth_provider, do:true"`
	Name      interface{} // 授权对象 qq 或 wechat
	Logo      interface{} // 授权图标地址
	Appid     interface{} // appid
	Appsecret interface{} // appsecret
	Status    interface{} // 状态 0 关闭 1 开启
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
