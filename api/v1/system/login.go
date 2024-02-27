package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"sagooiot/internal/model"
)

type LoginDoReq struct {
	g.Meta    `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	UserName  string `json:"userName" v:"required#请输入账号"   dc:"账号"`
	Password  string `json:"password" v:"required#请输入密码"   dc:"密码"`
	Captcha   string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
	VerifyKey string `json:"verifyKey"`
}
type LoginDoRes struct {
	UserInfo    *model.LoginUserRes `json:"userInfo"`
	Token       string              `json:"token"`
	IsChangePwd int                 `json:"isChangePwd" dc:"是否需要变更密码， 1是 0 否"`
}

type LoginOutReq struct {
	g.Meta `path:"/loginOut" method:"post" summary:"退出" tags:"登录"`
}
type LoginOutRes struct {
}

type EditPasswordReq struct {
	g.Meta          `path:"/user/editPassword" method:"post" summary:"用户修改密码" tags:"登录"`
	UserName        string `json:"userName"        description:"用户账户" v:"required#账户不能为空"`
	OldUserPassword string `json:"oldUserPassword"        description:"原密码" v:"required#原密码不能为空"`
	UserPassword    string `json:"userPassword"  description:"登录密码;cmf_password加密" v:"required#新密码不能为空"`
}
type EditPasswordRes struct {
}
