package main

import (
	"fmt"
	//	"github.com/gorilla/mux"
	"encoding/json"
  log "github.com/Sirupsen/logrus"

	//"io"
	"io/ioutil"
	"net/http"
)

const ( //TODO: Move this out, get it from config
	HOST = "localhost"
)

type UrlData struct {
	LongUrl string `json:"longUrl"`
  ShortUrl string `json:"shortUrl"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	url := UrlData{}
	 json.Unmarshal(body, &url)

	//check url , create key, throw err if url is bad
	shortKey, err := GetKey(url.LongUrl)
	if err != nil {
	 log.Error("error in url")
	}

  //TODO: Get it from conf
	shortUrl := "localhost:8000/"+shortKey.Id

	json.NewEncoder(w).Encode(shortUrl)

}

func LengthenHandler(w http.ResponseWriter, r *http.Request) {

  body, _ := ioutil.ReadAll(r.Body)

  url := UrlData{}
  json.Unmarshal(body, &url)

  longUrl, err := GetlongUrl(url.ShortUrl)
  if err != nil {
    log.Error("Error in url")
  }
  fmt.Println(longUrl)

}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {}
