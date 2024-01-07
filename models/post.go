package models

type Post struct {
	ID         string    `json:"id" bson:"id" binding:"required"`
	Title      string    `json:"title" bson:"title" binding:"required,max=60"`
	Content    string    `json:"content" bson:"content" binding:"required,max=1000"`
	Date       string    `json:"date" bson:"date" binding:"required"`
	Time       string    `json:"time" bson:"time" binding:"required"`
	PostAuthor Author    `json:"postAuthor" bson:"postAuthor" binding:"required"`
	Comments   []Comment `json:"comments" bson:"comments"`
}
