package route

import (
	"web-wechat/controller"

	"github.com/gin-gonic/gin"
)

// 初始化消息相关路由
func initMessageRoute(app *gin.Engine) {
	group := app.Group("/message")

	// 向指定好友发送图片消息
	group.POST("/user/img", controller.SendImgToFriend)

	// 向指定好友发送文本消息
	group.POST("/user", controller.SendTextToFriend)

	// 向指定群组发送消息
	group.PUT("/group", controller.SendTextToGroup)
}
