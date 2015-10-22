package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	//"io"
	"io/ioutil"
	"net/http"
)

const ( //TODO: Move this out, get it from config
	HOST = "localhost"
)

type UrlData struct {
	LongUrl  string `json:"longUrl"`
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
	shortUrl := "localhost:8000/" + shortKey.Id

	json.NewEncoder(w).Encode(shortUrl)

}

func LengthenHandler(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)

	//TODO: Check the url properly - make it more tighter!

	url := UrlData{}
	json.Unmarshal(body, &url)
	id := string(url.ShortUrl[len(url.ShortUrl)-6:])
	Url, err := GetlongUrl(id)
	if err != nil {
		log.Error("Error in url")
	}
	shortUrl := "http://" + Url.LongUrl
	json.NewEncoder(w).Encode(shortUrl)

}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	short := mux.Vars(r)["short"]
	fmt.Println(short)
	longUrl, err := GetlongUrl(short)
	if err != nil {
		//TODO: http.Redirect(w, r, notfound, http.StatusMovedPermanently)
		log.Error("Cannot redirect")
	}
	http.Redirect(w, r, "http://"+longUrl.LongUrl, http.StatusMovedPermanently)

}
