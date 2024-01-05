package services

import (
	"github.com/daffafaizan/blog-api/entity"
)

type PostService interface {
	CreatePost(entity.Post) entity.Post
	GetAllPosts() []entity.Post
}

type postService struct {
	posts []entity.Post
}

func New() PostService {
	return &postService{}
}

func (service *postService) CreatePost(post entity.Post) entity.Post {
	service.posts = append(service.posts, post)
	return post
}

func (service *postService) GetAllPosts() []entity.Post {
	return service.posts
}
