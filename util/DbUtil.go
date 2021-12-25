package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lab-reverse/app/middleware/log"
	"lab-reverse/constant"
	"time"
)

var db *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		constant.MysqlSetting.Username,
		constant.MysqlSetting.Password,
		constant.MysqlSetting.Ip,
		constant.MysqlSetting.Port,
		constant.MysqlSetting.Database,
	)
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Logger().Errorf("[mysql]连接数据库失败, %s", err)
		panic("failed to connect database")
	}

	sqlDB, err := conn.DB()

	// 设置数据库连接池
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
