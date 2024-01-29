package sales

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var database *sql.DB

// func SalesindexHandler(w http.ResponseWriter, r *http.Request) {
// 	//TODO fix db calling from here
// 	// rows, err := database.Query("select * from golang.products")
// 	// if err != nil {
// 	// 	fmt.Println("error")
// 	// }
// 	// defer rows.Close()
// 	// products := []Products{}
// 	// for rows.Next() {
// 	// 	p := Products{}
// 	// 	err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
// 	// 	if err != nil {
// 	// 		fmt.Println(err)
// 	// 		continue
// 	// 	}
// 	// 	products = append(products, p)
// 	// }
// 	p := Products{}
// 	tmpl, _ := template.ParseFiles("./sales/sales/templates/index.html")
// 	tmpl.Execute(w, p)
// 	// data := Products{}
// 	// w.Header().Set("Content-Type", "application/json")
// 	// w.WriteHeader(http.StatusCreated)
// 	// json.NewEncoder(w).Encode(data)
// }

func IndexHandler(w http.ResponseWriter, r *http.Request) {
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

	tmpl, _ := template.ParseFiles("./sales/templates/index.html")
	tmpl.Execute(w, products)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec("delete from golang.products where id = ?", id)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/", 301)
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
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
