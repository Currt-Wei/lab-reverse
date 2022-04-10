package common

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab-reverse/constant"
	"log"
)

func InitDB() *gorm.DB {
	constant.InitMysqlSetting()
	setting := constant.MysqlSetting

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.Username, setting.Password, setting.Ip, setting.Port, setting.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("数据库连接失败！")
	}
	return db
}

func InitRedis() *redis.Client {
	constant.InitRedisSetting()
	setting := constant.RedisSetting

	// 连接到Redis
	client := redis.NewClient(&redis.Options{
		Addr:     setting.Addr,
		Password: setting.Password, // no password set
		DB:       setting.DB,  // use default DB
	})

	return client
}

func InitRabbitMQConn() *amqp.Connection {
	constant.InitRabbitMQSetting()
	setting := constant.RabbitMQSetting

	amqps := fmt.Sprintf("%s://%s:%s@%s:%d/",
		setting.Protocol,
		setting.Username,
		setting.Password,
		setting.Host,
		setting.Port,
	)
	conn, _ := amqp.Dial(amqps)
	return conn
}