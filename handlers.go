package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func todoShow(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.Atoi(todoId)
	if err != nil {
		panic(err)
	}
	todo := findTodo(id)
	ctx.JSON(200, &todo)
}

func todos(ctx *gin.Context) {
	ctx.JSON(200, &todoList)
}

func index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Index!")
}

func todoCreate(ctx *gin.Context) {
	var todo Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		panic(err)
	}
	t := createTodo(todo)
	ctx.JSON(200, &t)
}

func todoDelete(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.Atoi(todoId)
	if err == nil {
		err := deleteTodo(id)
		if err == nil {
			ctx.JSON(200, &todoList)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func setCompleted(ctx *gin.Context) {
	todoId := ctx.Param("todoId")
	id, err := strconv.Atoi(todoId)
	if err == nil {
		todo, err := repoSetCompleted(id)
		if err == nil {
			ctx.JSON(200, &todo)
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
