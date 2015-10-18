package main

import (
	"fmt"
	"github.com/bmizerany/pat"
  "github.com/morpheyesh/bitty/api"
	"net/http"
//  "net/url"

)

func main() {
fmt.Println("asd")
fs := http.FileServer(http.Dir("views"))

mux := pat.New()

mux.Get("/shorten", http.HandlerFunc(ShortenHandler))
mux.Post("/lengthen", http.HandlerFunc(LengthenHandler))
mux.Post("/redirect", http.HandlerFunc(RedirectHandler))
//mux.Get("/", fs)

http.Handle("/", mux)
fmt.Println("[x] - Starting the server")
http.ListenAndServe(":8000", nil)
}
