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
	DeletePostById(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = controller.service.CreatePost(&post)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller postController) GetAllPosts(c *gin.Context) {
	posts, err := controller.service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (controller postController) GetPostById(c *gin.Context) {
	postId := c.Param("id")
	post, err := controller.service.GetPostById(&postId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (controller postController) DeletePostById(c *gin.Context) {
	postId := c.Param("id")
	err := controller.service.DeletePostById(&postId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
