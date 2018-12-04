package main

import "time"

type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type Users []User

type Todos []Todo