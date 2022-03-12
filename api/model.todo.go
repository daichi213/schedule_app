package main

import (
	"log"
)

func GetAllTodos() []Todo {
	db = GetDB()
	// tx := GetTransaction(db)
	DB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not find DB: %s", err)
	}
	defer DB.Close()
	db.Model(&todo).Find(&todos)
	return todos
}

func GetTodoByID(id int) error {
	db = GetDB()
	// tx := GetTransaction(db)
	DB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not find DB: %s", err)
	}
	defer DB.Close()
	errFirst := db.Debug().Where("id = ?", id).First(&todo).Error
	return errFirst
}

func CreateTodo(data *Todo) {
	db = GetDB()
	tx := GetTransaction(db)
	DB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not find DB: %s", err)
	}
	defer DB.Close()
	if err := tx.Model(&todo).Create(data).Error; err != nil {
		tx.Rollback()
		log.Fatalf("Could not create: %s", err.Error())
	}
	tx.Commit()
}

func UpdateTodo(id int, data *Todo) {
	db = GetDB()
	// tx := GetTransaction(db)
	DB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not find DB: %s", err)
	}
	defer DB.Close()
	db.Model(&todo).Where("id = ?", id).Updates(data)
}