package main

import (
    "github.com/gin-gonic/gin"
	"gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

var router *gin.Engine
var db *gorm.DB
var err error

func main() {
	router = gin.Default()

	initializeRoutes()

	router.Run()
}