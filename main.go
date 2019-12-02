package main

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"io/ioutil"
)

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

	_ = app.Run(iris.Addr(listenAddr))

}

func getListenAddr() (string, string, error) {
	config, err := ioutil.ReadFile("./config/listen_addr.json")
	if err != nil {
		fmt.Println(err.Error())
	}

	addr := ListenAddr{}
	jsonErr := json.Unmarshal([]byte(config), &addr)
	if jsonErr != nil {
		fmt.Println(jsonErr.Error())
	}

	return addr.Ip, addr.Port, nil
}
