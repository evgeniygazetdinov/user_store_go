package main

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"

	"work_in_que/manage"
	sales "work_in_que/sales"

	_ "github.com/go-sql-driver/mysql"
)

const PORT = ":8081"

var database *sql.DB

func main() {
	manage.CreateFile()
	router := mux.NewRouter()
	router.HandleFunc("/", sales.IndexHandler)
	router.HandleFunc("/create", sales.AddProduct)
	router.HandleFunc("/edit/{id:[0-9]+}", sales.EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", sales.EditHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", sales.DeleteHandler)

	http.Handle("/", router)

	fmt.Printf("Running on %s \n", PORT)
	http.ListenAndServe(PORT, nil)
}
