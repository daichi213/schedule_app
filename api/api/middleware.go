package api

import (
	"fmt"
	"time"
	"log"
    "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/appleboy/gin-jwt/v2"
)

// TODO MODELの変数をmiddlewareでも参照してしまっているため、分離できるように努める（DIコンテナで実現できるか？）
// jwt middleware
var IdentityKey = "email"

func CallAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	fmt.Println("checkpoint 2")
	AuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:		"test zone",
		Key:  		[]byte("secret key"),
		Timeout:	time.Hour,
		MaxRefresh:	time.Hour,
		IdentityKey: IdentityKey,
		// login後に呼び出される関数
		PayloadFunc: internalPayloadFunc,
		// Authorizatorへ値を渡すための関数
		IdentityHandler: internalIdentityHandlerFunction,
		// 認証(ユーザー本人かどうかの確認)
		Authenticator: internalAuthenticatorFunction,
		// 認可(権限の確認)
		// token発行後のページの読み込み制御についての関数
		Authorizator: internalAuthorizatorFunction,
		Unauthorized: internalUnauthorizedFunction,
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	return AuthMiddleware, err
}

func internalPayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*Login); ok {
		return jwt.MapClaims{
			"user_id": v.ID,
			IdentityKey: v.Email,
		}
	}
	return jwt.MapClaims{}
}

func internalIdentityHandlerFunction(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &Login{
		Email: claims[IdentityKey].(string),
	}
}

func internalAuthenticatorFunction(c *gin.Context) (interface{}, error) {
	var loginVals Login
	if err := c.ShouldBind(&loginVals); err != nil {
		fmt.Println("checkpoint shouldbind")
		return "", jwt.ErrMissingLoginValues
	}

	// TODOHTTPヘッダからIPアドレスを記録できるようにする
	if err := GetUserByEmail(loginVals.Email); err != nil {
		// log.Fatalf("No existing password is sent")
		fmt.Println("checkpoint get user")
		return "", jwt.ErrMissingLoginValues
	}

	if invalid := bcrypt.CompareHashAndPassword(UserFromDB.Password, []byte(loginVals.Password)); invalid != nil {
		log.Fatalf("Password is wrong...;dbside:%v,loginVals:%v", UserFromDB.Password, []byte(loginVals.Password))
		return "", jwt.ErrFailedAuthentication
	} else {
		return &Login{
			UserName: 	UserToDB.UserName,
			Email: 		UserToDB.Email,
			Password:	loginVals.Password,
		}, nil
	}
}

func internalAuthorizatorFunction(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*Login); ok {
		return true
	}
	return false
}

func internalUnauthorizedFunction(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":		code,
		"message":	message,
	})
}