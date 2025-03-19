package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDb() *mongo.Client {

	mongo_uri := os.Getenv("MONGO_URI")
	if mongo_uri == "" {
		mongo_uri = "mongodb://admin:password@mongo-container:27017"
	}

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create mongodb client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		fmt.Println("Error comnect mongo client")
	}

	// Ping the database to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = client
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("go-crud-api").Collection(collectionName)
}
