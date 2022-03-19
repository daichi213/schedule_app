package api

// import (
// 	"net/http"
// 	"time"
// 	"encoding/json"
//     "github.com/gin-gonic/gin"
	
// 	jwt "github.com/appleboy/gin-jwt/v2"
// )

// type login struct {
// 	Username string `form:"username" json:"username" binding:"required"`
// 	Email string `form:"email" json:"email" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

// var identityKey = "id"

// func helloHandler(c *gin.Context) {
// 	claims, err := jwt.ExtractClaims(c)
// 	user, _ := c.Get(identityKey)
// 	c.JSON(200, gin.H{
// 		"userID": claims[identityKey],
// 		"userName": user.(*User).UserName,
// 		"text": "Hello World.",
// 	})
// }

// type User struct {
// 	UserName string
// 	FirstName string
// 	LastName string
// }

// func main() {
// 	port := os.Getenv("PORT")
// 	r := gin.New()
// 	r.Use(gin.Logger())
// 	r.Use(gin.Recovery())

// 	if port == "" {
// 		port  = "8000"
// 	}
// }

// // jwt middleware
// authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
// 	Realm:		"test zone",
// 	Key:  		[]byte("secret key"),
// 	Timeout:	time.Hour,
// 	MaxRefresh:	time.Hour,
// 	IdentityKey: identityKey,
// 	PayloadFunc: func(data interface{}) jwt.MapClaims {
// 		if v, ok := data.(*User); ok {
// 			return jwt.Claims{
// 				identityKey: u.UserName,
// 			}
// 		}
// 		return jwt.MapClaims{}
// 	},
// 	IdentityHandler: func(c *gin.Context) interface{} {
// 		claims := jwt.ExtractClaims(c)
// 		return &User{
// 			UserName: claims[identityKey].(string),
// 		}
// 	},
// 	// 認証(ユーザー本人かどうかの確認)
// 	Authenticator: func(c *gin.Context) (interface{}, error) {
// 		var loginVals login
// 		if err := c.ShouldBind(&loginVals); err != nil {
// 			return "", jwt.ErrMissingLoginValues
// 		}
// 		userID := loginVals.UserName
// 		password := loginVals.Password

// 		if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
// 			return &User{
// 				UserName: 	userID,
// 				LastName: 	"Bo-Yi",
// 				FirstName:	"Wu",
// 			}, nil
// 		}

// 		return nil, jwt.ErrFailedAuthentication
// 	},
// 	// 認可(権限の確認)
// 	Authorization: func(data interface{}, c *gin.Context) bool {
// 		if v, ok := data.(*User); ok && v.UserName == "admin" {
// 			return true
// 		}
// 		return false
// 	},
// 	Unauthorized: func(c *gin.Context, code int, message string) {
// 		c.JSON(code, gin.H{
// 			"code":   code,
// 			"message":   message,
// 		})
// 	},
// 	TokenLookup: "header: Authorization, query: token, cookie: jwt",

// 	TokenHeadName: "Bearer",

// 	TimeFunc: time.Now(),
// })

// if err != nil {
// 	log.Fatal("JWT Error:" + err.Error())
// }

// errInit := authMiddleware.MiddlewareInit()

// if errInit != nil {
// 	log.Fatal("authMiddleware.MiddlewareInit failed: ", errInit.Error())
// }

// //TODO jwt.New()に含まれている？
// r.POST("/login",authMiddleware.LoginHandler)

// r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
// 	claims := jwt.ExtractClaims(c)
// 	log.Printf("NoRoute claims: %#v\n", claims)
// 	c.JSON(404, gin.H{"code":"PAGE_NOT_FOUND", "message": "Page not found"})
// })

// auth := r.Group("/auth")

// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
// auth.Use(authMiddleware.MiddlewareFunc())
// {
// 	auth.GET("/hello", helloHandler)
// }

// if err := http.ListenAndServe(":" + port, r); err != nil {
// 	log.Fatal(err)
// }