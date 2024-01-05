package models

type Author struct {
	Name        string `json:"name" binding:"required,max=50"`
	Institution string `json:"institution" binding:"max=60"`
	Occupation  string `json:"occupation" binding:"gte=1,lte=150"`
}
