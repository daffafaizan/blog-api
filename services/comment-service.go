package services

import (
	"github.com/daffafaizan/blog-api/models"
)

type CommentService interface {
	CreateComment(models.Comment, string) models.Comment
}

type commentService struct {
	postService PostService
}

func NewCommentService(postService PostService) CommentService {
	return &commentService{
		postService: postService,
	}
}

func (service *commentService) CreateComment(comment models.Comment, postId string) models.Comment {
	post := service.postService.GetPostById(postId)
	post.Comments = append(post.Comments, comment)
	return comment
}
