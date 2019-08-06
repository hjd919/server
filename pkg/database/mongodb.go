package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MDBConfig struct {
	DSN string `mapstructure:"dsn"`
}

var mdbs map[string]*mongo.Database
var mclient *mongo.Client

func init() {
	mdbs = make(map[string]*mongo.Database)
}

func NewMDB(c *MDBConfig, db string) *mongo.Database {

	mdb, ok := mdbs[db]
	if !ok {
		log.Println("初始化mongo db:" + db)
		if mclient == nil {
			log.Println("初始化mongo client")
			mclient = NewMClient(c)
		}
		mdb = mclient.Database(db)
		mdbs[db] = mdb
	}

	return mdb
}

func NewMClient(c *MDBConfig) *mongo.Client {

	if mclient != nil {
		log.Println("已初始化mongo client，返回")
		return mclient
	}

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + c.DSN).SetMaxPoolSize(20)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)

	if err != nil {
		log.Fatal(err)
	}

	return client
}
