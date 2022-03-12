package main

import (
	"testing"
)

func TestMain(m *testing.M) {
	db := GetDB()
	DB, err := db.DB()
	SetupFindFuncs()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	m.Run()
}
