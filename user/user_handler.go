package user

import (
	"github.com/kataras/iris/v12"
	_ "github.com/mattn/go-sqlite3"
	"iland/common"
)

type AddReq struct {
	Name      string    `json:"name"`
	Password  string    `json:"Password"`
}

type AddRes struct {
	common.Response
	UserId string `json:"user_id"`
}

func Add(ctx iris.Context){
	req := AddReq{}
	res := AddRes{}
	if err := ctx.ReadJSON(req); err != nil {
		res.Code = 1
		res.Message = err.Error()
		return
	}
	user := &UserOrm{
		Name: req.Name,
		Password: req.Password,
	}
	id, err2 := Insert(user)
	if err2 != nil {
		res.Code = 1
		res.Message = err2.Error()
		return
	}
	res.Code = 0
	res.Message = "ok"
	res.UserId = string(id)
	return
}

