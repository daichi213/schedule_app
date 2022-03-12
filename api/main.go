package main

import (
    "github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"api"
	_ "github.com/lib/pq"
)

var router *gin.Engine
var db *gorm.DB
var err error

func main() {
	router = gin.Default(

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Session("mysession", store)))

	// initializeRoutes()
	router.POST("/login", api.Login)
	// 認証済のみアクセス可能なグループ
	authUserGroup := router.Group("/auth")
	authUserGroup.Use(api.LoginCheckMiddleware()){
		authUserGroup.GET("/schedule",)
	}

	router.Run()
}