// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthProvider is the golang structure for table oauth_provider.
type OauthProvider struct {
	Name      string      `json:"name"      orm:"name"       ` // 授权对象 qq 或 wechat
	Logo      string      `json:"logo"      orm:"logo"       ` // 授权图标地址
	Appid     string      `json:"appid"     orm:"appid"      ` // appid
	Appsecret string      `json:"appsecret" orm:"appsecret"  ` // appsecret
	Status    int         `json:"status"    orm:"status"     ` // 状态 0 关闭 1 开启
	CallbackUrl    string  `json:"callbackUrl"    orm:"callback_url"`   // 回调url
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
