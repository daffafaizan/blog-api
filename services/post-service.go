package services

import (
	"errors"

	"github.com/daffafaizan/blog-api/models"
)

type PostService interface {
	CreatePost(models.Post) models.Post
	GetAllPosts() []models.Post
	GetPostById(string) (*models.Post, error)
}

type postService struct {
	posts []models.Post
}

func New() PostService {
	return &postService{}
}

func (service *postService) CreatePost(post models.Post) models.Post {
	service.posts = append(service.posts, post)
	return post
}

func (service *postService) GetAllPosts() []models.Post {
	return service.posts
}

func (service *postService) GetPostById(id string) (*models.Post, error) {
	for i, p := range service.posts {
		if p.ID == id {
			return &service.posts[i], nil
		}
	}
	return nil, errors.New("post not found")
}
