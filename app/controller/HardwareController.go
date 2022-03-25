package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/constant"
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