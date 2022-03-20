package model

import (
	"lab-reverse/app/middleware/log"
)


type User struct {
	Id uint `json:"id"`
	CreatedAt Time	`gorm:"type:timestamp;autoCreatTime" json:"created_at" swaggerignore:"true"`
	UpdatedAt Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at" swaggerignore:"true"`
	Account string	`gorm:"type:char(12);not null;uniqueIndex" json:"account" validate:"required"`
	Name string		`gorm:"type:varchar(50)" validate:"required"`
	Email string	`gorm:"type:varchar(50)" validate:"required,email"`
	Telephone string	`gorm:"type:varchar(11)" validate:"required,len=11"`
	College string		`gorm:"type:varchar(50)"`
	Password string		`gorm:"type:varchar(255)" validate:"required"`
	Degree string		`gorm:"type:varchar(10)"`
	Grade string		`gorm:"type:varchar(10)"`
	Identity string		`gorm:"type:varchar(10);default:student"`
	Enable int			`gorm:"type:smallint;default:1" json:"enable"`

	RoleId string		`gorm:"type:varchar(10);default:user"`
}

func (S User) TableName() string {
	return "user"
}

// GetAllUsers 查询所有用户
func GetAllUsers() ([]User, error) {
	users := make([]User, 10)
	err := DB.Preload("Role").Find(&users).Error
	if err != nil {
		log.Logger().Errorf("[user]查询所有用户失败，%s", err)
		return nil, err
	}
	return users, nil
}

// GetUserByEmail 根据email查找用户
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Logger().Errorf("[user]根据email查找用户失败，%s", err)
		return nil, err
	}
	return &user, nil
}

// GetUserByAccount 根据Account查找用户
func GetUserByAccount(account string) (*User, error) {
	var user User
	err := DB.Where("account = ?", account).First(&user).Error
	if err != nil {
		log.Logger().Errorf("[user]根据email查找用户失败，%s", err)
		return nil, err
	}
	return &user, nil
}

func GetUserById(id int) (*User, error) {
	var user User
	err := DB.Preload("Role").Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Logger().Errorf("[user]根据id查找用户失败，%s", err)
		return nil, err
	}
	return &user, nil
}

// AddUser 添加用户
func AddUser(user User) error {
	return DB.Create(&user).Error
}

// 查询历史预约信息
func GetReserveInfo(UserId uint) ([]Reservation,error){
	var ReservationInfo []Reservation
	err := DB.Where("user_id = ?", UserId).Find(&ReservationInfo).Error
	return ReservationInfo,err
}

func TurnToAdmin(user User) error{
	return DB.Model(&user).Where("account",user.Account).Update("RoleId","1").Error
}

func TurnToUser(user User) error{
	return DB.Model(&user).Where("account",user.Account).Update("RoleId","3").Error
}
