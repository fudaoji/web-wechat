package controller

import (
	"fmt"
	"os"
	"path"
	"web-wechat/core"
	"web-wechat/global"
	"web-wechat/protocol"

	"github.com/fudaoji/go-utils"

	"github.com/eatmoreapple/openwechat"
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

// SendImgToGroup 向指定群聊发图片
func SendImgToGroupHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	group, self := FindGroup(bot, res.To, ctx)
	if group == nil {
		return
	}

	filename := path.Base(res.Content)
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取图片失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := group.SendImage(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送图片失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendImgToFriend 向指定用户发图片
func SendImgToFriendHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	friend, self := FindFriend(bot, res.To, ctx)
	if friend == nil {
		return
	}

	filename := path.Base(res.Content)
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取图片失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := friend.SendImage(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送图片失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendTextToFriend 向指定用户发消息
func SendTextToFriendHandle(ctx *gin.Context) {
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
func SendTextToGroupHandle(ctx *gin.Context) {
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

//FindFriend 根据username获取好友
func FindFriend(bot *protocol.WechatBot, username string, ctx *gin.Context) (*openwechat.Friend, *openwechat.Self) {
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	friends, _ := self.Friends(true)
	// 查询指定好友
	friendSearchResult := friends.SearchByUserName(1, username)
	//friendSearchResult := friends.SearchByNickName(1, username)
	if friendSearchResult.Count() < 1 {
		core.FailWithMessage("指定好友不存在", ctx)
		return nil, self
	}
	// 取出好友
	return friendSearchResult.First(), self
}

//FindGroup 根据username获取群组
func FindGroup(bot *protocol.WechatBot, username string, ctx *gin.Context) (*openwechat.Group, *openwechat.Self) {
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	groups, _ := self.Groups(true)
	// 查询指定好友
	searchResult := groups.SearchByUserName(1, username)
	if searchResult.Count() < 1 {
		core.FailWithMessage("指定群组不存在", ctx)
		return nil, self
	}
	// 取出数据
	return searchResult.First(), self
}
