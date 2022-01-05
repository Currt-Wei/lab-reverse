package model

import "lab-reverse/app/middleware/log"


type User struct {
	Id         int    `gorm:"column:id" json:"id"`
	Name       string `gorm:"column:name" json:"name"`
	Email      string `gorm:"column:email" json:"email"`
	Password   string `gorm:"column:password" json:"password"`
	StuNo      string `gorm:"column:stu_no" json:"stu_no"`
	StuCollege string `gorm:"column:stu_college" json:"stu_college"`
	StuGrade   string `gorm:"column:stu_grade" json:"stu_grade"`

	//Role []Role `gorm:"many2many:user_role" json:"-"`
}

func (S User) TableName() string {
	return "users"
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
