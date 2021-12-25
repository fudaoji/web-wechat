package handler

import (
	"fmt"
	. "web-wechat/core"
	"web-wechat/model"

	"github.com/eatmoreapple/openwechat"
)

// 回调请求体
type CallbackRes struct {
	From    string      `form:"from" json:"from"`
	Type    string      `form:"type" json:"type"`
	Content interface{} `form:"content" json:"content"`
}

func HandleMessage(bot *openwechat.Bot) {
	// 定义一个处理器
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// 设置为异步处理
	dispatcher.SetAsync(true)
	// 处理消息为已读
	dispatcher.RegisterHandler(checkIsCanRead, setTheMessageAsRead)

	// 注册文本消息处理函数
	dispatcher.OnText(textMessageHandle)
	// 注册图片消息处理器
	dispatcher.OnImage(imageMessageHandle)
	// 注册表情包消息处理器
	dispatcher.OnEmoticon(emoticonMessageHandle)
	// APP消息处理
	dispatcher.OnMedia(appMessageHandle)
	// 保存消息
	//dispatcher.RegisterHandler(checkNeedSave, saveToDb)
	// 未定义消息处理
	dispatcher.RegisterHandler(checkIsOther, otherMessageHandle)

	// 注册消息处理器
	bot.MessageHandler = openwechat.DispatchMessage(dispatcher)
}

//NotifyWebhook  通知客户端平台
func NotifyWebhook(bot *openwechat.Bot, data *CallbackRes) {
	//uuid, _ := bot.Caller.GetLoginUUID()
	user, _ := bot.GetCurrentUser()
	appkeyRecord := model.Appkey{Uin: user.Uin}
	appkeyRecord.FindByUin()
	fmt.Printf("uin[%v]", user.Uin)
	if len(appkeyRecord.Webhook) > 0 {
		url := appkeyRecord.Webhook
		ReqPostJson(url, data, nil)
	} else {
		fmt.Println("未填写webhook")
	}
}
