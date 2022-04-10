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
		log.Printf("BindJSON is failed: %v", err)
	} else {
		err := CreateUser(&signupUser)
		if err != nil {
			// TODO エラー発生時にルートパスへリダイレクトさせる処理を追加する
			c.Status(http.StatusBadRequest)
			log.Printf("CreateUser is failed: %v", err)
		} else {
			c.JSON(200, gin.H{
				"UserName": signupUser.UserName,
				"Email": signupUser.Email,
			})
			// c.Redirect(http.StatusFound, "/auth/schedule")
		}
	}
}
