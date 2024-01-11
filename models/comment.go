package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	Content       string             `json:"content" bson:"content" binding:"required,max=300"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	CommentAuthor string             `json:"commentAuthor" bson:"commentAuthor" binding:"required"`
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	PostId        primitive.ObjectID `json:"postId" bson:"postId"`
}
