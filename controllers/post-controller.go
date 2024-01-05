package controllers

import (
	"github.com/daffafaizan/blog-api/models"
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(c *gin.Context) error
	GetAllPosts() []models.Post
	GetPostById(c *gin.Context) *models.Post
}

type postController struct {
	service services.PostService
}

func NewPostController(service services.PostService) PostController {
	return postController{
		service: service,
	}
}

func (controller postController) CreatePost(c *gin.Context) error {
	var post models.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	controller.service.CreatePost(post)
	return nil
}

func (controller postController) GetAllPosts() []models.Post {
	return controller.service.GetAllPosts()
}

func (controller postController) GetPostById(c *gin.Context) *models.Post {
	id := c.Param("id")
	return controller.service.GetPostById(id)
}
