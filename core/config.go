package core

import (
	"flag"
	"fmt"
	"web-wechat/utils"

	"github.com/spf13/viper"
)

// RedisConfig Redis配置
var (
	RedisConfig   redisConfig
	MySQLConfig   mysqlConfig
	OssConfig     ossConfig
	MongoDbConfig mongoConfig
)

// Redis配置
type redisConfig struct {
	Host     string // Redis主机
	Port     string // Redis端口
	Password string // Redis密码
	Db       int    // Redis库
}

// MySQL配置
type mysqlConfig struct {
	Host     string // 主机
	Port     string // 端口
	Username string // 用户名
	Password string // 密码
	DbName   string // 数据库名称
}

type ossConfig struct {
	Endpoint        string // 接口地址
	AccessKeyID     string // 账号
	SecretAccessKey string // 密码
	BucketName      string // 桶名称
	UseSsl          bool   // 是否使用SSL
}

type mongoConfig struct {
	Host     string // 地址
	Port     int    // 端口
	Username string // 用户名
	Password string // 密码
	DbName   string // 数据库名称
}

func (c mongoConfig) GetClientUri() string {
	return fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", c.Username, c.Password, c.Host, c.Port, c.DbName)
}

// InitRedisConfig 初始化Redis配置
func InitMysqlConfig() {
	// Mysql配置
	//主机
	host := getVal("db.host", "127.0.0.1")
	// 端口
	port := getVal("db.port", "3306")
	// 密码
	password := getVal("db.password", "")
	// 数据库
	database := getVal("db.database", "test")
	//用户名
	username := getVal("db.username", "root")

	MySQLConfig = mysqlConfig{
		Host:     host,
		Port:     port,
		Password: password,
		Username: username,
		DbName:   database,
	}
}

// InitRedisConfig 初始化Redis配置
func InitRedisConfig() {
	// RedisHost Redis主机
	//host := utils.GetEnvVal("REDIS_HOST", "127.0.0.1")
	host := getVal("redis.host", "127.0.0.1")
	// RedisPort Redis端口
	//port := utils.GetEnvVal("REDIS_PORT", "6379")
	port := getVal("redis.port", "6379")
	// RedisPassword Redis密码
	//password := utils.GetEnvVal("REDIS_PWD", "")
	password := getVal("redis.pwd", "")
	// Redis库
	//db := utils.GetEnvIntVal("REDIS_DB", 0)
	db := getIntVal("redis.db", 0)

	RedisConfig = redisConfig{
		Host:     host,
		Port:     port,
		Password: password,
		Db:       db,
	}
}

// InitOssConfig 初始化OSS配置
func InitOssConfig() {
	endpoint := utils.GetEnvVal("OSS_ENDPOINT", "wechat_oss")
	accessKeyID := utils.GetEnvVal("OSS_KEY", "minio")
	secretAccessKey := utils.GetEnvVal("OSS_SECRET", "minio")
	bucketName := utils.GetEnvVal("OSS_BUCKET", "wechat")
	useSSL := utils.GetEnvBoolVal("OSS_SSL", true)

	OssConfig = ossConfig{
		Endpoint:        endpoint,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		BucketName:      bucketName,
		UseSsl:          useSSL,
	}
}

// InitMongoConfig 初始化MongoDB配置
func InitMongoConfig() {
	host := utils.GetEnvVal("MONGO_HOST", "wechat_mongo")
	port := utils.GetEnvIntVal("MONGO_PORT", 27017)
	user := utils.GetEnvVal("MONGO_USER", "wechat")
	password := utils.GetEnvVal("MONGO_PWD", "wechat")
	dbName := utils.GetEnvVal("MONGO_DB", "web-wechat")

	MongoDbConfig = mongoConfig{host, port, user, password, dbName}
}

//InitConfig 读取配置文件
func InitConfig() {
	mode := flag.String("mode", "dev", "dev mode")
	flag.Parse()
	viper.SetConfigFile(fmt.Sprintf("./setting_%s.yaml", *mode))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

//getVal 获取配置文件的字符串配置值
func getVal(key string, defaultVal string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}
	return defaultVal
}

//getVal 获取配置文件的整形配置值
func getIntVal(key string, defaultVal int) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}
	return defaultVal
}
