package entity

type Comment struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Time    string `json:"time"`
	Author  Author `json:"author"`
	Post    Post   `json:"post"`
}
