package system

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sagoo-cloud/sagooiot/internal/model"
)

type LoginDoReq struct {
	g.Meta    `path:"/login" method:"post" summary:"执行登录请求" tags:"登录"`
	UserName  string `json:"userName" v:"required#请输入账号"   dc:"账号"`
	Password  string `json:"password" v:"required#请输入密码"   dc:"密码"`
	Captcha   string `json:"captcha"  v:"required#请输入验证码" dc:"验证码"`
	VerifyKey string `json:"verifyKey"`
}
type LoginDoRes struct {
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type LoginOutReq struct {
	g.Meta `path:"/loginOut" method:"post" summary:"退出" tags:"登录"`
}
type LoginOutRes struct {
}
