package api

import (
	"log"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Call the authMiddleware
	authMiddleware, err := CallAuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// Sign Up
	router.POST("/signup", SignUp)

	// Login
	router.POST("/login", authMiddleware.LoginHandler)

	// 404のRouting
	router.NoRoute(authMiddleware.MiddlewareFunc(), NoRouting)

	// 認証後のRouting
	auth := router.Group("/auth")
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", HelloHandler)
		auth.GET("/schedule", ScheduleHandler)
	}

	// AuthMiddleWareの初期化
	if errInit := authMiddleware.MiddlewareInit();errInit != nil {
		log.Fatal("AuthMiddleware.MiddlewareInit failed: ", errInit.Error())
	}
}