package api

import (
	"os"
	"fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
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

func getDB() (*gorm.DB, error) {
	postgresqlInfo := fmt.Sprintf("host=db port=${POSTGRES_PORT} dbname=${POSTGRES_DATABASE} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} sslmode=disable",
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	db, err := gorm.Open(postgres.Open(postgresqlInfo), &gorm.Config{})

	return db, err
}

func GetUserByEmail(email string) error {
	db , err := getDB()
	if err != nil {
		log.Fatalf("An Error occurred while connecting to database: %v", err)
		panic(err)
	} else {
		tx, err := db.DB()
		if err != nil {
			log.Fatalf("Could not find DB: %v", err)
			panic(err)
		}
		defer tx.Close()
		errFirst := db.Debug().Where("email= ?", email).First(&User).Error
		return errFirst
	}
}
