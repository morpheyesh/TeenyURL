package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	muxHandler := handlers()
	crossOrigin := cors.Default().Handler(muxHandler)
	http.Handle("/", crossOrigin)
	fmt.Println("[x] - Server started")
	http.ListenAndServe(":8080", nil)
}

func handlers() *mux.Router {
	router := mux.NewRouter()
  router.HandleFunc("/", AppRedirectHandler).Methods("GET")
	router.HandleFunc("/shorten", ShortenHandler).Methods("GET")
	router.HandleFunc("/lengthen", LengthenHandler).Methods("GET")
	router.HandleFunc("/{short:([a-zA-Z0-9]+$)}", RedirectHandler).Methods("GET")
	return router
}
