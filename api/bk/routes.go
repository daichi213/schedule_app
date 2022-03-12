package main

func initializeRoutes() {
	// Root Page
	router.GET("/", ShowTodoIndex)

	// Show Todo Page(Edit Page)
	router.GET("/:todo_id", ShowTodoContent)
	
	// Create Todo
	router.POST("/new", CreateTodoContent)
	
	// Edit Todo
	router.POST("/update/:todo_id", UpdateTodoContent)
}