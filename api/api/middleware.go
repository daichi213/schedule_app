package api

import (
	"time"
    "github.com/gin-gonic/gin"
	"crypto/bcrypt"

	jwt "github.com/appleboy/gin-jwt/v2"
)

// パスワードのハッシュ化関数
// TODO saltをenvファイルに生成しているので、それを読み込みpasswardに連結してハッシュ化する
// TODO dockerのbuild時にsaltをenvファイルに出力するようにする
// cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 36 | head -n 1 | sort | uniq > salt.env
func PasswordToHash(password string) []byte {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error generating password")
	}
	return hashed
}

// jwt middleware
var IdentityKey = "id"

func CallAuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	AuthMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:		"test zone",
		Key:  		[]byte("secret key"),
		Timeout:	time.Hour,
		MaxRefresh:	time.Hour,
		IdentityKey: IdentityKey,
		// login後に呼び出される関数
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*Login); ok {
				return jwt.MapClaims{
					IdentityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		// Authorizatorへ値を渡すための関数
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &Login{
				UserName: claims[IdentityKey].(string),
			}
		},
		// 認証(ユーザー本人かどうかの確認)
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.UserName
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &Login{
					UserName: 	userID,
					Email: 		"Bo-Yi",
					Password:	"Wu",
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		// 認可(権限の確認)
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*Login); ok && v.UserName == "admin" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":		code,
				"message":	message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	return AuthMiddleware, err
}
