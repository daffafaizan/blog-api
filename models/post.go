package models

type Post struct {
	ID         string    `json:"id" binding:"required"`
	Title      string    `json:"title" binding:"required,max=60"`
	Content    string    `json:"content" binding:"required,max=1000"`
	Time       string    `json:"time" binding:"required"`
	PostAuthor Author    `json:"postAuthor" binding:"required"`
	Comments   []Comment `json:"comments"`
}
