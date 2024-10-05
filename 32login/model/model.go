package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Auth struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
