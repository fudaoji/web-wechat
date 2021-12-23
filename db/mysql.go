package db

import (
	"fmt"
	"web-wechat/core"
	"web-wechat/logger"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type mysqlDB struct {
	*gorm.DB
}

var MysqlClient mysqlDB

// InitMongoConnHandle 初始化MongoDB连接
func InitMysqlConnHandle() {
	// 读取配置
	core.InitMysqlConfig()

	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&loc=Local",
			core.MySQLConfig.Username, core.MySQLConfig.Password, core.MySQLConfig.Host, core.MySQLConfig.Port, core.MySQLConfig.
				DbName))
	fmt.Println(core.MySQLConfig)
	if err != nil {
		panic("failed to connect mysql")
	}
	logger.Log.Info("MysqlDB连接初始化成功")
	MysqlClient = mysqlDB{db}
}
