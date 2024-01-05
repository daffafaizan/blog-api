package main

import (
	"net/http"

	"github.com/daffafaizan/blog-api/controllers"
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-gonic/gin"
)

var (
	postService    services.PostService       = services.NewPostService()
	postController controllers.PostController = controllers.NewPostController(postService)
)

func main() {
	server := gin.Default()

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", func(c *gin.Context) {
			c.JSON(200, postController.GetAllPosts())
		})
		apiRoutes.POST("/posts", func(c *gin.Context) {
			err := postController.CreatePost(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Post successfully created!"})
			}
		})
		apiRoutes.GET("/posts/:id", func(c *gin.Context) {
			post, err := postController.GetPostById(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, post)
			}
		})
	}

	server.Run("localhost:8080")
}
