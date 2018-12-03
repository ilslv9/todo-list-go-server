package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {

		var handler http.Handler

		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).HandlerFunc(route.HandlerFunc)
	}

	return router
}
