package model

type Apply struct {
	Id        int       `gorm:"column:id" json:"id"`
	CreatedAt	Time	`gorm:"type:timestamp;autoCreatTime" json:"created_at" swaggerignore:"true"`
	UserName     string    `gorm:"column:user_name" json:"user_name"`
	Account     string    `gorm:"column:account" json:"account"`
	LabName     string    `gorm:"column:lab_name" json:"lab_name"`
	ReserveDate	 string    `gorm:"column:reserve_date" json:"reserve_date"`
	Status    int    `gorm:"column:status" json:"status"`
	Description    string    `gorm:"column:description" json:"description"`
	Dates	string	`json:"dates"`
}

func (A Apply) TableName() string {
	return "apply"
}

func ApplyForLab(A Apply) error{
	return DB.Create(&A).Error
}

func AllowApply(A Apply) error{
	return DB.Model(&A).Where("account",A.Account).Where("lab_name",A.LabName).Update("status","1").Error
}

func RefuseApply(A Apply) error{
	return DB.Model(&A).Where("account",A.Account).Where("lab_name",A.LabName).Update("status","2").Error
}