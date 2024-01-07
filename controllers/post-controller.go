package controllers

import (
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(c *gin.Context)
	GetAllPosts(c *gin.Context)
	GetPostById(c *gin.Context)
}

type postController struct {
	service services.PostService
}

func NewPostController(service services.PostService) PostController {
	return postController{
		service: service,
	}
}

func (controller postController) CreatePost(c *gin.Context) {
	c.JSON(200, "")
}

func (controller postController) GetAllPosts(c *gin.Context) {
	c.JSON(200, "")
}

func (controller postController) GetPostById(c *gin.Context) {
	c.JSON(200, "")
}
