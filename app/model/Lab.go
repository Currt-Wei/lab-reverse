package model

import "lab-reverse/app/middleware/log"

type Lab struct {
	LabId     int    `gorm:"column:lab_id" json:"lab_id"`
	LabName    string    `gorm:"column:lab_name" json:"lab_name"`
	PersonNumber int `gorm:"column:person_number" json:"person_number"`
	Description   string `gorm:"column:description" json:"description"`
}

func (L Lab) TableName() string {
	return "laboratory"
}

func FindLabById(LabId int) (*Lab,error){
	var lab Lab
	err := DB.Where("lab_id = ?", LabId).First(&lab).Error
	if err != nil {
		log.Logger().Errorf("[lab]根据id查找实验室失败，%s", err)
		return nil, err
	}
	return &lab, nil
}

func UpdateLabInfo(lab *Lab) error{
	return DB.Where("lab_id = ?",lab.LabId).Updates(&lab).Error
}

func AddLab(lab *Lab) error{
	return DB.Create(lab).Error
}

func DeleteLab(lab *Lab) error{
	return DB.Where("lab_id=?",lab.LabId).Delete(lab).Error
}