package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{"Index",
		"GET",
		"/",
		index},
	{"TodoIndex",
		"GET",
		"/todos",
		todos},
	{"TodoShow",
		"GET",
		"/todos/{todoId}",
		todoShow},
	{"TodoCreate",
		"POST",
		"/todos",
		todoCreate},
	{"TodoDelete",
		"DELETE",
		"/todos/{todoId}",
		todoDelete},
	{"TodoSetCompleted",
		"PUT",
		"/todos/{todoId}",
		setCompleted},
}
