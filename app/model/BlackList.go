package model

type BlackList struct {
	Id        int       `gorm:"column:id" json:"id"`
	Account	string	`json:"account" gorm:"column:account"`
}

func (b BlackList) TableName() string {
	return "black_list"
}

func InBlackList(account string) (bool,error){
	var blackList BlackList
	err := DB.Where("account = ?", account).Find(&blackList).Error
	if err!=nil||blackList.Account!=""{
		return true,err
	}

	return false,err

}
