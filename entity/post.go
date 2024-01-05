package entity

import "time"

type Author struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Post struct {
	Date    time.Time `json:"date"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Author  string    `json:"author"`
}
