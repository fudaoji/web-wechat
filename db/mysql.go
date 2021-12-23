package db

import (
	"context"
	"time"
	"web-wechat/core"
	"web-wechat/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db struct {
	db *mongo.Client
}

var MongoClient mongoDBClient

// InitMongoConnHandle 初始化MongoDB连接
func InitMongoConnHandle() {
	// 读取配置
	core.InitMongoConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel() // 在调用WithTimeout之后defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(core.MongoDbConfig.GetClientUri()))
	if err != nil {
		logger.Log.Panicf("MongoDB初始化连接失败: %v", err.Error())
		//os.Exit(1)
	}
	logger.Log.Info("MongoDB连接初始化成功")
	//mongoClient = client
	MongoClient = mongoDBClient{client: client}
}
