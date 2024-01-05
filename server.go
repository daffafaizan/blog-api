package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", func(c *gin.Context) {
		})
	}

	server.Run("localhost:8080")
}
