package main

import (
	"fmt"
	"gopkg.in/check.v1"
	//   "encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

var (
	server   *httptest.Server
	reader   io.Reader
	usersUrl string
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})



func (s *S) TestShortenHandler(t *check.C) {
	jsonData := `{"longUrl": "bleeh.com"}`

	server = httptest.NewServer(handlers())

	shortenUrl := fmt.Sprintf("%s/shorten", server.URL)

	reader = strings.NewReader(jsonData)

	request, err := http.NewRequest("POST", shortenUrl, reader)
	res, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Error(err)
	}


}
