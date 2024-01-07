package services

import (
	"context"

	"github.com/daffafaizan/blog-api/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentService interface {
	CreateComment(*string, *models.Comment) error
}

type commentService struct {
	postService       PostService
	commentCollection *mongo.Collection
	c                 context.Context
}

func NewCommentService(postService PostService, commentCollection *mongo.Collection, c context.Context) CommentService {
	return &commentService{
		postService:       postService,
		commentCollection: commentCollection,
		c:                 c,
	}
}

func (service *commentService) CreateComment(postId *string, comment *models.Comment) error {
	_, err := service.commentCollection.InsertOne(service.c, comment)
	if err != nil {
		return err
	}

	err = service.postService.CreateComment(postId, comment)
	if err != nil {
		return err
	}

	return nil
}
