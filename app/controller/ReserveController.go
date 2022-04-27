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

	b,err:=model.RepeatReserve(r.LabName,r.ReserveDate,r.TimeInterval)
	if b==true{
		c.JSON(http.StatusOK, gin.H{
			"status": constant.AddSeatFail,
			"msg":    "预约失败,不可重复预约同一课室",
		})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.ReserveSeatFail,
			"msg":    err.Error(),
			"data":   "添加失败",
		})
		return
	}


	// todo 黑名单功能
	//b,err=model.InBlackList(r.Account)
	//
	//if err!=nil{
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"status": constant.ReserveSeatFail,
	//		"msg":    err.Error(),
	//		"data":   "添加失败",
	//	})
	//	return
	//}
	//
	//if b==true{
	//	c.JSON(http.StatusOK, gin.H{
	//		"status": constant.RefuseApplyFail,
	//		"data":   "黑名单用户无法预约座位",
	//	})
	//	return
	//}

	u,err := model.GetUserByAccount(r.Account)
	r.UserName=u.Name

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
	labName := r.LabName

	reservations, err:=model.SearchSeat(date,time,labName)

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

// 删除预约
func DeleteReserve(c *gin.Context) {
	var r model.Reservation
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.DeleteReserveFail,
			"msg":    "预约失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.DeleteReserve(&r)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.DeleteReserveFail,
			"msg":    err.Error(),
			"data":   "删除预约记录失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.DeleteReserveSuccess,
		"msg":    "删除预约记录成功",
	})

	return
}