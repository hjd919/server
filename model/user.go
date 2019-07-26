package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID     `bson:"_id,omitempty"`
	Username  string                 `bson:"username"`
	Nickname  string                 `bson:"nickname"`
	Avatar    string                 `bson:"avatar"`
	Token     string                 `bson:"token"`
	CreatedAt time.Time              `bson:"created_at"`
	Profile   map[string]interface{} `bson:"profile"`
}
