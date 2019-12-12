package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)
var store = base64Captcha.DefaultMemStore

func CreateCaptcha() (string, string, error) {
	dDigit := base64Captcha.DriverDigit{80, 240, 4, 0.7, 5}
	c := base64Captcha.NewCaptcha(&dDigit, store)
	return c.Generate()
}

func VerifyCaptcha(idkey, verifyValue string) bool {
	fmt.Println("## captcha: ", verifyValue)
	if store.Verify(idkey, verifyValue, false) {
		return true
	} else {
		return false
	}
}