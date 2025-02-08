package oauth

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model/entity"
)

type AuthLoginReq struct {
	g.Meta   `path:"/login" method:"get" summary:"sso授权" tags:"授权登录"`
	Provider string `json:"provider" description:"授权对象 qq | wechat" v:"required#provider不能为空"`
}

type AuthLoginRes struct{}

type AuthCallbackReq struct {
	g.Meta   `path:"/callback" method:"get" summary:"sso授权登录回调" tags:"授权登录"`
	Code     string `json:"code" description:"授权凭证" v:"required#code不能为空"`
	Provider string `json:"provider" description:"授权对象 qq | wechat" v:"required#provider不能为空"`
	State    string `json:"state"`
}

type AuthCallbackRes struct {
	OauthUser *entity.OauthUser `json:"oauthUser"`
	UserInfo  *LoginUserRes     `json:"loginUser"`
}

type LoginUserRes struct {
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	Avatar       string `orm:"avatar" json:"avatar"`                 //头像
	Token        string `json:"token"`
}

// 获取 系统用户绑定授权信息
type GetUserBindingReq struct {
	g.Meta   `path:"/users/{user_id}" method:"get" summary:"获取系统账户授权绑定信息" tags:"授权登录"`
	Provider string `json:"provider" description:"授权对象 qq | wechat"`
	UserId   int    `json:"user_id" description:"系统用户身份id" v:"required#user_id不能为空"`
}

type GetUserBindingRes struct {
	g.Meta `mime:"application/json"`
	Users  []*entity.OauthUser `json:"users" dc:"授权用户列表"`
}

// 系统用户登录授权绑定
type UserBindingReq struct {
	g.Meta   `path:"/binding" method:"post" summary:"系统用户登录绑定授权用户" tags:"授权登录"`
	Provider string `json:"provider" description:"授权对象 qq | wechat" v:"required#provider不能为空"`
	Openid   string `json:"openid" description:"授权用户身份id" v:"required#openid不能为空"`
}
type UserBindingRes struct {
	g.Meta            `mime:"application/json"`
	*entity.OauthUser `json:"user" dc:"授权用户列表"`
}
