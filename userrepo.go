package main

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var currentUserId int
var users Users

func createUser(user User) string {
	currentUserId += 1
	user.Id = currentUserId
	users = append(users, user)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["id"] = currentUserId
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token.Claims = claims
	tokenString, _ := token.SigningString()
	return tokenString
}
