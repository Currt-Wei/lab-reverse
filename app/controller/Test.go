package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/util"
	"log"
	"net/http"
)

func EmailTest(ctx *gin.Context){
	// test email
	err := util.SendEmail("2698230239@qq.com", "123")
	if err != nil {
		log.Println("[register]邮箱发送验证码失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"data": nil,
			"msg": "邮箱发送验证码失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": nil,
		"msg": "邮箱发送验证码成功",
	})
}
