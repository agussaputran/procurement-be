package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Migrate() {
	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensure the context is canceled to avoid resource leaks

	database := Client.Database("prucrement")

	// Create 'users' collection if it doesn't exist
	collectionNames, err := database.ListCollectionNames(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	collectionExists := false
	for _, name := range collectionNames {
		if name == "products" {
			collectionExists = true
			break
		}
	}

	if !collectionExists {
		err := database.CreateCollection(ctx, "products", options.CreateCollection())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Created 'products' collection")
	} else {
		log.Println("'products' collection already exists")
	}
}
