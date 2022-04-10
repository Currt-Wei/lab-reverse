package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"lab-reverse/app/common"
	"lab-reverse/app/model"
)

var DB *gorm.DB
var RedisClient *redis.Client
var Context context.Context
var RabbitMQ *common.Rabbitmq

func InitGlobalVariable() {
	DB = common.InitDB()
	RedisClient = common.InitRedis()
	Context = context.Background()
	RabbitMQ, _ = common.NewRabbitmq()

	// 创建队列
	RabbitMQ.CreateQueue("email.code.wjh")
	// 消费者
	go RabbitMQ.ConsumeQueue("email.code.wjh")
}

// ClaimsToUser 将jwt的claims转换成user，兼容原有的接口
func ClaimsToUser(claims *common.Claims) model.User {
	return model.User{
		Id: claims.UserId,
		Account: claims.Account,
		Enable: claims.Enable,
		RoleId: claims.RoleId,
		Identity: claims.Identity,
		Email: claims.Email,
	}
}

