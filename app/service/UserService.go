package service

import (
	"fmt"
	"lab-reverse/app/middleware/log"
	"lab-reverse/app/model"
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

func ReverseSeat(book model.Book) error {
	err:=model.ReverseSeat(book)
	return err
}