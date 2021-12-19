package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetLoginCode(t *testing.T) {
	method, url := "GET", Apis["getlogincode"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}

func TestCheckLogin(t *testing.T) {
	method, url := "POST", Apis["checklogin"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}

func TestJson(t *testing.T) {
	params := map[string]interface{}{"name": "doogie", "age": 12}
	fmt.Println(len(params))
	res, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("%#v", string(res))
	fmt.Printf("%#v", string(res))
}
