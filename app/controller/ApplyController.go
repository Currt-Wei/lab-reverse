package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/constant"
	"net/http"
	"strconv"
	"strings"
)

func ApplyForLab(ctx *gin.Context){
	var a model.Apply
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.ApplyForLabFail,
			"msg":    "添加失败",
			"data":   err.Error(),
		})
		return
	}

	var err error

	dates:=strings.Split(a.Dates,",")

	for _,date := range dates{
		var apply model.Apply
		apply.ReserveDate=date
		apply.LabName=a.LabName
		apply.Account=a.Account
		apply.UserName=a.UserName
		apply.Status=a.Status
		apply.Description=a.Description
		err=model.ApplyForLab(apply)
		if err!=nil{
			break
		}
	}



	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": constant.ApplyForLabFail,
			"msg":    err.Error(),
			"data":   "申请实验室失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": constant.ApplyForLabSuccess,
		"msg":    "申请实验室成功",
	})

	return
}

func AllowApply(ctx *gin.Context){
	var a model.Apply
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.AllowApplyFail,
			"msg":    "通过申请失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.AllowApply(a)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": constant.AllowApplyFail,
			"msg":    err.Error(),
			"data":   "通过申请失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": constant.AllowApplySuccess,
		"msg":    "通过申请成功",
	})

	return
}

func RefuseApply(ctx *gin.Context){
	var a model.Apply
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.RefuseApplyFail,
			"msg":    "拒绝申请失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.RefuseApply(a)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": constant.RefuseApplyFail,
			"msg":    err.Error(),
			"data":   "拒绝申请失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": constant.RefuseApplySuccess,
		"msg":    "拒绝申请成功",
	})

	return
}

func GetApply(ctx *gin.Context){
	// 获取查询数据
	data := ctx.Request.URL.Query()
	// 获取分页数据
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	//users, err:=service.FindAllUser(data,limit,page)
	offset := (page - 1) * limit

	var applies []model.Apply

	db:=model.DB
	db = db.Limit(limit).Offset(offset)
	// 添加筛选条件
	if data.Get("account") != "" {
		db = db.Where("account LIKE ?", data.Get("account"))
	}
	if data.Get("lab_name") != "" {
		db = db.Where("lab_name LIKE ?", data.Get("lab_name"))
	}else{
		db = db.Where("lab_name", "b3-351")
	}

	err := db.Find(&applies).Error

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.GetApplyFail,
			"data": nil,
			"msg": "管理员查询所有申请失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.GetApplySuccess,
		"data": applies,
		"msg": "查询成功",
	})
}

func GetPersonalApply(ctx *gin.Context){
	// 获取查询数据
	data := ctx.Request.URL.Query()
	// 获取分页数据
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	//users, err:=service.FindAllUser(data,limit,page)
	offset := (page - 1) * limit

	var applies []model.Apply

	db:=model.DB
	db = db.Limit(limit).Offset(offset)
	// 添加筛选条件
	//if data.Get("account") != "" {
		db = db.Where("account LIKE ?", data.Get("account"))
	//}
	if data.Get("lab_name") != "" {
		db = db.Where("lab_name LIKE ?", data.Get("lab_name"))
	}else{
		db = db.Where("lab_name", "b3-351")
	}

	err := db.Find(&applies).Error

	if err!=nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.GetPersonalApplyFail,
			"data": nil,
			"msg": "个人查询所有申请失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.GetPersonalApplySuccess,
		"data": applies,
		"msg": "查询成功",
	})
}