package api

import (
	// "os"
	"time"
	"log"
    "github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	// "github.com/joho/godotenv"

	jwt "github.com/appleboy/gin-jwt/v2"
)

// 廃止
// func PasswordToHash(password string) ([]byte, error) {
// 	err := godotenv.Load(CurrentDir + "/salt.env")
// 	if err != nil {
// 		log.Fatalf("An error occurred while loading salt")
// 		return []byte(""), err
// 	}
// 	salt := os.Getenv("SALT")
// 	passSalt := password + salt
// 	hashed, err := bcrypt.GenerateFromPassword([...]byte(passSalt), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Fatalf("An error occurred while hashing password")
// 		return []byte(""), err
// 	}
// 	return hashed, err
// }

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
			sentEmail := loginVals.Email
			sentPassword, err := bcrypt.GenerateFromPassword([]byte(loginVals.Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatalf("An error occurred while password is hashing")
				return "", jwt.ErrMissingLoginValues
			}

			// TODOHTTPヘッダからIPアドレスを記録できるようにする
			if err := GetUserByEmail(sentEmail); err != nil {
				log.Fatalf("No existing password is sent")
				return "", jwt.ErrMissingLoginValues
			}

			if invalid := bcrypt.CompareHashAndPassword(sentPassword, User.Password); invalid != nil {
				return nil, jwt.ErrFailedAuthentication
			} else {
				return &Login{
					UserName: 	User.UserName,
					Email: 		User.Email,
					Password:	loginVals.Password,
				}, nil
			}

			// if (sentEmail == User.Email && sentPassword == User.Password) {
			// 	return &Login{
			// 		UserName: 	User.UserName,
			// 		Email: 		User.Email,
			// 		Password:	loginVals.Password,
			// 	}, nil
			// }
			// return nil, jwt.ErrFailedAuthentication
		},
		// 認可(権限の確認)
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*Login); ok && v.AdminFlag == 1 {
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
