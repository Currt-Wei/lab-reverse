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

	//r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	// 使用自定义的jwt认证
	//r.Use(middleware.JWTAuth())
	// 权限验证
	//r.Use(middleware.Authorize())


	r.POST("/turnToUser", controller.TurnToUser)
	r.POST("/turnToAdmin", controller.TurnToAdmin)
	r.POST("/addAnnouncement", controller.AddAnnouncement)
	r.POST("/deleteAnnouncement", controller.DeleteAnnouncement)
	r.GET("/findAllAnnouncement", controller.FindAllAnnouncement)
	r.POST("/reverseSeat", controller.ReserveSeat)
	r.POST("/searchSeat", controller.SearchSeat)
	r.POST("/getInfoByAccount", controller.GetInfoByAccount)
	r.GET("/getReserveInfo", controller.GetReserveInfo)
	r.POST("/deleteReserve", controller.DeleteReserve)
	r.GET("/findAllUser", controller.FindAllUser)
	r.POST("/addSeat", controller.AddSeat)
	r.POST("/addLab",controller.AddLab)
	r.POST("/updateLabInfo",controller.UpdateLabInfo)
	r.GET("/findAllLab",controller.FindAllLab)
	r.POST("/deleteLab",controller.DeleteLab)

	err := r.Run()
	if err != nil {
		log.Logger().Errorf("路由初始化失败, %s", err)
		return
	}
}
