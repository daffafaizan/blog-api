package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Title      string             `json:"title" bson:"title" binding:"required,max=60"`
	Slug       string             `json:"slug" bson:"slug"`
	Summary    string             `json:"summary" bson:"summary" binding:"required,max=300"`
	Content    string             `json:"content" bson:"content" binding:"required"`
	CreatedAt  time.Time          `json:"createdAt" bson:"createdAt"`
	PostAuthor string             `json:"postAuthor" bson:"postAuthor" binding:"required"`
	Tags       []string           `json:"tags" bson:"tags" binding:"required"`
	Comments   []Comment          `json:"comments" bson:"comments"`
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
}
