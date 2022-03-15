package api

import (
	"net/http"
	"encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)

func SessionsSet(c *gin.Context, request *AuthUser) {
	session := sessions.Default(c)
	loginUser, err := json.Marshal(&request)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		session.Set("loginUser", string(loginUser))
		session.Save()
		c.Status(http.StatusOK)
	}
}

func SignUp(c *gin.Context) {
	var signupUser AuthUser
	err := c.BindJSON(&signupUser)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err := CreateUser(&signupUser)
		if err != nil {
			// TODO エラー発生時にルートパスへリダイレクトさせる処理を追加する
			c.Status(http.StatusBadRequest)
		} else {
			SessionsSet(c, &signupUser)
		}
	}
}

func Login(c *gin.Context) {
	var request AuthUser
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err := GetUserByEmail(request.Email)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			SessionsSet(c, &request)
		}
	}
}