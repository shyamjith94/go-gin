package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	Id           primitive.ObjectID `bson:"_id" json:"-"`
	CategoryId   string             `json:"category_id,omitempty"`
	CategoryName string             `json:"category_name,omitempty" validate:"required"`
}
