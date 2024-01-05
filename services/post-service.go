package services

import (
	"github.com/daffafaizan/blog-api/entity"
	"github.com/gin-gonic/gin"
)

type PostService interface {
	CreatePost(entity.Post) entity.Post
	GetAllPost(c *gin.Context) []entity.Post
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

func (service *postService) GetAllPost(c *gin.Context) []entity.Post {
	return service.posts
}
