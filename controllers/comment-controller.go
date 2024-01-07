package controllers

import (
	"net/http"

	"github.com/daffafaizan/blog-api/models"
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(c *gin.Context)
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return commentController{
		service: service,
	}
}

func (controller commentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	postId := c.Param("id")
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	err = controller.service.CreateComment(&postId, &comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
