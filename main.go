package main

import (
	"lab-reverse/app/model"
	"lab-reverse/app/router"
	"lab-reverse/constant"
	"lab-reverse/util"
)

func main() {
	constant.InitEmailSetting()
	//util.InitMQTT()
	constant.InitMysqlSetting()
	util.InitMysql()
	model.LoadModelDB() // 加载model中使用的db
	router.InitRouter()
}
