package api

// import (
//     "github.com/gin-gonic/gin"
// )

// func InitializeRoutes(router *gin.Engine) {
// 	// Sign Up
// 	router.POST("/signup", SignUp)

// 	// Login
// 	router.POST("/login", LoginHandler)
// 	// 認証済のみアクセス可能なグループ
// 	authUserGroup := router.Group("/auth")
// 	authUserGroup.Use(LoginCheckMiddleware())
// 	{
// 		// authUserGroup.GET("/schedule",)
// 	}
// }