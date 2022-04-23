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

func GetInsideWeather(ctx *gin.Context){

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.GetInsideWeatherSuccess,
		"data": util.InsideWeather,
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
	util.OpenDoor()

	ctx.JSON(http.StatusOK, gin.H{
		"status": constant.EntranceGuardSuccess,
		"msg":    "申请门禁成功",
	})

	return
}

func GetElectricMeterData(ctx *gin.Context){
	var electricMeterData model.ElectricMeterData
	electricMeterData.Voltage=util.Vol
	electricMeterData.Current=util.Cur
	electricMeterData.Active_power=util.ActivePower
	electricMeterData.Reactive_power=util.ReactivePower
	electricMeterData.Apparent_power=util.ApparentPower
	electricMeterData.Factor=util.Fac
	electricMeterData.Angel=util.Ang
	electricMeterData.Neutral=util.Neutral
	electricMeterData.Frequency=util.Frequency
	electricMeterData.Temperature=util.Temperature
	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.GetOutsideWeatherSuccess,
		"data": electricMeterData,
		"msg": "查询成功",
	})
}

func LightOn(ctx *gin.Context) {

	util.LightOn()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg": "操作成功",
	})
}

func LightOff(ctx *gin.Context) {

	util.LightOff()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg": "操作成功",
	})
}