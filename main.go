package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Book struct {
	Title  string
	Author string
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get all users")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get a user")
}

func createUser(w http.ResponseWriter, r *http.Request) {

	col := client.Database("some_database").Collection("Some Collection")

	fmt.Println("Collection type:", reflect.TypeOf(col))

	// doc := Book{Title: "Atonement", Author: "Ian McEwan"}
	// result, err := coll.InsertOne(context.TODO(), doc)
	// fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	// fmt.Println(err)

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update a user")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete a user")
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://0.0.0.0:27017/?timeoutMS=5000")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	defer client.Disconnect(context.Background())

	http.HandleFunc("/users", getUsers)
	http.HandleFunc("/users/", getUser)
	http.HandleFunc("/users/create/", createUser)
	http.HandleFunc("/users/update/", updateUser)
	http.HandleFunc("/users/delete/", deleteUser)

	http.ListenAndServe(":8080", nil)
}
