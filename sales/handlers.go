package sales

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	products := MainService()
	tmpl, _ := template.ParseFiles("./sales/templates/index.html")
	tmpl.Execute(w, products)

}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idOFProduct := vars["id"]
	DeleteProductService(idOFProduct)
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

		AddProductService(model, company, price)
		http.Redirect(w, r, "/", 301)
	} else {
		http.ServeFile(w, r, "./sales/templates/create.html")
	}
}

func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	prod, err := EditProductService(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	} else {
		tmpl, _ := template.ParseFiles("./sales/templates/edit.html")
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
		err = EditProductSpecitcService(id, model, company, price)
		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	}
}
