package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employe struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName string             `json:"fullname,omitempty"`
	Present  bool               `json:"present,omitempty"`
}
