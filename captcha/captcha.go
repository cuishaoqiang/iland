package captcha

import (
	"github.com/mojocn/base64Captcha"
)
var store = base64Captcha.DefaultMemStore

func CreateCaptcha() (string, string, error) {
	var driver base64Captcha.DriverDigit
	c := base64Captcha.NewCaptcha(&driver, store)
	return c.Generate()
}

func VerifyCaptcha(idkey,verifyValue string) bool {
	if store.Verify(idkey, verifyValue, false) {
		return true
	} else {
		return false
	}
}