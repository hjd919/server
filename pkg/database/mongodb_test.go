package database

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ConnectMongodbClient() (client *mongo.Client) {
	conf := &MDBConfig{
		DSN: "root:Hjd123%25%5E*@39.96.187.72:27017",
	}
	client = NewMClient(conf)
	return
}

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

func TestMongodbTransaction(t *testing.T) {
	client := ConnectMongodbClient()
	ctx := context.Background()
	defer client.Disconnect(ctx)
	testDB := client.Database("test")
	col := testDB.Collection("test")
	var err error
	//先在事务外写一条id为“111”的记录
	/* 	_, err := col.InsertOne(ctx, bson.M{"_id": "111", "name": "ddd", "age": 50})
	   	if err != nil {
	   		t.Error("InsertOne TypeOf" + err.Error())
	   		return
	   	}
	   	session, err := client.StartSession()
	   	if err != nil {
	   		t.Error("InsertOne TypeOf" + err.Error())
	   	}

	   	//开始事务
	   	err = session.StartTransaction()
	   	if err != nil {
	   		fmt.Println(err)
	   		return
	   	}

	   	//在事务内写一条id为“222”的记录
	   	_, err = col.InsertOne(ctx, bson.M{"_id": "222", "name": "ddd", "age": 50})
	   	if err != nil {
	   		fmt.Println(err)
	   		return
	   	} */

	// //写重复id
	// _, err = col.InsertOne(ctx, bson.M{"_id": "111", "name": "ddd", "age": 50})
	// if err != nil {
	// 	session.AbortTransaction(ctx)
	// } else {
	// 	session.CommitTransaction(ctx)
	// }

	//第一个事务：成功执行
	client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err = sessionContext.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“222”的记录
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "555", "name": "ddd", "age": 50})
		if err != nil {
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“333”的记录
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "333", "name": "ddd", "age": 50})
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			sessionContext.CommitTransaction(sessionContext)
		}
		return nil
	})

	//第二个事务：执行失败，事务没提交，因最后插入了一条重复id "111",
	err = client.UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}

		//在事务内写一条id为“222”的记录
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "444", "name": "ddd", "age": 50})
		if err != nil {
			fmt.Println(err)
			return err
		}

		//写重复id
		_, err = col.InsertOne(sessionContext, bson.M{"_id": "111", "name": "ddd", "age": 50})
		if err != nil {
			sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			sessionContext.CommitTransaction(sessionContext)
		}
		return nil
	})

}
