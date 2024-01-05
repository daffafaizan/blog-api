package entity

type Comment struct {
	Content string `json:"content"`
	Author  Author `json:"author"`
	Post    Post   `json:"post"`
}
