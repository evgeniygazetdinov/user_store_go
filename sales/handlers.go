package sales

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/mux"
)

func DoConnection() *sql.DB {

	db, err := sql.Open("mysql", "docker:password@tcp(0.0.0.0:3306)/golang")

	if err != nil {
		log.Println(err)
	}
	database := db
	defer db.Close()
	return database
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	database := DoConnection()
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
			fmt.Println("here!!!!")
		}
		products = append(products, p)
	}
	tmpl, _ := template.ParseFiles("./sales/templates/index.html")
	tmpl.Execute(w, products)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	database := DoConnection()
	_, err := database.Exec("delete from golang.products where id = ?", id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	database := DoConnection()
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		model := r.FormValue("model")
		company := r.FormValue("company")
		price := r.FormValue("price")

		_, err = database.Exec("insert into golang.products(model, company, price) values (?, ?,?)", model, company, price)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "/sales/templates/create.html")
	}
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	database := DoConnection()
	vars := mux.Vars(r)
	id := vars["id"]
	prod := Products{}
	row := database.QueryRow("select * from golang.products where id = ?", id)
	err := row.Scan(&prod.Id, &prod.Model, &prod.Company, &prod.Price)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("/sales/templates/edit.html")
		tmpl.Execute(w, prod)
	}
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	database := DoConnection()
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		id := r.FormValue("id")
		model := r.FormValue("model")
		company := r.FormValue("company")
		price := r.FormValue("price")

		_, err = database.Exec("update golang.products set model = ?, company = ?, price = ? where id = ?", model, company, price, id)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	}
}
