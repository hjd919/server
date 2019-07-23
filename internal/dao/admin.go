package dao

import (
	"context"

	"github.com/hjd919/server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *Dao) AdminCheckAuth(username, password string) (admin *model.Admin, err error) {
	filter := bson.M{"username": username, "password": password}
	err = d.MDB.Collection("admin").FindOne(context.Background(), filter).Decode(admin)
	if err != nil {
		return
	}

	return
}
