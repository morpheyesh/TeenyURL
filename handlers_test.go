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
/*
func (s *S) TestShortenHandler(c *check.C) {
	jsonData := `{"ShortUrl": "localhost:8000/tuW34o"}`

	server = httptest.NewServer(handlers())
	shortenUrl := fmt.Sprintf("%s/shorten?url=google.com", server.URL)

	reader = strings.NewReader(jsonData)

	request, err := http.NewRequest("GET", shortenUrl, reader)
	_, err = http.DefaultClient.Do(request)
	if err != nil {
		c.Error(err)
	}
	c.Assert(err, check.IsNil)

}

*/
func (s *S) TestLengthenHandler(c *check.C) {
	jsonData := `{"ShortUrl": "localhost:8000/tuW34o"}`

	server = httptest.NewServer(handlers())
	lengthenUrl := fmt.Sprintf("%s/lengthen?url=localhost:8000/tuW34o", server.URL)

	reader = strings.NewReader(jsonData)

	request, err := http.NewRequest("GET", lengthenUrl, reader)
	x, err := http.DefaultClient.Do(request)
	fmt.Println(x)
	if err != nil {
		c.Error(err)
	}
	c.Assert(err, check.IsNil)

}

/*
func (s *S) TestRedirectHandler(c *check.C) {

	server = httptest.NewServer(handlers())

	shortenUrl := fmt.Sprintf("%s/tuW34o", server.URL)

	request, err := http.NewRequest("GET", shortenUrl, reader)
	_, err = http.DefaultClient.Do(request)
	if err != nil {
		c.Error(err)
	}
	c.Assert(err, check.IsNil)

}
*/
