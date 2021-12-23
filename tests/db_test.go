package tests

import (
	"fmt"
	"testing"
	. "web-wechat/model"
)

func TestGetAppkey(t *testing.T) {
	var key = Appkey{Appkey: "1231231"}
	fmt.Println(key)
	key.FindByAppkey()
	if key.Deadline > 0 {
		fmt.Sprintf("%s", "exists")
	} else {
		fmt.Println("not exists")
	}
}
