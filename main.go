package main

import (
	"database/sql"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

const PORT = ":8081"

var database *sql.DB

func main() {
	db, err := sql.Open("mysql", "docker:password@tcp(0.0.0.0:3306)/golang")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/create", addProduct)
	router.HandleFunc("/edit/{id:[0-9]+}", editPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", editHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", DeleteHandler)

	http.Handle("/", router)

	fmt.Printf("Running on %s \n", PORT)
	http.ListenAndServe(PORT, nil)
}
