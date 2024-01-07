package models

type Comment struct {
	ID            string `json:"id" bson:"id" binding:"required"`
	Content       string `json:"content" bson:"content" binding:"required,max=300"`
	Date          string `json:"date" bson:"date" binding:"required"`
	Time          string `json:"time" bson:"time" binding:"required"`
	CommentAuthor Author `json:"commentAuthor" bson:"commentAuthor" binding:"required"`
}
