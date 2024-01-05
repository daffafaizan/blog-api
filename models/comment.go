package models

type Comment struct {
	ID      string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required,max=300"`
	Time    string `json:"time" binding:"required"`
	PostID  string `json:"postid" binding:"required"`
	Author  Author `json:"author" binding:"required"`
}
