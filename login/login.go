package login

import (
	"github.com/kataras/iris/v12"
	"iland/auth"
	"iland/common"
)

type LoginReq struct {
	Name string `json:"name"`
	Password string `json:"password"`
	CaptchaCode int `json:"captcha_code"`
}

type LoginRes struct {
	common.Response
	Token string `json:"token"`
}

func Login(ctx iris.Context){
	req := LoginReq{}
	res := LoginRes{}

	if err := ctx.ReadJSON(&req); err != nil {
		res.Code = 1
		res.Message = err.Error()
		_, _ = ctx.JSON(res)
		return
	}

	// TODO 检查密码

	// TODO 生成token
	id := ""
	token, err := auth.GenerateToken(id)
	if err != nil {
		res.Code = 1
		res.Message = err.Error()
		_, _ = ctx.JSON(res)
	}

	res.Code = 0
	res.Message = "ok"
	res.Token = token
	_, _ = ctx.JSON(res)
}