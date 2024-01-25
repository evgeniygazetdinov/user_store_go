package sales

import (
	"fmt"

	"github.com/gorilla/mux"
)

func DoSomeThing() {
	fmt.Println("here")
}

func CallRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/main", indexHandler)
	return router
}
