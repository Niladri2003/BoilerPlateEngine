package db

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToMongo() error {
	var err error
	// Use the connection string for MongoDB running in the Docker container
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://root:rootpass@localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB
	err = Client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Defer disconnection when main function ends

	fmt.Println("Connected to MongoDB!")
	return nil
}
