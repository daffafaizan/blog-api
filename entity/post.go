package entity

type Post struct {
	PostID  string `json:"postid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Author  Author `json:"author"`
}
