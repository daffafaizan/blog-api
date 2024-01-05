package entity

type Post struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Author  Author `json:"author"`
}
