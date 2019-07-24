package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// MID
type MID struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
