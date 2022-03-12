package main

import (
	// "database/sql"
	"os"
	"fmt"
	// "log"
	// "strconv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	// gorm.ModelはIDやcreated_atなどが予め定義されている
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Todo struct {
	// gorm.ModelはIDやcreated_atなどが予め定義されている
	gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	Status int `json:"status"`
	User_id int `json:"user_id"`
}

var todo Todo
var todos []Todo
var user User
var users []User

func GetDB() *gorm.DB {
	// テスト環境と開発環境でDBの接続先を変更する
	// if testMode, err := strconv.ParseBool(os.Getenv("TEST_MODE")); testMode == false {
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	mysqlInfo := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
	// 		os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	// } else {
	// 	mysqlInfo := fmt.Sprintf("root:%s@tcp(localhost:%s)/%s?charset=utf8&parseTime=True",
	// 		os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	// }

	// Dev Mode
	// mysqlInfo := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
	// 	os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))

	// Test Mode
	mysqlInfo := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_TEST_DATABASE"))
	
	db, err = gorm.Open(mysql.Open(mysqlInfo), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}

func GetTransaction(db *gorm.DB) *gorm.DB {
	return db.Begin()
}

func InitDb() {
	db = GetDB()
	DB, err := db.DB()
	defer DB.Close()
	if err != nil {
		panic(err)
	}
	if hasDB := db.Migrator().HasTable("todo"); hasDB ==false {
		db.Migrator().CreateTable(&Todo{})
	}
}
