package main

import (
    "github.com/gin-gonic/gin"
	"api/api"
)

// TODO どこかでGOROUTINEを使用した並列処理を入れるようにする

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api.InitializeRoutes(router)

	router.Run()
}
