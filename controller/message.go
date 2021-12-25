package controller

import (
	"web-wechat/core"
	"web-wechat/global"

	"github.com/gin-gonic/gin"
)

// 发送消息请求体
type sendMsgRes struct {
	// username
	To string `form:"to" json:"to"`
	// 消息类型
	Type string `form:"type" json:"type"`
	// 正文
	Content string `form:"content" json:"content"`
}

// SendTextToFriend 向指定用户发消息
func SendTextToFriend(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	friends, _ := self.Friends(true)
	// 查询指定好友
	friendSearchResult := friends.SearchByUserName(1, res.To)
	if friendSearchResult.Count() < 1 {
		core.FailWithMessage("指定好友不存在", ctx)
		return
	}
	// 取出好友
	friend := friendSearchResult.First()
	// 发送消息
	if _, err := friend.SendText(res.Content); err != nil {
		core.FailWithMessage(err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendTextToGroup 向指定群组发送消息
func SendTextToGroup(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 获取所有群组
	groups, err := self.Groups(true)
	if err != nil {
		core.FailWithMessage("群组获取失败", ctx)
		return
	}
	// 判断指定群组是否存在
	search := groups.SearchByUserName(1, res.To)
	if search.Count() < 1 {
		core.FailWithMessage("指定群组不存在", ctx)
		return
	}
	// 取出指定群组
	group := search.First()
	// 发送消息
	if _, err := group.SendText(res.Content); err != nil {
		core.FailWithMessage(err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}
