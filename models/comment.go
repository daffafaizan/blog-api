package models

type Comment struct {
	ID            string `json:"id" binding:"required"`
	Content       string `json:"content" binding:"required,max=300"`
	Date          string `json:"date" binding:"required"`
	Time          string `json:"time" binding:"required"`
	CommentAuthor Author `json:"commentAuthor" binding:"required"`
}
