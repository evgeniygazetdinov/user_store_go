package getcollection

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	collection := client.Database("ctx").Collection("Posts")
	fmt.Println("after collection")
	return collection
}
