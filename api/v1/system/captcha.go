package system

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CaptchaIndexReq struct {
	g.Meta `path:"/captcha" method:"get" tags:"登录" summary:"获取默认的验证码"`
}
type CaptchaIndexRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key" dc:"key"`
	Img    string `json:"img" dc:"图片"`
}
