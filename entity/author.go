package entity

type Author struct {
	Name string `json:"name" binding:"required,max=50"`
	Age  int    `json:"age" binding:"required,gte=1,lte=150"`
}
