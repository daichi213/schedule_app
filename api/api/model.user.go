package api

import (
	"log"
	_ "github.com/lib/pq"
)

type Login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var User Login

func CreateUser(user *Login) error {
	db , err := GetDB()
	tx := db.Begin()
	if err != nil {
		log.Fatalf("An Error occurred while connecting to database: %v", err)
		panic(err)
	} else {
		DB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not find DB: %v", err)
			panic(err)
		}
		defer DB.Close()
		if err := tx.Model(user).Create(user).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Could not create: %s", err.Error())
		} else {
			tx.Commit()
		}
		return err
	}
}
