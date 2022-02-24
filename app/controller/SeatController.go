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

func FindAllSeat(ctx *gin.Context){

	var s model.Seat
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.FindAllSeatFail,
			"msg":    "管理员查询所有座位失败",
			"data":   err.Error(),
		})
		return
	}

	lab_name := s.LabName

	var seats []model.Seat

	db:=model.DB
	err:=db.Where("lab_name",lab_name).Find(&seats).Error

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.FindAllSeatFail,
			"data": nil,
			"msg": "管理员查询所有座位失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.FindAllSeatSuccess,
		"data": seats,
		"msg": "查询成功",
	})
}

func SetBreakdown(ctx *gin.Context){

	var s model.Seat
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.SetBreakdownFail,
			"msg":    "添加失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.SetBreakdown(s)

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.SetBreakdownFail,
			"data": nil,
			"msg": "设置座位故障失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.SetBreakdownSuccess,
		"msg": "设置座位故障成功",
	})
}

func SetNormal(ctx *gin.Context){

	var s model.Seat
	if err := ctx.ShouldBindJSON(&s); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.SetNormalFail,
			"msg":    "添加失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.SetNormal(s)

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.SetNormalFail,
			"data": nil,
			"msg": "恢复座位正常失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.SetNormalSuccess,
		"msg": "恢复座位正常成功",
	})
}