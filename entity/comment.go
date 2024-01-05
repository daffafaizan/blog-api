package entity

type Comment struct {
	CommentID string `json:"commentid"`
	Content   string `json:"content"`
	Time      string `json:"time"`
	Author    Author `json:"author"`
	Post      Post   `json:"post"`
}
