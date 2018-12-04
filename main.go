package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/auth", auth)
	router.GET("/", index)
	router.GET("/todos", todos)
	router.GET("/todos/:todoId", todoShow)
	router.POST("/todos", todoCreate)
	router.PUT("/todos/:todoId", setCompleted)
	router.DELETE("/todos/:todoId", todoDelete)
	router.Run(":8080")
}
