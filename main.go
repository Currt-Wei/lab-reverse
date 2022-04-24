package main

import (
	"lab-reverse/app/model"
	"lab-reverse/app/router"
	"lab-reverse/constant"
	"lab-reverse/global"
	"lab-reverse/util"
)

func main() {
	global.InitGlobalVariable()
	constant.InitEmailSetting()
	//util.InitElecMQTT()
	//util.InitESPTHMQTT()
	//util.InitESPDoorMQTT()
	constant.InitMysqlSetting()
	util.InitMysql()
	model.LoadModelDB() // 加载model中使用的db
	router.InitRouter()
}



