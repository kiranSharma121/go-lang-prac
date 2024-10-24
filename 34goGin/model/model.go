package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GinAuth struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username" validation:"required,min=2,max=100"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
