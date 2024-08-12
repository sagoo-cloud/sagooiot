// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthUser is the golang structure of table oauth_user for DAO operations like Where/Data.
type OauthUser struct {
	g.Meta    `orm:"table:oauth_user, do:true"`
	Id        interface{} // 自增id
	Nickname  interface{} // 用户昵称
	AvatarUrl interface{} // 用户授权头像
	Openid    interface{} // 用户授权唯一标识
	Provider  interface{} // 授权对象名 qq 或 wechat
	UserId    interface{} // 用户系统身份 id
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
