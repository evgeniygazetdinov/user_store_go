package sales

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// fix in future one call off db connection from one place

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// products := MainService()
	products := MainService()
	tmpl, _ := template.ParseFiles("./sales/templates/index.html")
	tmpl.Execute(w, products)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idOFProduct := vars["id"]
	DeleteService(idOFProduct)
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

		AddService(model, company, price)
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
