package main

import (
	"lab-reverse/app/model"
	"lab-reverse/constant"
	"lab-reverse/util"
	//"lab-reverse/util/casbin"
)

func init() {
	constant.InitMysqlSetting()
	util.InitMysql()
	model.LoadModelDB() // 加载model中使用的db

	// 加载casbin策略
	//casbin.InitCasbinPolicyData()
}
