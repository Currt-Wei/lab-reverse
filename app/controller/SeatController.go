package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/constant"
	"net/http"
)

func AddSeat(c *gin.Context) {
	var s model.Seat
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.AddSeatFail,
			"msg":    "添加失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.AddSeat(s)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.AddSeatFail,
			"msg":    err.Error(),
			"data":   "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.AddSeatSuccess,
		"msg":    "添加成功",
	})

	return
}
