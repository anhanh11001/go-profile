package middlewares

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Constants on MongoDB
const connectionString = "TODO: Check connection string"
const dbName = "TODO: check DB name"
const collName = "TODO: check collection name"

// collection object/instance

func SetUpMongoDb() *mongo.Collection {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	var collection *mongo.Collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created!")
	return collection
}
