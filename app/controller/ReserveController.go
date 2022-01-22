package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/constant"
	"net/http"
)

func ReserveSeat(c *gin.Context) {
	var r model.Reservation
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.AddSeatFail,
			"msg":    "预约失败",
			"data":   err.Error(),
		})
		return
	}

	seat,err := model.FindSeatBySeatId(r.SeatId)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.ReserveSeatFail,
			"msg":    err.Error(),
			"data":   "添加失败",
		})
		return
	}

	r.SeatName=seat.SeatName

	err=model.ReserveSeat(&r)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.ReserveSeatFail,
			"msg":    err.Error(),
			"data":   "添加失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.ReserveSeatSuccess,
		"msg":    "添加成功",
	})

	return
}

func SearchSeat(c *gin.Context) {
	var r model.Reservation
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.AddSeatFail,
			"msg":    "预约失败",
			"data":   err.Error(),
		})
		return
	}

	date := r.ReserveDate
	time := r.TimeInterval
	labId := r.LabId

	reservations, err:=model.SearchSeat(date,time,labId)

	var seatNames []string

	for _,reservation := range reservations{
		seatNames=append(seatNames,reservation.SeatName)
	}


	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.SearchSeatFail,
			"msg":    err.Error(),
			"data":   "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.SearchSeatSuccess,
		"msg":    "查询成功",
		"data": seatNames,
	})

	return
}