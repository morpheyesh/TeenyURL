package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	http.Handle("/", handlers())
	fmt.Println("[x] - Starting the server")
	http.ListenAndServe(":8000", nil)
}

func handlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", ShortenHandler).Methods("GET")
  router.HandleFunc("/lengthen", LengthenHandler).Methods("GET")
  router.HandleFunc("/{short:([a-zA-Z0-9]+$)}", RedirectHandler).Methods("GET")
	return router
}
