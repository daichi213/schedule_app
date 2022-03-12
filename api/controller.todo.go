package main

import (
	// "fmt"
	"log"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

func ShowTodoIndex(ctx *gin.Context) {
	todos := GetAllTodos()

	ctx.JSON(http.StatusOK, gin.H{
		"page_title": "Index page", 
		"todos":todos,
	},)
}

func ShowTodoContent(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("todo_id")); err == nil {
		// todo_empty := Todo{}
		errFirst := GetTodoByID(articleID)
		if errFirst != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"page_title": "The Page is not found.",
			})
		}else {
			ctx.JSON(http.StatusOK, gin.H{
				"page_title": todo.Title,
				"todo":todo,
			},)
		}
	}else{
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

func CreateTodoContent(ctx *gin.Context){
	// title := ctx.Request.URL.Query().Get("Title")
	// content := ctx.Request.URL.Query().Get("Content")
	// status, err := strconv.Atoi(ctx.Request.URL.Query().Get("Status"))

	var req Todo
	ctx.BindJSON(&req)

	title := req.Title
	content := req.Content
	status := req.Status
	
	log.Println("title:" + title + "\tcontent:" + content)
	log.Println("status:" + ctx.Request.URL.Query().Get("Status"))
	if err != nil{
		log.Fatalf("%s", err)
		// panic(err)
	}
	input := Todo{Title:title, Content:content, Status:status}
	log.Println("input is read.")
	CreateTodo(&input)
	ctx.JSON(http.StatusOK, input)
	return
}

func UpdateTodoContent(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("todo_id")); err == nil {
		// title := ctx.Request.URL.Query().Get("Title")
		// content := ctx.Request.URL.Query().Get("Content")
		// status, err := strconv.Atoi(ctx.Request.URL.Query().Get("Status"))

		var req Todo
		ctx.BindJSON(&req)

		title := req.Title
		content := req.Content
		status := req.Status

		log.Println("title:" + title + "\tcontent:" + content)
		log.Println("status:" + ctx.Request.URL.Query().Get("Status"))
		if err != nil{
			panic(err)
		}
		input := Todo{Title:title, Content:content, Status:status}
		log.Println(status)
		log.Println("input is read.")
		UpdateTodo(articleID, &input)
		ctx.JSON(http.StatusOK, input)
	}else{
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}