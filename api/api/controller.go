package api

import (
	"net/http"
	"encoding/json"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
)

func Login(c *gin.Context) {
	var request EmailLoginRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err := GetUserByEmail(request.Email)
		if err != nil {
			c.Status(http.StatusBadRequest)
		} else {
			session := sessions.Default(c)
			loginUser, err := json.Marshal(&User)
			if err != nil {
				c.Status(http.StatusInternalServerError)
			} else {
				session.Set("loginUser", string(loginUser))
				session.Save()
				c.Status(http.StatusOK)
			}
		}
	}
}