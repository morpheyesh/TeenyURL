package main

import (
	"fmt"
	//	"github.com/bmizerany/pat"
	"github.com/gorilla/mux"


	//  "encoding/json"
	//  "github.com/morpheyesh/bitty/api"
	"net/http"
	//  "net/url"
)

func main() {

	http.Handle("/", handlers())
	fmt.Println("[x] - Starting the server")
	http.ListenAndServe(":8000", nil)
}

func handlers() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", ShortenHandler).Methods("POST")
	return router
}
