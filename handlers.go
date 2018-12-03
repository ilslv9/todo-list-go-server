package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func todoShow(writer http.ResponseWriter, request *http.Request) {
	todoId := mux.Vars(request)["todoId"]
	id, err := strconv.Atoi(todoId)
	if err == nil {
		todo := findTodo(id)
		writer.Header().Set("Content-Type",
			"application-json; charset=UTF-8")
		writer.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(writer).Encode(todo); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func todos(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type",
		"application-json; charset=UTF-8")
	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(todoList); err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index!")
}

func todoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := createTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["todoId"]
	id, err := strconv.Atoi(todoId)
	if err == nil {
		err := deleteTodo(id)
		if err == nil {
			w.Header().Set("Content-Type",
				"application-json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(todoList); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func setCompleted(w http.ResponseWriter, r *http.Request) {
	todoId := mux.Vars(r)["todoId"]
	id, err := strconv.Atoi(todoId)
	if err == nil {
		todo, err := repoSetCompleted(id)
		if err == nil {
			w.Header().Set("Content-Type",
				"application-json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(todo); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
