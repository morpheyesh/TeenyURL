package main

import (
	"encoding/json"
//	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	//"io"
	//"io/ioutil"
	"net/http"
)

const ( //TODO: Move this out, get it from config
	HOST = "localhost"
)

type UrlData struct {
	LongUrl  string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}

//TODO: url checking needs to be tighter
func ShortenHandler(w http.ResponseWriter, r *http.Request) {

	longUrl := r.FormValue("url")
	//check url , create key, throw err if url is bad
	shortKey, err := GetKey(longUrl)
	if err != nil {
		log.Error("error in url")
	}

	//TODO: Get it from conf
	shortUrl := "localhost:9000/" + shortKey.Id

	json.NewEncoder(w).Encode(shortUrl)

}

func LengthenHandler(w http.ResponseWriter, r *http.Request) {

	ShortUrl := r.FormValue("url")

	id := string(ShortUrl[len(ShortUrl)-6:])
	Url, err := GetlongUrl(id)
	if err != nil {
		log.Error("Error in url")
	}
	shortUrl := "http://" + Url.LongUrl
	json.NewEncoder(w).Encode(shortUrl)

}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	short := mux.Vars(r)["short"]
	longUrl, err := GetlongUrl(short)
	if err != nil {
		//TODO: http.Redirect(w, r, notfound, http.StatusMovedPermanently)
		log.Error("Cannot redirect")
	}
	http.Redirect(w, r, "http://"+longUrl.LongUrl, http.StatusMovedPermanently)

}
