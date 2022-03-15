package api

import (
	"log"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)

type EmailLoginRequest struct {
	gorm.Model
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type AuthUser EmailLoginRequest

var User AuthUser

func CreateUser(user *AuthUser) error {
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

func GetUserByEmail(email string) error {
	db , err := GetDB()
	DB, err := db.DB()
	if err != nil {
		log.Fatalf("An Error occurred while connecting to database: %v", err)
		panic(err)
	} else {
		if err != nil {
			log.Fatalf("Could not find DB: %v", err)
			panic(err)
		}
		defer DB.Close()
		errFirst := db.Debug().Where("email= ?", email).First(&User).Error
		return errFirst
	}
}
