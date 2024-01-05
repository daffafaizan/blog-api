package controllers

import (
	"github.com/daffafaizan/blog-api/models"
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	CreateComment(c *gin.Context) error
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return commentController{
		service: service,
	}
}

func (controller commentController) CreateComment(c *gin.Context) error {
	var comment models.Comment
	postId := c.Param("id")
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		return err
	}
	controller.service.CreateComment(comment, postId)
	return nil
}
