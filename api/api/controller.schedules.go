package api

import (
	"log"
	"net/http"
    "github.com/gin-gonic/gin"
)

// TODO 認可はjwt.ExtractClaimsを使用してuser_idを参照して行う

func ScheduleCreateHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(IdentityKey)
	// TODO ログインユーザーにひもづくschedule, todoのレコードを取得する関数を呼び出す
}

func AllScheduleGetHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(IdentityKey)
	// TODO ログインユーザーにひもづくschedule, todoのレコードを取得する関数を呼び出す
}

func ScheduleGetHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(IdentityKey)
	// TODO ログインユーザーにひもづくschedule, todoのレコードを取得する関数を呼び出す
}

func ScheduleUpdateHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(IdentityKey)
	// TODO ログインユーザーにひもづくschedule, todoのレコードを取得する関数を呼び出す
}

func ScheduleDeleteHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(IdentityKey)
	// TODO ログインユーザーにひもづくschedule, todoのレコードを取得する関数を呼び出す
}
