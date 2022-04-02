package api

import (
	"testing"
	"log"
	// "golang.org/x/crypto/bcrypt"
	"github.com/stretchr/testify/assert"
)

// email := "testUser@gin.org"
// username := "testUser"
// invalid := bcrypt.CompareHashAndPassword([]byte(loginVals.Password), UserFromDB.Password)

// TODO GetUserByEmailのメソッド呼び出しで転けているため、テストコードからデバッグする
func TestGetUserByEmail(t *testing.T) {
	email := "testUser@gin.org"
	username := "testUser"
	// should get user by email
	err := GetUserByEmail(email)
	if err != nil {
		log.Fatalf("GetUserByEmail is failed: %v", err)
	}
	act := UserFromDB.UserName
	exp := username
	log.Printf(act,'\n', UserFromDB.Password)
	assert.Equal(t, act, exp)
}