package services

import (
	"context"
	"errors"

	"github.com/daffafaizan/blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentService interface {
	GetCommentById(*string) (*models.Comment, error)
	CreateComment(*string, *models.Comment) error
	DeleteCommentById(*string, *string) error
}

type commentService struct {
	postService       PostService
	commentCollection *mongo.Collection
	postCollection    *mongo.Collection
	c                 context.Context
}

func NewCommentService(postService PostService, commentCollection *mongo.Collection, postCollection *mongo.Collection, c context.Context) CommentService {
	return &commentService{
		postService:       postService,
		commentCollection: commentCollection,
		postCollection:    postCollection,
		c:                 c,
	}
}

func (service *commentService) GetCommentById(commentId *string) (*models.Comment, error) {
	var comment *models.Comment
	objectId, err := primitive.ObjectIDFromHex(*commentId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	err = service.commentCollection.FindOne(service.c, filter).Decode(&comment)
	return comment, err
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

func (service *commentService) DeleteCommentById(commentId, postId *string) error {
	postObjectId, err := primitive.ObjectIDFromHex(*postId)
	if err != nil {
		return err
	}

	commentObjectId, err := primitive.ObjectIDFromHex(*commentId)
	if err != nil {
		return err
	}

	commentFilter := bson.D{bson.E{Key: "_id", Value: commentObjectId}}
	commentResult, _ := service.commentCollection.DeleteOne(service.c, commentFilter)
	if commentResult.DeletedCount != 1 {
		return errors.New("no matched comment for delete")
	}

	postFilter := bson.D{bson.E{Key: "_id", Value: postObjectId}, bson.E{Key: "comments._id", Value: commentObjectId}}
	postUpdate := bson.D{bson.E{Key: "$pull", Value: bson.D{bson.E{Key: "comments", Value: bson.D{bson.E{Key: "_id", Value: commentObjectId}}}}}}
	postResult, err := service.postCollection.UpdateOne(service.c, postFilter, postUpdate)
	if err != nil {
		return err
	}

	if postResult.MatchedCount != 1 {
		return errors.New("no matched comment for delete")
	}

	return nil
}
