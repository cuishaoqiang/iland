package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"iland/auth"
	"iland/captcha"
	"iland/log"
	"iland/login"
	"io/ioutil"
)

type ListenAddr struct {
	IP string `json:"ip"`
	Port string `json:"port"`
}

func main() {
	app := iris.New()

	log.OpenLogsPath()
	app.Logger().SetLevel("debug")

	customLogger, close2 := log.NewRequestLogger()
	defer close2()
	app.Use(customLogger)

	// TODO vue.js template 解析是否正常
	view := iris.HTML("web", ".html").Delims("{<", ">}")
	app.RegisterView(view)

	app.Get("/", func(ctx iris.Context) {
		_ = ctx.View("index_test.html")
	})

	lCaptcha := app.Party("/captcha")
	{
		lCaptcha.Get("/", captcha.Get)
		lCaptcha.Post("/", captcha.Check)
	}

	app.Post("/login", login.Login)

	app.Post("/logout", func(c context.Context) {

	})

	lUser := app.Party("/user", auth.CheckToken)
	{
		lUser.Post("/register", func(context context.Context) {})
		lUser.Post("/list", func(context context.Context) {})
	}

	listenAddr, err := getListenAddr()
	if err != nil {
		fmt.Println("## get listen address failed")
	}
	_ = app.Run(iris.Addr(fmt.Sprintf("%s:%s", listenAddr.IP, listenAddr.Port)))

}

func getListenAddr() (ListenAddr, error) {
	config, err := ioutil.ReadFile("./config/listenaddr.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	addr := ListenAddr{}
	jsonErr := json.Unmarshal([]byte(config), &addr)
	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
	}

	return addr, nil
}
