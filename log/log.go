package log

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"os"
	"strings"
	"time"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func OpenLogsPath() {
	var logPath string = "./logs"
	exist, _ := PathExists(logPath)
	if !exist {
		dirErr := os.Mkdir(logPath, os.ModePerm)
		if dirErr != nil {
			fmt.Printf("mkdir failed![%v]\n", dirErr)
		}
	}
}

//根据日期获取文件名，文件日志以最常用的方式工作
//但这些只是好的命名方式。
func todayFilename() string {
	today := time.Now().Format("2006-01-02")
	return "./logs/" + today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	//打开一个输出文件，如果重新启动服务器，它将追加到今天的文件中
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

func NewRequestLogger() (h iris.Handler, close func() error) {
	close = func() error { return nil }
	c := logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// Columns: true,
		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	}
	logFile := newLogFile()
	close = func() error {
		err := logFile.Close()
		return err
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := logger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		_, _ = logFile.Write([]byte(output))
	}
	//我们不想使用记录器，一些静态请求等
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	h = logger.New(c)
	return
}
