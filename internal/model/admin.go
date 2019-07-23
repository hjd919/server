package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Admin
type Admin struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username string
	Password string
}
