package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"lab-reverse/app/middleware/log"
	"lab-reverse/app/model"
	"lab-reverse/global"
)

func AddUser(user *model.User) (err error) {

	var user1 *model.User
	user1, _ = model.GetUserByEmail(user.Email)
	if user1 != nil {
		log.Logger().Error("注册失败, 邮箱已注册")
		return fmt.Errorf("注册失败, 邮箱已注册")
	}

	err = model.AddUser(*user)
	if err != nil {
		log.Logger().Errorf("注册失败, %s", err)
		return fmt.Errorf("注册失败, 无法新建用户: %s", err)
	}

	return nil
}

func FindUserByEmail(email string) (*model.User, error) {

	user1, err := model.GetUserByEmail(email)
	return user1, err
}

func FindUserByAccount(account string) (*model.User, error) {

	user1, err := model.GetUserByAccount(account)
	return user1, err
}

func ChangeRole(userId uint, roleId int) (*model.User, error) {
	// 查看该用户是否拥有该角色
	err := global.DB.Where("user_id = ? AND role_id = ?", userId, roleId).First(&model.UserRole{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该用户无此角色")
	}
	var user model.User
	err = global.DB.Where("id = ?", userId).First(&user).Update("role_id", roleId).Error
	return &user, err
}
