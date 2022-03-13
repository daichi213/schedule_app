package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
	"api/api"
)

// TODO どこかでGOROUTINEを使用した並列処理を入れるようにする

var router *gin.Engine
var err error

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// initializeRoutes()
	router.POST("/login", api.Login)
	// 認証済のみアクセス可能なグループ
	authUserGroup := router.Group("/auth")
	authUserGroup.Use(api.LoginCheckMiddleware())
	{
		authUserGroup.GET("/schedule",)
	}

	router.Run()
}