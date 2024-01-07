package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	Name        string             `json:"name" bson:"name" binding:"required,max=50"`
	Institution string             `json:"institution" bson:"institution" binding:"max=60"`
	Occupation  string             `json:"occupation" bson:"occupation" binding:"gte=1,lte=150"`
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
}
