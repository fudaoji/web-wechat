package controller

import (
	"web-wechat/core"
	"web-wechat/global"
	"web-wechat/logger"
	"web-wechat/protocol"

	"github.com/eatmoreapple/openwechat"
	"github.com/gin-gonic/gin"
)

// 返回用户信息包装类
type responseUserInfo struct {
	Uin         int64              `json:"uin"`          // 用户唯一ID
	Sex         int                `json:"sex"`          // 性别
	Province    string             `json:"province"`     // 省
	City        string             `json:"city"`         // 市
	Alias       string             `json:"alias"`        // 别名
	DisplayName string             `json:"display_name"` // 显示名称
	NickName    string             `json:"nick_name"`    // 昵称
	RemarkName  string             `json:"remark_name"`  // 备注
	HeadImgUrl  string             `json:"head_img_url"` // 头像
	UserName    string             `json:"user_name"`    // 当前登录中用户的唯一标识
	Members     []*openwechat.User `json:"members"`      // 群成员(群独有)
}

// 返回的好友列表的实体
type friendsResponse struct {
	Count   int                `json:"count"`
	Friends []responseUserInfo `json:"friends"`
}

// 返回的群聊列表的实体
type groupsResponse struct {
	Count  int                `json:"count"`
	Groups []responseUserInfo `json:"groups"`
}

// 修改备注名请求体
type setRemarkNameRes struct {
	// 用户名
	To string `form:"to" json:"to"`
	// 正文
	RemarkName string `form:"remark_name" json:"remark_name"`
}

// SetRemarkNameHandle 修改指定用户的备注
func SetFriendRemarkNameHandle(ctx *gin.Context) {
	// 取出请求参数
	var res setRemarkNameRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	bot := GetCurBot(ctx)
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
	// 设置备注
	if err := self.SetRemarkNameToFriend(friend, res.RemarkName); err != nil {
		core.FailWithMessage("设置备注失败："+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// GetCurrentUserInfoHandle 获取当前登录用户
func GetCurrentUserInfoHandle(ctx *gin.Context) {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")
	bot := global.GetBot(appKey)
	// 获取登录用户信息
	user, err := bot.GetCurrentUser()
	if err != nil {
		core.FailWithMessage("获取登录用户信息失败", ctx)
		return
	}

	logger.Log.Infof("登录用户：%v", user.NickName)
	userDataVO := responseUserInfo{
		Uin:         user.Uin,
		Sex:         user.Sex,
		Province:    user.Province,
		City:        user.City,
		Alias:       user.Alias,
		DisplayName: user.DisplayName,
		NickName:    user.NickName,
		RemarkName:  user.RemarkName,
		HeadImgUrl:  user.HeadImgUrl,
		UserName:    user.UserName,
	}
	core.OkWithData(userDataVO, ctx)
}

// GetFriendsListHandle 获取好友列表
func GetFriendsListHandle(ctx *gin.Context) {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	// 获取好友列表
	user, _ := bot.GetCurrentUser()
	friends, err := user.Friends(true)
	if err != nil {
		core.FailWithMessage("获取好友列表失败", ctx)
		return
	}

	// 循环处理数据
	var friendList []responseUserInfo
	for _, friend := range friends {
		friendList = append(friendList, responseUserInfo{
			Uin:         friend.Uin,
			Sex:         friend.Sex,
			Province:    friend.Province,
			City:        friend.City,
			Alias:       friend.Alias,
			DisplayName: friend.DisplayName,
			NickName:    friend.NickName,
			RemarkName:  friend.RemarkName,
			HeadImgUrl:  friend.HeadImgUrl,
			UserName:    friend.UserName,
		})
	}

	// 返回给前端
	core.OkWithData(friendsResponse{Count: friends.Count(), Friends: friendList}, ctx)
}

// GetGroupsListHandle 获取群聊列表
func GetGroupsListHandle(ctx *gin.Context) {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	// 获取好友列表
	user, _ := bot.GetCurrentUser()

	groups, err := user.Groups(true)
	if err != nil {
		core.FailWithMessage("获取群聊列表失败", ctx)
		return
	}

	logger.Log.Infof("%v", groups)
	// 循环处理数据
	var groupList []responseUserInfo
	for _, group := range groups {
		// 取出群成员
		//members, _ := group.Members()
		groupList = append(groupList, responseUserInfo{
			Uin:         group.Uin,
			Sex:         group.Sex,
			Province:    group.Province,
			City:        group.City,
			Alias:       group.Alias,
			DisplayName: group.DisplayName,
			NickName:    group.NickName,
			RemarkName:  group.RemarkName,
			HeadImgUrl:  group.HeadImgUrl,
			UserName:    group.UserName,
			//Members:     members,
		})
	}

	// 返回给前端
	core.OkWithData(groupsResponse{Count: groups.Count(), Groups: groupList}, ctx)
}

//GetBot 获取当前bot
func GetCurBot(ctx *gin.Context) *protocol.WechatBot {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")
	return global.GetBot(appKey)
}
