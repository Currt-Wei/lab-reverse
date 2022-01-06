package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/app/service"
	"lab-reverse/constant"
	"log"
	"net/http"
	"strconv"
)

func FindAllUser(ctx *gin.Context){
	// 获取查询数据
	data := ctx.Request.URL.Query()
	// 获取分页数据
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	//users, err:=service.FindAllUser(data,limit,page)
	offset := (page - 1) * limit

	var users []model.User

	db:=model.DB
	db = db.Limit(limit).Offset(offset)
	// 添加筛选条件
	if data.Get("account") != "" {
		db = db.Where("account LIKE ?", data.Get("account"))
	}
	if data.Get("name") != "" {
		db = db.Where("name LIKE ?", data.Get("name"))
	}
	if data.Get("college") != "" {
		db = db.Where("college LIKE ?", data.Get("college"))
	}
	if data.Get("degree") != "" {
		db = db.Where("degree LIKE ?", data.Get("degree"))
	}
	if data.Get("identity") != "" {
		db = db.Where("identity LIKE ?", data.Get("identity"))
	}

	err := db.Find(&users).Error

	if err!=nil {
		log.Println("[FindAllUser]管理员查询所有用户失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.FindAllUserFail,
			"data": nil,
			"msg": "管理员查询所有用户失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.FindAllUserSuccess,
		"data": users,
		"msg": "查询成功",
	})
}

func UpdateLabInfo(c *gin.Context){
	var lab model.Lab
	if err := c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.UpdateLabFail,
			"msg":    "更新失败",
			"data":   err.Error(),
		})
		return
	}

	err:=service.UpdateLabInfo(&lab)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.UpdateLabFail,
			"msg":    err.Error(),
			"data":   "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.UpdateLabSuccess,
		"msg":    "更新成功",
	})

	return
}

// 添加实验室
func AddLab(c *gin.Context){
	var lab model.Lab
	if err := c.ShouldBindJSON(&lab); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.AddLabFail,
			"msg":    "添加失败",
			"data":   err.Error(),
		})
		return
	}

	err:=service.AddLab(&lab)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.AddLabFail,
			"msg":    err.Error(),
			"data":   "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.AddLabSuccess,
		"msg":    "更新成功",
	})

	return
}

func FindAllLab(ctx *gin.Context){
	// 获取查询数据
	data := ctx.Request.URL.Query()
	// 获取分页数据
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	//users, err:=service.FindAllUser(data,limit,page)
	offset := (page - 1) * limit

	var labs []model.Lab

	db:=model.DB
	db = db.Limit(limit).Offset(offset)
	// 添加筛选条件
	if data.Get("lab_name") != "" {
		db = db.Where("lab_name LIKE ?", data.Get("lab_name"))
	}
	if data.Get("person_number") != "" {
		db = db.Where("person_number LIKE ?", data.Get("person_number"))
	}

	err := db.Find(&labs).Error

	if err!=nil {
		log.Println("[FindAllLab]管理员查询所有实验室失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.FindAllLabFail,
			"data": nil,
			"msg": "管理员查询所有用户失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.FindAllLabSuccess,
		"data": labs,
		"msg": "查询成功",
	})
}

func DeleteLab(ctx *gin.Context) {
	var lab model.Lab
	if err := ctx.ShouldBindJSON(&lab); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.DeleteLabFail,
			"msg":    "添加失败",
			"data":   err.Error(),
		})
		return
	}

	err:=service.DeleteLab(&lab)

	if err!=nil {
		log.Println("[DeleteLab]管理员删除实验室失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.FindAllLabFail,
			"err": err,
			"msg": "删除指定实验室失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.FindAllLabSuccess,
		"msg": "删除成功",
	})
}