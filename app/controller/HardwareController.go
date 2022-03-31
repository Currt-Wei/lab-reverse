package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/constant"
	"lab-reverse/util"
	"net/http"
)

func GetOutsideWeather(ctx *gin.Context){

	weather, err:=model.GetOutsideWeather()

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.GetOutsideWeatherFail,
			"data": nil,
			"msg": "个人查询所有申请失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.GetOutsideWeatherSuccess,
		"data": weather,
		"msg": "查询成功",
	})
}

func EntranceGuard(ctx *gin.Context)  {
	var u model.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.EntranceGuardFail,
			"msg":    "申请门禁失败",
			"data":   err.Error(),
		})
		return
	}

	b,err:=model.SearchReserve(u.Account)

	if err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.EntranceGuardFail,
			"msg":    "申请门禁失败",
			"data":   err.Error(),
		})
		return
	}

	if b==false{

		user,_:=model.GetUserByAccount(u.Account)
		// email
		util.SendEmail(user.Email, "申请门禁失败,未到预约时间")

		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.EntranceGuardFail,
			"msg":    "申请门禁失败,未到预约时间",
			"data":   "",
		})
		return
	}
	// todo 请求门禁系统


	ctx.JSON(http.StatusOK, gin.H{
		"status": constant.EntranceGuardSuccess,
		"msg":    "申请门禁成功",
	})

	return
}