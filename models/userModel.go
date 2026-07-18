package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"Name,omitempty" bson:"Name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=6"`
	ImageUrl  string             `json:"imageUrl,omitempty" bson:"imageUrl,omitempty"`
	Bio       string             `json:"bio,omitempty" bson:"bio,omitempty"`
	Followers []string           `json:"followers,omitempty" bson:"followers,omitempty"`
	Following []string           `json:"following,omitempty" bson:"following,omitempty"`
}
