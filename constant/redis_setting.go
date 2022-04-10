package constant

import (
	"gopkg.in/ini.v1"
	"log"
)

type Redis struct {
	Addr       	string `ini:"addr"`
	Password   	string `ini:"password"`
	DB			int `ini:"db"`
}

var RedisSetting = &Redis{}

func InitRedisSetting() {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Printf("[redis]配置文件加载错误, %s", err)
	}

	err = cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Printf("[redis]配置文件映射错误, %s", err)
	}
}
