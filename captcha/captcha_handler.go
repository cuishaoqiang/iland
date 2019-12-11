package captcha

import (
	"github.com/kataras/iris/v12"
	"iland/common"
)

type GetCaptchaRes struct {
	common.Response
	IDKey string `json:"idkey"`
	Captcha string `json:"captcha"`
}

type CheckCaptchaReq struct {
	IDKey string `json:"idkey"`
	Verify string `json:"verify"`
}

type CheckCaptchaRes struct {
	common.Response
	CaptchaCode string `json:"captcha_code"`
}

func Get(ctx iris.Context){
	res := GetCaptchaRes{}

	idkey, captcha, err := CreateCaptcha()
	if err != nil {
		res.Code = 1
		res.Message = err.Error()
		return
	}
	res.Code = 0
	res.Message = "ok"
	res.IDKey = idkey
	res.Captcha = captcha
	_, _ = ctx.JSON(res)
}

func Check(ctx iris.Context) {
	req := CheckCaptchaReq{}
	res := CheckCaptchaRes{}

	if err := ctx.ReadJSON(&req); err != nil {
		res.Code = 1
		res.Message = err.Error()
		_, _ = ctx.JSON(res)
		return
	}

	if VerifyCaptcha(req.IDKey, req.Verify) {
		res.Code = 1
		res.Message = "验证码错误"
		_, _ = ctx.JSON(res)
		return
	}
	res.Code = 0
	res.Message = "ok"
	res.CaptchaCode = "1"
	_, _ = ctx.JSON(res)
}