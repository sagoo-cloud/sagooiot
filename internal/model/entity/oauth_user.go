// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthUser is the golang structure for table oauth_user.
type OauthUser struct {
	Id        int         `json:"id"        orm:"id"         ` // 自增id
	Nickname  string      `json:"nickname"  orm:"nickname"   ` // 用户昵称
	AvatarUrl string      `json:"avatarUrl" orm:"avatar_url" ` // 用户授权头像
	Openid    string      `json:"openid"    orm:"openid"     ` // 用户授权唯一标识
	Provider  string      `json:"provider"  orm:"provider"   ` // 授权对象名 qq 或 wechat
	UserId    int         `json:"userId"    orm:"user_id"    ` // 用户系统身份 id
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` //
}
