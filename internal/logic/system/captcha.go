package system

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/sagoo-cloud/sagooiot/internal/service"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

type sCaptcha struct{}

var (
	captchaStore  = base64Captcha.DefaultMemStore
	captchaDriver = &base64Captcha.DriverString{
		Height:          80,
		Width:           240,
		NoiseCount:      50,
		ShowLineOptions: 20,
		Length:          4,
		Source:          "abcdefghjkmnpqrstuvwxyz23456789",
		Fonts:           []string{"chromohv.ttf"},
		BgColor:         &color.RGBA{R: 209, G: 205, B: 205, A: 90},
	}
)

func init() {
	service.RegisterCaptcha(New())
}

// Captcha 验证码管理服务
func New() *sCaptcha {
	return &sCaptcha{}
}

// GetVerifyImgString 获取字母数字混合验证码
func (s *sCaptcha) GetVerifyImgString(ctx context.Context) (idKeyC string, base64stringC string, err error) {
	driver := captchaDriver.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, captchaStore)
	idKeyC, base64stringC, err = c.Generate()
	return
}

// VerifyString 验证输入的验证码是否正确
func (s *sCaptcha) VerifyString(id, answer string) bool {
	c := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	answer = gstr.ToLower(answer)
	return c.Verify(id, answer, true)
}
