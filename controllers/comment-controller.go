package controllers

import (
	"net/http"

	"github.com/daffafaizan/blog-api/models"
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-gonic/gin"
)

type CommentController interface {
	GetCommentById(c *gin.Context)
	GetAllCommentsByPostId(c *gin.Context)
	CreateComment(c *gin.Context)
	DeleteCommentById(c *gin.Context)
}

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) CommentController {
	return commentController{
		service: service,
	}
}

func (controller commentController) GetCommentById(c *gin.Context) {
	commentId := c.Param("commentId")
	comment, err := controller.service.GetCommentById(&commentId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

func (controller commentController) GetAllCommentsByPostId(c *gin.Context) {
	postId := c.Param("postId")
	comments, err := controller.service.GetAllCommentsByPostId(&postId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, comments)
}

func (controller commentController) CreateComment(c *gin.Context) {
	var comment models.Comment
	postId := c.Param("postId")
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

func (controller commentController) DeleteCommentById(c *gin.Context) {
	postId := c.Param("postId")
	commentId := c.Param("commentId")
	err := controller.service.DeleteCommentById(&commentId, &postId)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
