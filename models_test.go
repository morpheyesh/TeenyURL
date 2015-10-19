package main

import (
	"fmt"
	"gopkg.in/check.v1"
	//   "encoding/json"
	//"io"
//	"net/http"
//	"net/http/httptest"
//	"strings"

	"testing"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})


func (s *S) TestGetKey(t *check.C) {

x,_ := GetKey("gogole.com")
 fmt.Println(x.Id)
 fmt.Println(x.LongUrl)
 fmt.Println(x.CreatedAt)
}
