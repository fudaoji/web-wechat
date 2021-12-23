package model

import (
	. "web-wechat/db"
)

const tableName = "appkey"

// Appkey appkey表
type Appkey struct {
	Id       int    `json:"id" gorm:"type:int(10);comment:ID"`
	Appkey   string `json:"appkey" gorm:"type:varchar(32);comment:appKey"`
	Deadline int64  `json:"deadline" gorm:"type:int(10);comment:到期时间"`
}

//findByAppkey 根据appkey找记录
func (a *Appkey) FindByAppkey() {
	MysqlClient.Table(tableName).Where("appkey = ?", a.Appkey).First(a)
}
