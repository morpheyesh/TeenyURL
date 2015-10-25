package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/rs/cors"
)

func main() {
  muxHandler := handlers()
 crossOrigin := cors.Default().Handler(muxHandler)
	http.Handle("/", crossOrigin)
	fmt.Println("[x] - Starting the server")
	http.ListenAndServe(":9000", nil)
}

func handlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", ShortenHandler).Methods("GET")
  router.HandleFunc("/lengthen", LengthenHandler).Methods("GET")
  router.HandleFunc("/{short:([a-zA-Z0-9]+$)}", RedirectHandler).Methods("GET")
	return router
}
