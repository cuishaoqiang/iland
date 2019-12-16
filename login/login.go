package login

import (
	"github.com/kataras/iris/v12"
	"iland/auth"
	"iland/common"
	"iland/encoder"
	"iland/user"
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

	// 检查验证码
	if req.CaptchaCode != 1 {
		res.Code = 1
		res.Message = "验证码无效"
		_, _ = ctx.JSON(res)
		return
	}

	// 检查密码
	userOrm := &user.UserOrm{
		Name:      req.Name,
	}

	ok, err := user.Get(userOrm)
	if !ok {
		res.Code = 1
		res.Message = err.Error()
		_, _ = ctx.JSON(res)
		return
	}

	plainPassword, errE := encoder.RSADecrypt([]byte(req.Password))
	if errE != nil {
		res.Code = 1
		res.Message = "解密错误"
		_, _ = ctx.JSON(res)
	}

	if string(plainPassword) != userOrm.Password {
		res.Code = 1
		res.Message = "密码错误"
		_, _ = ctx.JSON(res)
		return
	}

	// 生成token
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