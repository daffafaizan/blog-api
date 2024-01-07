package services

import (
	"context"
	"errors"

	"github.com/daffafaizan/blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentService interface {
	CreateComment(*models.Comment, *string) error
}

type commentService struct {
	postService    PostService
	postCollection *mongo.Collection
	c              context.Context
}

func NewCommentService(postService PostService, postCollection *mongo.Collection, c context.Context) CommentService {
	return &commentService{
		postService:    postService,
		postCollection: postCollection,
		c:              c,
	}
}

func (service *commentService) CreateComment(comment *models.Comment, postId *string) error {
	post, err := service.postService.GetPostById(postId)
	if err != nil {
		return err
	}
	post.Comments = append(post.Comments, *comment)
	filter := bson.D{bson.E{Key: "id", Value: postId}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "comments", Value: post.Comments}}}}
	result, err := service.postCollection.UpdateOne(service.c, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("no matched post found for update")
	}
	return nil
}
