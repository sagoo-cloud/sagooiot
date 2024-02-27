package system

import (
	"context"
	"sagooiot/internal/service"

	systemV1 "sagooiot/api/v1/system"
)

// 图形验证码
var Captcha = cCaptcha{}

type cCaptcha struct{}

func (a *cCaptcha) Index(ctx context.Context, req *systemV1.CaptchaIndexReq) (res *systemV1.CaptchaIndexRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &systemV1.CaptchaIndexRes{
		Key: idKeyC,
		Img: base64stringC,
	}

	return
}
