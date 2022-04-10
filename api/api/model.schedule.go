package api

import (
	// "log"
	_ "github.com/lib/pq"
)

type Schedule struct {
	Title string `form:"Title" json:"Title" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag"`
}

var schedule Schedule