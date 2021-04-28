package test

import (
	"Api-go/lib"
	"testing"
)

func TestCreateToken(t *testing.T) {
	userName := "mahaoyuan"
	token := lib.CreateToken(userName)
	t.Log(token)
}
