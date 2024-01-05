package entity

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Date    string `json:"date"`
	Author  Author `json:"author"`
}
