package constant

import (
	"github.com/patrickmn/go-cache"
	"gopkg.in/ini.v1"
	"log"
	"time"
)

type Email struct {
	Username	string `ini:"username"`
	Password	string `ini:"password"`
	Host string `ini:"host"`
	Port string `ini:"port"`
}

var EmailSetting = &Email{}
var EmailCache *cache.Cache

func InitEmailSetting() {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Printf("[email]配置文件加载错误, %s", err)
	}

	err = cfg.Section("email").MapTo(EmailSetting)
	if err != nil {
		log.Printf("[email]配置文件映射错误, %s", err)
	}

	// 新建缓存来保存邮箱验证码
	EmailCache = cache.New(5 * time.Minute, 10 * time.Minute)
}

