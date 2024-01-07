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
	GetAllPosts() ([]*models.Post, error)
	GetPostById(*string) (*models.Post, error)
}

type postService struct {
	postCollection *mongo.Collection
	c              context.Context
}

func NewPostService(postCollection *mongo.Collection, c context.Context) PostService {
	return &postService{
		postCollection: postCollection,
		c:              c,
	}
}

func (service *postService) CreatePost(post *models.Post) error {
	_, err := service.postCollection.InsertOne(service.c, post)
	return err
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

func (service *postService) GetPostById(id *string) (*models.Post, error) {
	var post *models.Post
	objectId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}
	filter := bson.D{bson.E{Key: "_id", Value: objectId}}
	err = service.postCollection.FindOne(service.c, filter).Decode(&post)
	return post, err
}
