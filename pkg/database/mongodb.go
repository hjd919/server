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

func NewMDB(c *MDBConfig) *mongo.Database {
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

	return client.Database("jishua")
}
