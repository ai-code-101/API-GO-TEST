package main

import "fmt"

// in-memory user store
var users = []User{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAllUsers() []User {
	return users
}

func CreateUser(name, email string) User {
	user := User{Name: name, Email: email}
	users = append(users, user)
	return user
}

func UserCount() string {
	return fmt.Sprintf("%d", len(users))
}
