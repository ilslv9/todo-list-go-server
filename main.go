package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/todos", todos)
	router.GET("/todos/:todoId", todoShow)
	router.POST("/todos", todoCreate)
	router.PUT("/todos/:todoId", setCompleted)
	router.DELETE("/todos/:todoId", todoDelete)
	router.Run(":8080")
}


