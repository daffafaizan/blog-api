package main

import (
	"context"
	"log"
	"os"

	"github.com/daffafaizan/blog-api/controllers"
	"github.com/daffafaizan/blog-api/initializers"
	"github.com/daffafaizan/blog-api/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	postService       services.PostService
	commentService    services.CommentService
	postController    controllers.PostController
	commentController controllers.CommentController

	postCollection    *mongo.Collection
	commentCollection *mongo.Collection
	mongoClient       *mongo.Client

	c      context.Context
	server *gin.Engine
)

func init() {
	initializers.LoadEnv()
	c = context.TODO()

	mongoConn := options.Client().ApplyURI(os.Getenv("MONGODB"))
	mongoClient, err := mongo.Connect(c, mongoConn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(c, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	postCollection = mongoClient.Database("blogdb").Collection("posts")
	commentCollection = mongoClient.Database("blogdb").Collection("comments")

	postService = services.NewPostService(postCollection, c)
	commentService = services.NewCommentService(postService, commentCollection, c)
	postController = controllers.NewPostController(postService)
	commentController = controllers.NewCommentController(commentService)

	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(c)

	server.Use(cors.Default())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/posts", postController.GetAllPosts)
		apiRoutes.POST("/posts", postController.CreatePost)
		apiRoutes.GET("/posts/:id", postController.GetPostById)
		apiRoutes.POST("/posts/:id/comment", commentController.CreateComment)
	}

	log.Fatal(server.Run(os.Getenv("SERVER")))
}
