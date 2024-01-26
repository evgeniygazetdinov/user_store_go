package sales

import (
	"database/sql"

	// "encoding/json"

	"html/template"
	"net/http"
)

var database *sql.DB

func SalesindexHandler(w http.ResponseWriter, r *http.Request) {
	//TODO fix db calling from here
	// rows, err := database.Query("select * from golang.products")
	// if err != nil {
	// 	fmt.Println("error")
	// }
	// defer rows.Close()
	// products := []Products{}
	// for rows.Next() {
	// 	p := Products{}
	// 	err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	products = append(products, p)
	// }
	p := Products{}
	tmpl, _ := template.ParseFiles("./templates/index.html")
	tmpl.Execute(w, p)
	// data := Products{}
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusCreated)
	// json.NewEncoder(w).Encode(data)
}
