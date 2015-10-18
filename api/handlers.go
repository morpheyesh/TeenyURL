package main

import (
  "fmt"
  "encoding/json"

)

func ShortenHandler(w http.ResponseWriter, r *http.Request){
//get the long url and shorten it and store it in redis

fmt.Println("booyah")


fmt.Println("booyah")
x := "test"
json.NewEncoder(w).Encode(x)




}

func LengthenHandler(w http.ResponseWriter, r *http.Request){}


func RedirectHandler(w http.ResponseWriter, r *http.Request) {}
