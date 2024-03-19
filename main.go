package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client
var collection *mongo.Collection

type Event struct {
	ID          string `json:"ID,omitempty" bson:"ID,omitempty"`
	Title       string `json:"Title,omitempty" bson:"Title,omitempty"`
	Description string `json:"Description,omitempty" bson:"Description,omitempty"`
}

type EventUpdate struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func CreateEvent(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Starting CreateEvent Function...")
	response.Header().Set("content-type", "application/json")
	var newEvent Event
	err := json.NewDecoder(request.Body).Decode(&newEvent)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	insertResult, err := collection.InsertOne(context.TODO(), newEvent)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(insertResult)
}

func GetAllEvents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []Event
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	if err = cursor.All(ctx, &events); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(events)
}

func main() {
	port := ":9090"
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/c", CreateEvent)
	router.HandleFunc("/a", GetAllEvents)
	fmt.Println("running on http://0.0.0.0" + port)
	log.Fatal(http.ListenAndServe(port, router))

}
