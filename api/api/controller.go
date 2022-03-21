package api

import (
	"log"
	"net/http"
    "github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
)

func SignUp(c *gin.Context) {
	var signupUser Login
	err := c.BindJSON(&signupUser)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err := CreateUser(&signupUser)
		if err != nil {
			// TODO エラー発生時にルートパスへリダイレクトさせる処理を追加する
			c.Status(http.StatusBadRequest)
		} else {
			c.JSON(200, gin.H{
				"UserName": signupUser.UserName,
				"Email": signupUser.Email,
			})
			// c.Redirect(http.StatusFound, "/auth/schedule")
		}
	}
}

func ScheduleHandler(c *gin.Context) {
	// claims := jwt.ExtractClaims(c)
	// user, _ := c.Get(IdentityKey)
	// TODO ログインユーザーにひもづくschedule, todoのレコードを取得する関数を呼び出す
}

func NoRouting(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	log.Printf("NoRoute claims: %#v\n", claims)
	c.JSON(404, gin.H{"code":"PAGE_NOT_FOUND", "message": "Page not found"})
}

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(IdentityKey)
	c.JSON(200, gin.H{
		"userID": claims[IdentityKey],
		"userName": user.(*Login).UserName,
		"text": "Hello World.",
	})
}
