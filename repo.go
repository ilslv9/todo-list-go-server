package main

import "fmt"

var currentId int

var todoList Todos

func init() {
	createTodo(Todo{Name: "lol"})
	createTodo(Todo{Name: "kek"})
}

func findTodo(id int) Todo {
	for _, t := range todoList {
		if t.Id == id {
			return t
		}
	}

	return Todo{}
}

func createTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todoList = append(todoList, t)
	return t
}

func repoSetCompleted(id int) (Todo, error) {
	for i, t := range todoList {
		if t.Id == id {
			todo := Todo{Name: t.Name, Id:t.Id, Due: t.Due, Completed: true}
			todoList = append(todoList[:i], todoList[i+1:]...)
			todoList = append(todoList, todo)
			return todo, nil
		}
	}
	return Todo{}, fmt.Errorf("Todo with id:%d not found", id)
}

func deleteTodo(id int) error {
	for i, t := range todoList {
		if t.Id == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Todo with id:%d not found", id)
}
