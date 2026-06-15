package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	RegisterRoutes(r)

	log.Println("Server running on http://127.0.0.1:8080")
	r.Run("0.0.0.0:8080")
}
