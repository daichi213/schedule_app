package api

import (
	"net/http"
	"encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
	"github.com/koron/go-dproxy"
)

func SettingCookie(router *gin.Engine) {
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
}

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		// TODO dproxyを使用せずにsession.Getで値が取れないか確認する
		// TODO sessions.Default().Get()の返却型を確認する
		loginUser, err := dproxy.New(session.Get("loginUser")).String()
		if err != nil {
			c.Status(http.StatusUnauthorized)
			// TODO 返却値を確認して挙動を把握する
			c.Abort()
		} else {
			var loginInfo AuthUser
			err := json.Unmarshal([]byte(loginUser), &loginInfo)
			if err != nil {
				c.Status(http.StatusUnauthorized)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}