package services

import (
	"errors"

	"github.com/daffafaizan/blog-api/entity"
)

type PostService interface {
	CreatePost(entity.Post) entity.Post
	GetAllPosts() []entity.Post
	GetPostById(string) (*entity.Post, error)
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

func (service *postService) GetPostById(postId string) (*entity.Post, error) {
	for i, p := range service.posts {
		if p.PostID == postId {
			return &service.posts[i], nil
		}
	}
	return nil, errors.New("post not found")
}
