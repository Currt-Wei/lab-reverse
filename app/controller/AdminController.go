package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
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
