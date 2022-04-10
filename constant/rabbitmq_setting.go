package constant

import (
	"gopkg.in/ini.v1"
	"log"
)

type RabbitMQSet struct {
	Protocol	string	`ini:"protocol"`
	Username    string 	`ini:"username"`
	Password   	string 	`ini:"password"`
	Host		string 	`ini:"host"`
	Port		int 	`ini:"port"`
}

var RabbitMQSetting = &RabbitMQSet{}

func InitRabbitMQSetting() {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Printf("[rabbitmq]配置文件加载错误, %s", err)
	}

	err = cfg.Section("rabbitmq").MapTo(&RabbitMQSetting)
	if err != nil {
		log.Printf("[rabbitmq]配置文件映射错误, %s", err)
	}
}

