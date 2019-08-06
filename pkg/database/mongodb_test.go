package database

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectMongodb() (db *mongo.Database) {
	conf := &MDBConfig{
		DSN: "root:Hjd123%25%5E*@39.96.187.72:27017",
	}
	db = NewMDB(conf, "test")
	return
}

func SelectMongodbCollection() (coll *mongo.Collection) {
	db := ConnectMongodb()
	coll = db.Collection("coll")
	return
}

func TestMongodbConnect(t *testing.T) {
	db := ConnectMongodb()
	if db == nil {
		t.Error(`IsPalindrome("detartrated") = false`)
	}
}

// 增加
func TestMongodbInsertOne(t *testing.T) {
	coll := SelectMongodbCollection()

	result, err := coll.InsertOne(
		context.Background(),
		bson.D{
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
		})
	if err != nil {
		t.Error("InsertOne" + err.Error())
	}
	typeName := reflect.TypeOf(result.InsertedID).String()
	fmt.Println("类型typeName:" + typeName)
	if typeName != "primitive.ObjectID" {
		t.Error("InsertOne TypeOf" + err.Error())
	}
}

func TestMongodbInsertMany(t *testing.T) {
	coll := SelectMongodbCollection()

	docs := []interface{}{
		bson.D{
			{"item", "journal"},
			{"qty", 25},
			{"size", bson.D{
				{"h", 14},
				{"w", 21},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "notebook"},
			{"qty", 50},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "A"},
		},
		bson.D{
			{"item", "paper"},
			{"qty", 100},
			{"size", bson.D{
				{"h", 8.5},
				{"w", 11},
				{"uom", "in"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "planner"},
			{"qty", 75},
			{"size", bson.D{
				{"h", 22.85},
				{"w", 30},
				{"uom", "cm"},
			}},
			{"status", "D"},
		},
		bson.D{
			{"item", "postcard"},
			{"qty", 45},
			{"size", bson.D{
				{"h", 10},
				{"w", 15.25},
				{"uom", "cm"},
			}},
			{"status", "A"},
		},
	}

	result, err := coll.InsertMany(context.Background(), docs)
	if err != nil {
		t.Error("InsertMany" + err.Error())
	}
	for _, insertID := range result.InsertedIDs {
		typeName := reflect.TypeOf(insertID).String()
		fmt.Println("类型typeName:" + typeName)
		if typeName != "primitive.ObjectID" {
			t.Error("InsertMany TypeOf" + err.Error())
		}
	}
}
