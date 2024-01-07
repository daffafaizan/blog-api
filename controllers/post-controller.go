package controllers

import (
	"net/http"

	"github.com/daffafaizan/blog-api/models"
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
	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	err = controller.service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller postController) GetAllPosts(c *gin.Context) {
	c.JSON(200, "")
}

func (controller postController) GetPostById(c *gin.Context) {
	c.JSON(200, "")
}
