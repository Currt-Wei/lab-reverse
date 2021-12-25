package model

import (
	"gorm.io/gorm"
	"lab-reverse/util"
)

var DB *gorm.DB

func LoadModelDB() {
	DB = util.GetDB()
}
