package services

import "github.com/daffafaizan/blog-api/models"

type CommentService interface {
	CreateComment(models.Comment) models.Comment
}

type commentService struct {
	comments []models.Comment
}

func NewCommentService() CommentService {
	return &commentService{}
}

func (service *commentService) CreateComment(comment models.Comment) models.Comment {
	service.comments = append(service.comments, comment)
	return comment
}
