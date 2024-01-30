package main

import (
	"database/sql"
	"log"
)

func DoConnection() *sql.DB {

	db, err := sql.Open("mysql", "docker:password@tcp(0.0.0.0:3306)/golang")

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	return database
}
