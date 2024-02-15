package main

import (
	"database/sql"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"

	reports "work_in_que/reports"
	sales "work_in_que/sales"

	_ "github.com/go-sql-driver/mysql"
)

const PORT = ":8081"

var database *sql.DB

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", sales.IndexHandler)
	router.HandleFunc("/create", sales.AddProduct)
	router.HandleFunc("/edit/{id:[0-9]+}", sales.EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", sales.EditHandler).Methods("POST")
	router.HandleFunc("/delete/{id:[0-9]+}", sales.DeleteHandler)
	router.HandleFunc("/reports/index/", reports.ReportsMainHandler).Methods("GET")

	http.Handle("/", router)

	fmt.Printf("Running on %s \n", PORT)
	http.ListenAndServe(PORT, nil)
}
