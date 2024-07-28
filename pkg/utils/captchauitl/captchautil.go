package captchauitl

import (
	"github.com/mojocn/base64Captcha"
)

type StringCaptcha struct {
	captcha *base64Captcha.Captcha
}

// NewStringCaptcha 创建验证码
func NewStringCaptcha() *StringCaptcha {
	store := base64Captcha.DefaultMemStore

	source := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	driver := base64Captcha.NewDriverString(
		38,
		120,
		0,
		1,
		4,
		source,
		nil,
		nil,
		nil,
	)

	captcha := base64Captcha.NewCaptcha(driver, store)
	return &StringCaptcha{
		captcha: captcha,
	}
}

// Generate 生成验证码
func (StringCaptcha *StringCaptcha) Generate() (string, string, string) {
	id, b64s, answer, _ := StringCaptcha.captcha.Generate()
	return id, b64s, answer
}

// Verify 校验验证码
func (StringCaptcha *StringCaptcha) Verify(id, answer string) bool {
	return StringCaptcha.captcha.Verify(id, answer, true)
}
