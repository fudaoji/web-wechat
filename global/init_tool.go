package global

import (
	"fmt"
	"log"
	"web-wechat/db"
	"web-wechat/protocol"
)

// InitBotWithStart 系统启动的时候从Redis加载登录信息自动登录
func InitBotWithStart() {
	keys, err := db.GetRedisKeys("")
	if err != nil {
		log.Println("获取Key失败")
		return
	}
	log.Printf("获取到登录用户信息数量：%v", len(keys))
	for _, key := range keys {
		// 提取出AppKey
		appKey := key[13:]
		// 调用热登录
		fmt.Println(appKey)
		bot := InitWechatBotHandle()
		storage := protocol.NewJsonFileHotReloadStorage(key)
		if err := bot.HotLogin(storage, false); err != nil {
			log.Printf("[%v] 热登录失败，错误信息：%v\n", appKey, err.Error())
			// 登录失败，删除热登录数据
			if err := db.DelRedis(key); err != nil {
				log.Printf("[%v] Redis缓存删除失败，错误信息：%v\n", key, err.Error())
			}
			return
		}
		loginUser, _ := bot.GetCurrentUser()
		log.Printf("[%v]初始化自动登录成功，用户名：%v\n", appKey, loginUser.NickName)
		// 登录成功，写入到WechatBots
		SetBot(appKey, bot)
	}
}
