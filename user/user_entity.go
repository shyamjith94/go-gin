package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       string             `json:"userid,omitempty"`
	UserName     string             `json:"username,omitempty" validate:"required"`
	FirstName    string             `json:"firstname,omitempty" validate:"required"`
	LastName     string             `json:"lastname,omitempty" validate:"required"`
	Email        string             `json:"email,omitempty" validate:"required"`
	Phone        int                `json:"phone,omitempty" validate:"required"`
	Location     string             `json:"location,omitempty" validate:"required"`
	Password     string             `json:"password,omitempty" validate:"required"`
	CreatedAt    time.Time          `json:"created_at,omitempty"`
	UpdatedAt    time.Time          `json:"updated_at,omitempty"`
	Token        string             `json:"token,omitempty"`
	RefreshToken string             `json:"refresh_token,omitempty"`
}

type Login struct {
	UserName string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
