package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Content       string             `json:"content" bson:"content" binding:"required,max=300"`
	Date          string             `json:"date" bson:"date" binding:"required"`
	Time          string             `json:"time" bson:"time" binding:"required"`
	CommentAuthor Author             `json:"commentAuthor" bson:"commentAuthor" binding:"required"`
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	PostId        primitive.ObjectID `json:"postId" bson:"postId" binding:"required"`
}
