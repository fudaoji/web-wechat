package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	APPKEY  = "3156"
	APIHOST = "http://127.0.0.1:8888"
)

var (
	Apis = map[string]string{
		"getlogincode":   APIHOST + "/getlogincode",
		"checklogin":     APIHOST + "/checklogin",
		"getcurrentuser": APIHOST + "/user/info",
		"listfriends":    APIHOST + "/user/friends",
		"listgroups":     APIHOST + "/user/groups",
		"msgToUser":      APIHOST + "/message/user",
		"msgToGroup":     APIHOST + "/message/group",
	}
	Uuid = "YaW_Qwx3PA=="
)

func dealPayload(data interface{}) io.Reader {
	if data != nil {
		var payload *bytes.Reader
		if c, err := json.Marshal(data); err == nil {
			payload = bytes.NewReader(c)
		}
		return payload
	} else {
		return nil
	}
}

//请求封装
func Request(method string, url string, data interface{}) map[string]interface{} {
	client := &http.Client{}
	req, err := http.NewRequest(method, url+"?uuid="+Uuid, dealPayload(data))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("AppKey", APPKEY)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	//fmt.Println(string(body))
	m := make(map[string]interface{})
	if err := json.Unmarshal(body, &m); err != nil {
		// 处理错误
		fmt.Println("unmarshal错误:" + err.Error())
	}
	return m
}
