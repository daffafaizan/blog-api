package controllers

import (
	"github.com/daffafaizan/blog-api/entity"
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(c *gin.Context) error
	GetAllPosts() []entity.Post
	GetPostById(c *gin.Context) (*entity.Post, error)
}

type postController struct {
	service services.PostService
}

func New(service services.PostService) PostController {
	return postController{
		service: service,
	}
}

func (controller postController) CreatePost(c *gin.Context) error {
	var post entity.Post
	err := c.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	controller.service.CreatePost(post)
	return nil
}

func (controller postController) GetAllPosts() []entity.Post {
	return controller.service.GetAllPosts()
}

func (controller postController) GetPostById(c *gin.Context) (*entity.Post, error) {
	id := c.Param("id")
	return controller.service.GetPostById(id)
}
