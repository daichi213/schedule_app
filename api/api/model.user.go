package api

import (
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
	_ "github.com/lib/pq"
)

// Clientから受け取る構造体
type Login struct {
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Email string `form:"Email" json:"Email" binding:"required"`
	Password string `form:"Password" json:"Password" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag"`
}

// Tableへ格納するために使用する構造体
type User struct {
	ID uint `gorm:"primary_key"`
	UserName string `form:"UserName" json:"UserName" binding:"required"`
	Email string `form:"Email" json:"Email" binding:"required"`
	Password []byte `form:"Password" json:"Password" binding:"required"`
	AdminFlag int `form:"AdminFlag" json:"AdminFlag" binding:"required"`
	CreatedAt time.Time
  	UpdatedAt time.Time
  	DeletedAt *time.Time
}

var UserFromDB User
var UserToDB User

func CreateUser(user *Login) error {
	db , err := GetDB()
	tx := db.Begin()
	// TODO エラーハンドリングのelse文が煩雑になってしまっているので、テストコードを書いてリファクタリングする
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

		// TODO パスワードのハッシュ化のための関数を個別で定義する
		// また、AdminFlagを引数から制御できるように実装する
		pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("An Error occurred while hashing the password: %v", err)
			return err
		} else {
			UserToDB = User{UserName: user.UserName, Email: user.Email, Password: pass}
		}
		if err := tx.Model(&UserToDB).Create(&UserToDB).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Could not create: %v", err)
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
		errFirst := db.Debug().Where("email= ?", email).First(&UserFromDB).Error
		return errFirst
	}
}