package services

import (
	"context"
	"errors"

	"github.com/daffafaizan/blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService interface {
	CreatePost(*models.Post) error
	UpdatePost(*string, *models.Post) error
	CreateComment(*string, *models.Comment) error
	GetAllPosts() ([]*models.Post, error)
	GetPostById(*string) (*models.Post, error)
	DeletePostById(*string) error
}

type postService struct {
	postCollection    *mongo.Collection
	commentCollection *mongo.Collection
	c                 context.Context
}

func NewPostService(postCollection *mongo.Collection, commentCollection *mongo.Collection, c context.Context) PostService {
	return &postService{
		postCollection:    postCollection,
		commentCollection: commentCollection,
		c:                 c,
	}
}

func (service *postService) CreatePost(post *models.Post) error {
	post.ID = primitive.NewObjectID()
	_, err := service.postCollection.InsertOne(service.c, post)
	return err
}

func (service *postService) UpdatePost(postId *string, post *models.Post) error {
	objectId, err := primitive.ObjectIDFromHex(*postId)
	if err != nil {
		return err
	}

	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "title", Value: post.Title}, bson.E{Key: "summary", Value: post.Summary}, bson.E{Key: "content", Value: post.Content}, bson.E{Key: "tags", Value: post.Tags}}}}
	result, _ := service.postCollection.UpdateOne(service.c, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched post found for update")
	}
	return nil
}

func (service *postService) GetAllPosts() ([]*models.Post, error) {
	var posts []*models.Post
	query := bson.D{{}}
	cursor, err := service.postCollection.Find(service.c, query)
	if err != nil {
		return nil, err
	}
	for cursor.Next(service.c) {
		var post models.Post
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(service.c)
	if len(posts) == 0 {
		return nil, errors.New("posts not found")
	}
	return posts, nil
}

func (service *postService) GetPostById(postId *string) (*models.Post, error) {
	var post *models.Post
	objectId, err := primitive.ObjectIDFromHex(*postId)
	if err != nil {
		return nil, err
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	err = service.postCollection.FindOne(service.c, filter).Decode(&post)
	return post, err
}

func (service *postService) CreateComment(postId *string, comment *models.Comment) error {
	postObjectId, err := primitive.ObjectIDFromHex(*postId)
	if err != nil {
		return err
	}

	post, err := service.GetPostById(postId)
	if err != nil {
		return err
	}
	post.Comments = append(post.Comments, *comment)

	filter := bson.D{bson.E{Key: "_id", Value: postObjectId}}
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

func (service *postService) DeletePostById(postId *string) error {
	postObjectId, err := primitive.ObjectIDFromHex(*postId)
	if err != nil {
		return err
	}

	postFilter := bson.D{bson.E{Key: "_id", Value: postObjectId}}
	postResult, _ := service.postCollection.DeleteOne(service.c, postFilter)

	if postResult.DeletedCount != 1 {
		return errors.New("no matched post found for delete")
	}

	commentFilter := bson.D{bson.E{Key: "postId", Value: postObjectId}}
	_, err = service.commentCollection.DeleteMany(service.c, commentFilter)
	if err != nil {
		return err
	}

	return nil
}
