package main

import (
	// "github.com/DATA-DOG/go-sqlmock"
	// "time"
	"testing"
	"log"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTodos(t *testing.T) {
	act := GetAllTodos()

	db = GetDB()
	DB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Model(&todo).Find(&todos)
	ext := todos
	defer DB.Close()
	assert.Equal(t, len(ext), len(act))
}

func TestGetTodo(t *testing.T) {
	GetTodoByID(7)
	act := todo.Title

	target := Todo{Title: "test2", Content:"This message is in test2.", Status: 1}
	ext := target.Title
	assert.Equal(t, ext, act)
}

func TestCreateTodoContent(t *testing.T){
	nowTime := GetTimeNumber()
	target := Todo{Title: nowTime, Content:"This message is in test2 at " + nowTime, Status: 1}
	ext := target.Title

	CreateTodo(&target)

	db = GetDB()
	DB, err := db.DB()
	defer DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	db.Debug().Last(&todo)
	act := todo.Title
	assert.Equal(t, ext, act)
}

func TestPutTodoContent(t *testing.T) {
	nowTime := GetTimeNumber()
	target := Todo{Title: nowTime, Content:"This message is in test2.", Status: 1}
	ext := target.Title
	UpdateTodo(2, &target)

	db = GetDB()
	DB, err := db.DB()
	defer DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	db.Debug().Where("id=?", 2).First(&todo)
	act := todo.Title
	assert.Equal(t, ext, act)
}