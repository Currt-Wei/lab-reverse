package controller

import (
	"github.com/gin-gonic/gin"
	"lab-reverse/app/model"
	"lab-reverse/app/service"
	"lab-reverse/constant"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func AddAnnouncement(ctx *gin.Context){
	var a model.Announcement
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.AddAnnouncementFail,
			"msg":    "添加公告失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.AddAnnouncement(&a)

	if err!=nil {
		log.Println("[AddAnnouncement]添加公告失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.AddAnnouncementFail,
			"err": err,
			"msg": "添加公告失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.AddAnnouncementSuccess,
		"msg": "添加公告成功",
	})
}

func DeleteAnnouncement(ctx *gin.Context){
	var a model.Announcement
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status": constant.DeleteAnnouncementFail,
			"msg":    "删除公告失败",
			"data":   err.Error(),
		})
		return
	}

	err:=model.DeleteAnnouncement(&a)

	if err!=nil {
			log.Println("[DeleteAnnouncement]添加公告失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.DeleteAnnouncementFail,
			"err": err,
			"msg": "删除公告失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.DeleteAnnouncementSuccess,
		"msg": "删除公告成功",
	})
}

func FindAllAnnouncement(ctx *gin.Context){
	// 获取查询数据
	data := ctx.Request.URL.Query()

	// 获取分页数据
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))

	//users, err:=service.FindAllUser(data,limit,page)
	offset := (page - 1) * limit

	var announcements []model.Announcement

	db:=model.DB
	db = db.Limit(limit).Offset(offset)
	// 添加筛选条件
	if data.Get("title") != "" {
		db = db.Where("title LIKE ?", data.Get("title"))
	}

	err := db.Find(&announcements).Error

	if err!=nil {
		log.Println("[DeleteAnnouncement]添加公告失败")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": constant.FindAllAnnouncementFail,
			"err": err,
			"msg": "查找公告失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": constant.FindAllAnnouncementSuccess,
		"msg": "查找公告成功",
		"announcements":announcements,
	})
}

func TurnToAdmin(c *gin.Context){
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.TurnToAdminFail,
			"msg":    "更新失败",
			"data":   err.Error(),
		})
		return
	}

	err := model.TurnToAdmin(u)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.TurnToAdminFail,
			"msg":    err.Error(),
			"data":   "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.TurnToAdminSuccess,
		"msg":    "更新成功",
	})

	return
}

func TurnToUser(c *gin.Context){
	var u model.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": constant.TurnToUserFail,
			"msg":    "更新失败",
			"data":   err.Error(),
		})
		return
	}

	err := model.TurnToUser(u)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.TurnToUserFail,
			"msg":    err.Error(),
			"data":   "更新失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.TurnToAUserSuccess,
		"msg":    "更新成功",
	})

	return
}

func GetHumiture(ctx *gin.Context){

}

func GetAllReserveInfo(c *gin.Context){
	// 获取查询数据
	data := c.Request.URL.Query()


	// 获取分页数据
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	//users, err:=service.FindAllUser(data,limit,page)
	offset := (page - 1) * limit

	db:=model.DB
	db = db.Limit(limit).Offset(offset)

	var reserveInfo []model.Reservation

	// 添加筛选条件
	if data.Get("lab_name") != "" {
		db = db.Where("lab_name LIKE ?", data.Get("lab_name"))
	}

	if data.Get("reserve_date") != "" {
		db = db.Where("reserve_date LIKE ?", data.Get("reserve_date"))
	}

	if data.Get("time_interval") != "" {
		db = db.Where("time_interval LIKE ?", data.Get("time_interval"))
	}

	err:=db.Find(&reserveInfo).Error

	var todo []model.Reservation
	var done []model.Reservation
	for _,info :=range reserveInfo {
		interval :=strings.Split(info.TimeInterval,"-")
		t,_ := time.ParseInLocation("2006-01-02 15:04:05", info.ReserveDate+" "+interval[0],time.Local)
		if t.After(time.Now()){
			todo = append(todo,info)
		} else {
			done = append(done, info)
		}


	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": constant.GetAllReserveInfoFail,
			"msg":    err.Error(),
			"data":   "查询失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": constant.GetAllReserveInfoSuccess,
		"msg":    "查询成功",
		//"data": 	reserveInfo,
		"todo": 	todo,
		"done":		done,
	})

	return
}
