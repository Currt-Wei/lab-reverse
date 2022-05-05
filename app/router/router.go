package router

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/controller"
	"lab-reverse/app/middleware"
	"lab-reverse/app/middleware/log"
)

func InitRouter() {
	r := gin.New()
	// 使用自定义的日志中间件
	r.Use(log.LoggerToFile())
	// 默认跨域
	r.Use(middleware.Cors())


	// 使用自定义的jwt认证
	//r.Use(middleware.JWTAuth())
	// 权限验证
	//r.Use(middleware.Authorize())

	r.POST("/addBlackList", controller.AddBlackList)
	r.POST("/deleteBlackList", controller.DeleteBlackList)
	r.GET("/findAllBlackList", controller.FindAllBlackList)
	r.POST("/addCardInfo", controller.AddCardInfo)
	r.GET("/getCardId", controller.GetCardId)
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.POST("/send-verification-code", controller.SendVerificationCode)
	r.POST("/getInsideWeather", controller.GetInsideWeather)
	//r.POST("/saveElectricMeterData", controller.SaveElectricMeterData)
	r.GET("/getHistoryElectricMeterData", controller.GetHistoryElectricMeterData)
	r.POST("/getOutsideWeather", controller.GetOutsideWeather)
	r.POST("/entranceGuard", controller.EntranceGuard)
	r.GET("/findAllAnnouncement", controller.FindAllAnnouncement)
	r.GET("/getElectricMeterData", controller.GetElectricMeterData)
	r.POST("/lightOn", controller.LightOn)
	r.POST("/lightOff", controller.LightOff)
	//r.POST("/PushTopic", controller.PushTopic)
	api:=r.Group("")
	api.Use(middleware.JWTAuth())
	//api.POST("/entranceGuard", controller.EntranceGuard)
	//api.POST("/getOutsideWeather", controller.GetOutsideWeather)

	api.POST("/searchApply", controller.SearchApply)
	api.GET("/getPersonalApply", controller.GetPersonalApply)
	api.GET("/getApply", controller.GetApply)
	api.POST("/refuseApply", controller.RefuseApply)
	api.POST("/allowApply", controller.AllowApply)
	api.POST("/applyForLab", controller.ApplyForLab)

	api.POST("/changeRole", controller.ChangeRole)
	//api.POST("/turnToAdmin", controller.TurnToAdmin)
	api.POST("/addAnnouncement", controller.AddAnnouncement)
	api.POST("/deleteAnnouncement", controller.DeleteAnnouncement)
	//api.GET("/findAllAnnouncement", controller.FindAllAnnouncement)
	//r.POST("/reverseSeat", controller.ReserveSeat)
	api.POST("/reverseSeat", controller.ReserveSeat)
	api.POST("/searchSeat", controller.SearchSeat)
	api.POST("/getInfoByAccount", controller.GetInfoByAccount)
	api.GET("/getReserveInfo", controller.GetReserveInfo)
	api.GET("/getAllReserveInfo", controller.GetAllReserveInfo)
	api.POST("/deleteReserve", controller.DeleteReserve)
	api.GET("/findAllUser", controller.FindAllUser)
	api.POST("/addSeat", controller.AddSeat)
	api.POST("/setBreakdown", controller.SetBreakdown)
	api.POST("/setNormal", controller.SetNormal)
	api.POST("/addLab",controller.AddLab)
	api.POST("/updateLabInfo",controller.UpdateLabInfo)
	api.GET("/findAllLab",controller.FindAllLab)
	api.POST("/findAllSeat",controller.FindAllSeat)
	api.POST("/deleteLab",controller.DeleteLab)

	err := r.Run(":8081")
	if err != nil {
		log.Logger().Errorf("路由初始化失败, %s", err)
		return
	}
}
