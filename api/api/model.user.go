package api

import (
	"log"
	_ "github.com/lib/pq"
)

type Login struct {
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Email string `form:"Email" json:"Email" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag" binding:"required"`
}

type UserFromDB struct {
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Email string `form:"Email" json:"Email" binding:"required"`
	Password []byte `form:"Password" json:"Password" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag" binding:"required"`}

var User UserFromDB

func CreateUser(user *Login) error {
	db , err := GetDB()
	tx := db.Begin()
	if err != nil {
		log.Fatalf("An Error occurred while connecting to database: %v", err)
		return err
	} else {
		DB, err := db.DB()
		if err != nil {
			log.Fatalf("Could not find DB: %v", err)
			return err
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

func GetUserByEmail(email string) error {
	db , err := GetDB()
	DB, err := db.DB()
	if err != nil {
		log.Fatalf("An Error occurred while connecting to database: %v", err)
		return err
	} else {
		if err != nil {
			log.Fatalf("Could not find DB: %v", err)
			return err
		}
		defer DB.Close()
		errFirst := db.Debug().Where("email= ?", email).First(&User).Error
		return errFirst
	}
}