package tests

import (
	"fmt"
	"testing"
)

// 发送消息请求体
type sendMsg struct {
	// 送达人UserName
	To string
	// 消息类型
	Type int
	// 正文
	Content string
}

//发送消息给好友
func TestMsgToUser(t *testing.T) {
	method, url := "PUT", Apis["msgToUser"]
	res := Request(method, url, sendMsg{To: "道集", Type: 1, Content: "hi"})
	fmt.Printf("%#v", res)
}

//发送消息给群聊
func TestMsgToGroup(t *testing.T) {
	method, url := "PUT", Apis["msgToGroup"]
	res := Request(method, url, sendMsg{To: "有家、有爱", Type: 1, Content: "hi"})
	fmt.Printf("%#v", res)
}
