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
		ctx.JSON(http.StatusOK, gin.H{
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
		"data": model.InsideWeather,
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
		ctx.JSON(http.StatusOK, gin.H{
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

		ctx.JSON(http.StatusOK, gin.H{
			"status": constant.EntranceGuardFail,
			"msg":    "申请门禁失败,未到预约时间",
			"data":   "",
		})
		return
	}
	// todo 请求门禁系统
	model.OpenDoor("111")

	ctx.JSON(http.StatusOK, gin.H{
		"status": constant.EntranceGuardSuccess,
		"msg":    "申请门禁成功",
	})

	return
}

func GetElectricMeterData(ctx *gin.Context){
	var electricMeterData model.ElectricMeterData
	electricMeterData.Voltage=model.Vol
	electricMeterData.Current=model.Cur
	electricMeterData.Active_power=model.ActivePower
	electricMeterData.Reactive_power=model.ReactivePower
	electricMeterData.Apparent_power=model.ApparentPower
	electricMeterData.Factor=model.Fac
	electricMeterData.Angel=model.Ang
	electricMeterData.Neutral=model.Neutral
	electricMeterData.Frequency=model.Frequency
	electricMeterData.Temperature=model.Temperature
	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.GetOutsideWeatherSuccess,
		"data": electricMeterData,
		"msg": "查询成功",
	})
}

func LightOn(ctx *gin.Context) {

	model.LightOn()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg": "操作成功",
	})
}

func LightOff(ctx *gin.Context) {

	model.LightOff()

	ctx.JSON(http.StatusOK, gin.H{
		"code": 400,
		"msg": "操作成功",
	})
}

func GetCardId(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": model.CardId,
	})
}

func AddCardInfo(ctx *gin.Context){
	var card model.Card
	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "添加卡信心失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.AddCard(card)

	if err!=nil{
		ctx.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"msg":    "添加卡信心失败",
			"data":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "添加卡信心成功",
		"data":   err.Error(),
	})
}