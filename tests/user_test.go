package tests

import (
	"fmt"
	"testing"
)

func TestListGroups(t *testing.T) {
	method, url := "GET", Apis["listgroups"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}

func TestListFriends(t *testing.T) {
	method, url := "GET", Apis["listfriends"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}

func TestGetCurrentUser(t *testing.T) {
	method, url := "GET", Apis["getcurrentuser"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}
