package services

import (
	"context"

	"github.com/daffafaizan/blog-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostService interface {
	CreatePost(*models.Post) error
	GetAllPosts() (*[]models.Post, error)
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

func (service *postService) GetAllPosts() (*[]models.Post, error) {
	return nil, nil
}

func (service *postService) GetPostById(id *string) (*models.Post, error) {
	var post *models.Post
	query := bson.D{bson.E{Key: "id", Value: id}}
	err := service.postCollection.FindOne(service.c, query).Decode(&post)
	return post, err
}
