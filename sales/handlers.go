package sales

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

var database *sql.DB

func indexHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Query("select * from golang.products")
	if err != nil {
		fmt.Println("error")
	}
	defer rows.Close()
	products := []Products{}
	for rows.Next() {
		p := Products{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	tmpl, _ := template.ParseFiles("/work_in_que/templates/index.html")
	tmpl.Execute(w, products)
}
