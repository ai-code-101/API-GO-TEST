package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	// API routes
	r.GET("/ping", Ping)
	r.GET("/users", GetUsers)
	r.POST("/users", CreateUserHandler)

	// HTML dashboard
	r.GET("/", Dashboard)
}
