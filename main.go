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

	postService = services.NewPostService(postCollection, commentCollection, c)
	commentService = services.NewCommentService(postService, commentCollection, postCollection, c)
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
		apiRoutes.PATCH("/posts/:postId", postController.UpdatePost)
		apiRoutes.GET("/posts/:postId", postController.GetPostById)
		apiRoutes.DELETE("/posts/:postId", postController.DeletePostById)
		apiRoutes.GET("/posts/:postId/comments/:commentId", commentController.GetCommentById)
		apiRoutes.GET("/posts/:postId/comments", commentController.GetAllCommentsByPostId)
		apiRoutes.POST("/posts/:postId/comments", commentController.CreateComment)
		apiRoutes.DELETE("/posts/:postId/comments/:commentId", commentController.DeleteCommentById)
	}

	log.Fatal(server.Run())
}
