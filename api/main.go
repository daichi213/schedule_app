package main

import (
	"os"
	"io"
    "github.com/gin-gonic/gin"
	"api/api"
)

// TODO どこかでGOROUTINEを使用した並列処理を入れるようにする

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(os.Stdout, f)

	router := gin.Default()
	// TODO FrontEndのIPまたはDNSを指定する
	// router.SetTrustedProxies([]string{"192.168.1.2"})
	router.SetTrustedProxies([]string{"0.0.0.0"})

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api.InitializeRoutes(router)

	router.Run()
}
