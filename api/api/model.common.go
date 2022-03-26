package api

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	_ "github.com/lib/pq"
)

func GetDB() (*gorm.DB, error) {
	postgresqlInfo := fmt.Sprintf("host=db port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DATABASE"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"))

	db, err := gorm.Open(postgres.Open(postgresqlInfo), &gorm.Config{})

	return db, err
}