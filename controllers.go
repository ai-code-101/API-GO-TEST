package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func GetUsers(ctx *gin.Context) {
	users := GetAllUsers()
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUserHandler(ctx *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required,email"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := CreateUser(req.Name, req.Email)
	ctx.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

// renders the HTML dashboard
func Dashboard(ctx *gin.Context) {
	users := GetAllUsers()
	ctx.HTML(http.StatusOK, "dashboard.html", gin.H{
		"title": "My API Dashboard",
		"users": users,
		"count": UserCount(),
	})
}
