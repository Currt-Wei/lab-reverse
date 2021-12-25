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

	r.POST("/test", controller.Test)

	err := r.Run()
	if err != nil {
		//log.Logger().Errorf("路由初始化失败, %s", err)
		return
	}
}
