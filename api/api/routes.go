package api

import (
	"fmt"
	"log"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	fmt.Println("checkpoint 1")
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
		// USER関連
		auth.PUT("/:user_id", UserEditHandler)
		auth.DELETE("/:user_id", UserEditHandler)
		// schedule関連
		auth.POST("/schedule", ScheduleCreateHandler)
		// auth.GET("/schedules", AllScheduleGetHandler)
		// auth.GET("/schedule/:schedule_id", ScheduleGetHandler)
		// auth.PUT("/schedule/:schedule_id", ScheduleUpdateHandler)
		// auth.DELETE("/schedule/:schedule_id", ScheduleDeleteHandler)
		// schedule関連に紐づくtodo関連
		// auth.POST("/todo", TodoCreateHandler)
		// auth.GET("/todos", AllTodoGetHandler)
		// auth.GET("/todos", TodoGetHandler)
		// auth.PUT("/todo", TodoUpdateHandler)
		// auth.DELETE("/todo", TodoDeleteHandler)
	}

	// AuthMiddleWareの初期化
	if errInit := authMiddleware.MiddlewareInit();errInit != nil {
		log.Fatal("AuthMiddleware.MiddlewareInit failed: ", errInit.Error())
	}
}