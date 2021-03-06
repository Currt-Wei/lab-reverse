package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/constant"
	"net/http"
	"strconv"
	"strings"
	"time"
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

	user,err :=model.GetUserByAccount(a.Account)

	dates:=strings.Split(a.Dates,",")

	for _,date := range dates{
		var apply model.Apply
		apply.ReserveDate=date
		apply.LabName=a.LabName
		apply.Account=a.Account
		apply.UserName=user.Name
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
	//// 获取分页数据
	//limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	//page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	//
	////users, err:=service.FindAllUser(data,limit,page)
	//offset := (page - 1) * limit

	var applies []model.Apply

	db:=model.DB
	//db = db.Limit(limit).Offset(offset)
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

	var todo []model.Apply
	var done []model.Apply
	for _,info :=range applies {

		t,_ := time.ParseInLocation("2006-01-02 15:04:05", info.ReserveDate+" "+"00:00:00",time.Local)
		if t.After(time.Now()){
			todo = append(todo,info)
		} else {
			done = append(done, info)
		}


	}

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
		"todo": todo,
		"done": done,
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

func SearchApply(ctx *gin.Context){
	var a model.Apply
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.SearchApplyFail,
			"msg":    "查询失败",
			"data":   err.Error(),
		})
		return
	}

	applies,err:=model.SearchApply(a.LabName)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": constant.SearchApplyFail,
			"msg":    "申请实验室失败",
			"data":   err.Error(),
		})
		return
	}

	t:=time.Now().Format("2006-01-02")

	var data []string

	for _,apply := range applies{
		if t<=apply.ReserveDate{
			data=append(data,apply.ReserveDate)
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": constant.SearchApplySuccess,
		"msg":    "申请实验室成功",
		"data": data,
	})

	return
}