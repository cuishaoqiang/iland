package auth

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"time"
)

var JWT_KEY string = "JWT_KEY_SECRET"

func GenerateToken(id string) (string, error) {
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
       "id": id,
       "iss":"iland",
       "iat":time.Now().Unix(),
       "exp":time.Now().Add(24 * time.Hour * time.Duration(1)).Unix(), // 24h
   })
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	return tokenString, err
}

func CheckToken(ctx iris.Context) {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			//自己加密的秘钥或者说盐值
			return []byte(JWT_KEY), nil
		},
		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
		//重要的是要避免此处的安全问题：https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		//验证未通过错误处理方式
		//ErrorHandler: func(context.Context, string)

		//debug 模式
		//Debug: bool
	})

	jwtHandler.Serve(ctx)
}

