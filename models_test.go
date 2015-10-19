package main

import (
	"gopkg.in/check.v1"

	//	"testing"
)

func (s *S) TestGetKey(c *check.C) {

	_, error := GetKey("google.com")
	c.Assert(error, check.IsNil)

}

func (s *S) TestGetlongUrl(c *check.C) {

	_, error := GetlongUrl("asdsd.com/asdasd")
	c.Assert(error, check.IsNil)

}
