package sales

import (
	"database/sql"
	"fmt"
	"log"
)

// 	"github.com/gorilla/mux"
// )
func DoConnection() *sql.DB {

	db, err := sql.Open("mysql", "docker:password@tcp(0.0.0.0:3306)/golang")

	if err != nil {
		log.Println(err)
	}
	database := db
	return database
}

func MainService() []Products {
	database := DoConnection()
	rows, err := database.Query("select * from golang.products")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)
	products := []Products{}
	for rows.Next() {
		p := Products{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println("here!!!!")
			fmt.Println(err)

		}
		products = append(products, p)
	}
	return products
}
