package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://michael:secret@localhost:27017/"),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to mongoDBs")
	return client
}
