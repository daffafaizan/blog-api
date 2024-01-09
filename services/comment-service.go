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
	GetAllCommentsByPostId(*string) ([]*models.Comment, error)
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

func (service *commentService) GetAllCommentsByPostId(postId *string) ([]*models.Comment, error) {
	var comments []*models.Comment
	objectId, err := primitive.ObjectIDFromHex(*postId)
	if err != nil {
		return nil, err
	}

	query := bson.D{bson.E{Key: "postId", Value: objectId}}
	cursor, err := service.commentCollection.Find(service.c, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(service.c) {
		var comment models.Comment
		err := cursor.Decode(&comment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(service.c)
	if len(comments) == 0 {
		return nil, errors.New("posts not found")
	}
	return comments, nil
}

func (service *commentService) CreateComment(postId *string, comment *models.Comment) error {
	comment.ID = primitive.NewObjectID()
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
